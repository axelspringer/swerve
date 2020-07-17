package database

import (
	"encoding/base64"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// CreateRedirect creates a new redirect entry
func (d *Database) CreateRedirect(redirect Redirect) error {
	tablePrefix := d.Config.TableNamePrefix
	tableName := d.Config.TableRedirects

	redirect.Created = int(time.Now().Unix())
	redirect.Modified = redirect.Created

	_, err := d.GetRedirectByDomain(redirect.RedirectTo)
	if err == nil {
		return errors.New(ErrRedirectExists)
	}

	payload, err := dynamodbattribute.MarshalMap(redirect)
	if err != nil {
		return errors.WithMessage(err, ErrRedirectMarshal)
	}

	_, err = d.Service.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tablePrefix + tableName),
		Item:      payload,
	})

	return err
}

// GetRedirectByDomain returns one redirect entry by domain name
func (d *Database) GetRedirectByDomain(name string) (Redirect, error) {
	tablePrefix := d.Config.TableNamePrefix
	tableName := d.Config.TableRedirects
	redirect := Redirect{}

	res, err := d.Service.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tablePrefix + tableName),
		Key: map[string]*dynamodb.AttributeValue{
			keyNameRedirectsTable: {
				S: aws.String(name),
			},
		},
	})
	if err != nil {
		return redirect, errors.WithMessage(err, ErrRedirectsFetch)
	}

	if len(res.Item) == 0 {
		return redirect, errors.New(ErrRedirectNotFound)
	}

	if err = dynamodbattribute.UnmarshalMap(res.Item, &redirect); err != nil {
		return redirect, errors.WithMessage(err, ErrRedirectMarshal)
	}

	return redirect, nil
}

// DeleteRedirectByDomain deletes a redirect entry by domain name
func (d *Database) DeleteRedirectByDomain(name string) error {
	tablePrefix := d.Config.TableNamePrefix
	tableName := d.Config.TableRedirects

	_, err := d.Service.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(tablePrefix + tableName),
		Key: map[string]*dynamodb.AttributeValue{
			keyNameRedirectsTable: {
				S: aws.String(name),
			},
		},
	})
	return err
}

// UpdateRedirectByDomain updates a redirect entry
func (d *Database) UpdateRedirectByDomain(name string, redirect Redirect) error {
	tablePrefix := d.Config.TableNamePrefix
	tableName := d.Config.TableRedirects

	currentRedirect, err := d.GetRedirectByDomain(name)
	if err != nil {
		return errors.WithMessage(err, ErrRedirectNotExist)
	}

	redirect.Created = currentRedirect.Created
	redirect.Modified = int(time.Now().Unix())

	if name != redirect.RedirectFrom {
		if err := d.CreateRedirect(redirect); err != nil {
			return errors.WithMessage(err, ErrRedirectCreate)
		}
		if err := d.DeleteRedirectByDomain(name); err != nil {
			return errors.WithMessage(err, ErrRedirectDelete)
		}
		return nil
	}

	update := expression.UpdateBuilder{}
	if currentRedirect.Description != redirect.Description {
		update = update.Set(expression.Name(attrNameDescription),
			expression.Value(redirect.Description))
	}
	if currentRedirect.RedirectTo != redirect.RedirectTo {
		update = update.Set(expression.Name(attrNameRedirect),
			expression.Value(redirect.RedirectTo))
	}
	if currentRedirect.Promotable != redirect.Promotable {
		update = update.Set(expression.Name(attrNamePromotable),
			expression.Value(redirect.Promotable))
	}
	if currentRedirect.Code != redirect.Code {
		update = update.Set(expression.Name(attrNameCode),
			expression.Value(redirect.Code))
	}
	if currentRedirect.CPathMaps != redirect.CPathMaps {
		update = update.Set(expression.Name(attrNamePathMap),
			expression.Value(redirect.CPathMaps))
	}
	if currentRedirect.Modified != redirect.Modified {
		update = update.Set(expression.Name(attrNameModified),
			expression.Value(redirect.Modified))
	}

	expr, err := expression.NewBuilder().
		WithUpdate(update).
		Build()

	res, err := d.Service.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(tablePrefix + tableName),
		Key: map[string]*dynamodb.AttributeValue{
			keyNameRedirectsTable: {
				S: aws.String(name),
			},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              aws.String("UPDATED_NEW"),
	})
	if err != nil {
		return err
	}
	if len(res.Attributes) == 0 {
		return errors.New(ErrRedirectUpdate)
	}
	if len(res.Attributes) != len(expr.Values()) {
		return errors.New(ErrRedirectUpdatePartly)
	}

	return nil
}

