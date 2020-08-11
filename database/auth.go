package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/pkg/errors"
)

// GetPwdHash returns a password hash from the database
func (d *Database) GetPwdHash(username string) (string, error) {
	tablePrefix := d.Config.TableNamePrefix
	tableName := d.Config.TableUsers
	res, err := d.Service.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tablePrefix + tableName),
		Key: map[string]*dynamodb.AttributeValue{
			keyNameUsersTable: {
				S: aws.String(username),
			},
		},
	})
	if err != nil {
		return "", errors.WithMessage(err, ErrRedirectsFetch)
	}

	if len(res.Item) == 0 {
		return "", errors.New(ErrUserNotFound)
	}

	userRes := &User{}
	if err = dynamodbattribute.UnmarshalMap(res.Item, &userRes); err != nil {
		return "", errors.WithMessage(err, ErrRedirectMarshal)
	}

	return userRes.Pwd, nil
}
