package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/axelspringer/swerve/log"
	"github.com/pkg/errors"
)

// NewDatabase creates a new instance
func NewDatabase(c Config) (*Database, error) {
	db := &Database{}

	dynamoConfig := &aws.Config{
		Region: aws.String(c.Region),
	}

	if c.Endpoint != "" {
		dynamoConfig.Endpoint = aws.String(c.Endpoint)
	}

	if c.Key != "" && c.Secret != "" {
		dynamoConfig.Credentials = credentials.NewStaticCredentials(c.Key, c.Secret, "")
	}

	sess, err := session.NewSession(dynamoConfig)
	if err != nil {
		return db, errors.WithMessage(err, ErrSessionCreate)
	}

	db.Service = dynamodb.New(sess)
	db.Config = c

	return db, nil
}

// Prepare prepares the database
func (d *Database) Prepare() error {
	log.Debug("call Prepare() dynamodb")
	err := d.prepareTable(d.Config.TableRedirects, keyNameRedirectsTable)
	if err != nil {
		return errors.WithMessagef(err, ErrfTableCreate, d.Config.TableRedirects)
	}
	err = d.prepareTable(d.Config.TableCertCache, keyNameCertCacheTable)
	if err != nil {
		return errors.WithMessagef(err, ErrfTableCreate, d.Config.TableCertCache)
	}
	err = d.prepareTable(d.Config.TableUsers, keyNameUsersTable)
	if err != nil {
		return errors.WithMessagef(err, ErrfTableCreate, d.Config.TableUsers)
	}
	_, err = d.Service.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(d.Config.TableUsers),
		Item: map[string]*dynamodb.AttributeValue{
			keyNameUsersTable: {
				S: aws.String(defaultDynamoUser),
			},
			attrNamePwd: {
				S: aws.String(defaultDynamoPassword),
			},
		},
	})
	if err != nil {
		return errors.WithMessagef(err, "can't create default user", d.Config.TableUsers)
	}

	return nil
}

func (d *Database) prepareTable(tableName string, keyName string) error {
	log.Debugf("try to create table %s", tableName)
	tablePrefix := d.Config.TableNamePrefix

	tableCreate := &dynamodb.CreateTableInput{
		TableName: aws.String(tablePrefix + tableName),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String(keyName), KeyType: aws.String("HASH")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String(keyName), AttributeType: aws.String("S")},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}
	tableDescribe := &dynamodb.DescribeTableInput{
		TableName: aws.String(tablePrefix + tableName),
	}

	if _, err := d.Service.DescribeTable(tableDescribe); err != nil {
		log.Warn(errors.WithMessagef(err, "Table '%s' does not exist", tableName).Error())
		if _, cerr := d.Service.CreateTable(tableCreate); cerr != nil {
			return cerr
		}
		log.Infof("Table '%s' created", tableName)
	}
	return nil
}