// GetRedirectsPaginated returns a limited set of redirect entries and a cursor to use if there are more items to read
func (d *Database) GetRedirectsPaginated(cursor *string) ([]Redirect, *string, error) {
	var startkey map[string]*dynamodb.AttributeValue
	var newCursor string
	tablePrefix := d.Config.TableNamePrefix
	tableName := d.Config.TableRedirects

	if cursor != nil {
		encodedLastKey, err := url.QueryUnescape(*cursor)
		if err != nil {
			return nil, nil, errors.WithMessage(err, ErrCursorDecode)
		}
		lastKey, err := base64.StdEncoding.DecodeString(encodedLastKey)
		if err != nil {
			return nil, nil, errors.WithMessage(err, ErrCursorDecode)
		}
		startkey = map[string]*dynamodb.AttributeValue{
			keyNameRedirectsTable: {
				S: aws.String(string(lastKey)),
			},
		}
	}

	res, err := d.Service.Scan(&dynamodb.ScanInput{
		TableName:         aws.String(tablePrefix + tableName),
		Limit:             aws.Int64(25),
		ExclusiveStartKey: startkey,
	})
	if err != nil {
		return nil, nil, errors.WithMessage(err, ErrRedirectsScan)
	}

	redirects := []Redirect{}
	err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &redirects)
	if err != nil {
		return nil, nil, errors.WithMessage(err, ErrRedirectsUnmarshal)
	}

	val, ok := res.LastEvaluatedKey[keyNameRedirectsTable]
	if ok {
		newCursor = base64.StdEncoding.EncodeToString([]byte(*val.S))
		newCursor = url.QueryEscape(newCursor)
	} else {
		newCursor = "EOF"
	}

	return redirects, &newCursor, nil
}

// ExportRedirects paginates through all redirects entries and returns them
func (d *Database) ExportRedirects() ([]Redirect, error) {
	var cursor *string
	redirects := []Redirect{}

	for cursor == nil || (cursor != nil && *cursor != "EOF") {
		redirectsX, newCursor, err := d.GetRedirectsPaginated(cursor)
		cursor = newCursor
		if err != nil {
			return nil, errors.WithMessage(err, ErrRedirectsFetch)
		}
		if redirectsX == nil {
			return nil, errors.New(ErrRedirectListEmpty)
		}
		redirects = append(redirects, redirectsX...)
	}

	return redirects, nil
}

// ImportRedirects imports a set of redirect entries
// TODO: add limit so write capacity isn't exceeded
func (d *Database) ImportRedirects(redirects []Redirect) error {
	tablePrefix := d.Config.TableNamePrefix
	tableName := d.Config.TableRedirects

	for _, redirect := range redirects {
		redirect.Created = int(time.Now().Unix())
		redirect.Modified = redirect.Created

		payload, err := dynamodbattribute.MarshalMap(redirect)
		if err != nil {
			return errors.WithMessage(err, ErrRedirectMarshal)
		}

		_, err = d.Service.PutItem(&dynamodb.PutItemInput{
			TableName: aws.String(tablePrefix + tableName),
			Item:      payload,
		})

		return err
	}

	return nil
}
