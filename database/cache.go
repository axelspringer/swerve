package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/pkg/errors"
)

// UpdateCacheEntry puts a new certificate into the database
func (d *Database) UpdateCacheEntry(key string, data []byte) error {
	tablePrefix := d.Config.TableNamePrefix
	tableName := d.Config.TableCertCache

	if _, err := d.Service.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tablePrefix + tableName),
		Item: map[string]*dynamodb.AttributeValue{
			keyNameCertCacheTable: {
				S: aws.String(key),
			},
			attrNameCacheValue: {
				S: aws.String(string(data)),
			},
		},
	}); err != nil {
		return errors.WithMessage(err, ErrCertCacheUpdate)
	}

	return nil
}

// GetCacheEntry returns the certificate cache entry corresponding to the key
func (d *Database) GetCacheEntry(key string) ([]byte, error) {
	tablePrefix := d.Config.TableNamePrefix
	tableName := d.Config.TableCertCache

	res, err := d.Service.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tablePrefix + tableName),
		Key: map[string]*dynamodb.AttributeValue{
			keyNameCertCacheTable: {
				S: aws.String(key),
			},
		},
	})
	if err != nil {
		return nil, errors.WithMessage(err, ErrCertCacheFetch)
	}

	certCache := &CertCacheEntry{}
	if err := dynamodbattribute.UnmarshalMap(res.Item, &certCache); err != nil {
		return nil, errors.WithMessage(err, ErrCertCacheMarshal)
	}

	return []byte(certCache.Value), nil
}

// DeleteCacheEntry deletes a certificate cache entry
func (d *Database) DeleteCacheEntry(key string) error {
	tablePrefix := d.Config.TableNamePrefix
	tableName := d.Config.TableCertCache

	if _, err := d.Service.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(tablePrefix + tableName),
		Key: map[string]*dynamodb.AttributeValue{
			keyNameCertCacheTable: {
				S: aws.String(key),
			},
		},
	}); err != nil {
		return errors.WithMessage(err, ErrCertCacheDelete)
	}

	return nil
}
