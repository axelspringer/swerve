package database

import (
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	d *Database
)

func TestMain(m *testing.M) {
	var err error
	cfg := Config{
		TableNamePrefix: "",
		Region:          "eu-west-1",
		TableRedirects:  "Redirects",
		TableCertCache:  "DomainsTLSCache",
		TableUsers:      "Users",
		Key:             "0",
		Secret:          "0",
		Endpoint:        "http://localhost:8000",
	}
	d, err = NewDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func TestPutItem(t *testing.T) {
	_, err := d.Service.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(d.Config.TableNamePrefix + d.Config.TableUsers),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String("testuser"),
			},
			"password": {
				S: aws.String("c3ZvYnFvb2pic2R2bnN2ZG5zbnZ3Cg=="),
			},
		}})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserTable(t *testing.T) {
	var startkey map[string]*dynamodb.AttributeValue
	tablePrefix := d.Config.TableNamePrefix
	tableName := d.Config.TableUsers

	res, err := d.Service.Scan(&dynamodb.ScanInput{
		TableName:         aws.String(tablePrefix + tableName),
		Limit:             aws.Int64(25),
		ExclusiveStartKey: startkey,
	})
	if err != nil {
		t.Fatal(err)
	}

	users := []User{}
	err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &users)
	if err != nil {
		t.Fatal(err)
	}

	if len(users) == 0 {
		t.Fatal("No Users found in dynamodb")
	}

	for _, v := range users {
		t.Logf("Found User %s with PW %s", v.Name, v.Pwd)
	}

}
