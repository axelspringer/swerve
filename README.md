# SWERVE

Swerve is a redirection service that uses autocert to generate https certificates automatically. The domain and certificate data are stored in a DynamoDB

## Test

    make test/local

## Build

    make

or

    make build/local

### Build docker container

    make build/docker

## Parameter

### Environment parameter

* SWERVE_DB_ENDPOINT - AWS endpoint for the DynamoDB
* SWERVE_DB_REGION - AWS region for the DynamoDB
* SWERVE_DB_KEY - AWS key for credential based access
* SWERVE_DB_SECRET - AWS secret for credential based access
* SWERVE_API - Address for the API listener
* SWERVE_HTTP - Address for the HTTP listener
* SWERVE_HTTPS - Address for the HTTPS listener
* BOOTSTRAP - DB table preparation
* LOG_LEVEL - Log level info, debug, warning, error, fatal and panic

### Application parameter

* db-endpoint - AWS endpoint for the DynamoDB
* db-region - AWS region for the DynamoDB
* db-key - AWS key for credential based access
* db-secret - AWS secret for credential based access
* bootstrap - DB table preparation
* api - Address for the API listener
* http - Address for the HTTP listener
* https - Address for the HTTPS listener

## Run

### run local dynamodb

    make run/dynamo

### run swerve (examples)

    SWERVE_DB_ENDPOINT=http://localhost:8000 SWERVE_DB_REGION=eu-west-1 SWERVE_HTTPS=:8081 ./bin/swerve

## API

### Domain model

    {
        "id": "guid v4 will be generated",
        "domain": "my.domain.com",
        "path": [
            {
                "from": "/match/path/prefix",
                "to": "/foo"
            },
            {
                "from": "/other/target",
                "to": "https://the.other.one/"
            }
        ],
        "redirect": "https://my.redirect.com"
        "promotable": false,
        "code": 301,
        "description": "Meanful description of this redirection",
        "created": "generated date",
        "modified": "generated date"
    }

#### id

Will be generated

#### domain

The domain name to keep track on. e.g. ```my.redirect.com```

#### path

You can add an aditional path mapping conditional list. When defined the redirection based on the matching result of this list. Fallback is the default redirect

#### redirect

Redirection target

#### promotable

Promotable redirects are attaching the path of the request to the redirection location e.g.
with ```"promotable": true``` my.domain.com/foo/bar/baz.html will be redirected to https://my.redirect.com/foo/bar/baz.html
with ```"promotable": false``` my.domain.com/foo/bar/baz.html will be redirected to https://my.redirect.com

#### code

The redirection code. It has to be greater or equal 300 and less or equal than 399

#### description

Meanful description of the domain entry

## API calls

### Get all domains

    curl -X GET http://localhost:8082/domain

### Get a single domain by name

    curl -X GET http://localhost:8082/domain/<name>

### Insert a new domain

    curl -X POST \
        http://localhost:8082/domain/ \
        -d '{
            "domain": "<name>",
            "redirect": "https://my.redirect.target.com",
            "code": 308,
            "description": "Example domain entry"
        }'

### Purge a domain by name

    curl -X DELETE http://localhost:8082/domain/<name>

### Export all domains

    curl -X GET http://localhost:8082/export

### Import a export set

    curl -X POST \
    http://localhost:8082/import \
    -H 'cache-control: no-cache' \
    -H 'content-type: application/json' \
    -d '{
        "domains": [
            {
                "id": "6a368e16-ea73-40e0-a4cd-e12148c15ffd",
                "domain": "my.domain.com",
                "redirect": "https://example.com",
                "promotable": false,
                "wildcard": false,
                "certificate": "",
                "code": 308,
                "description": "example registration 2",
                "created": "2018-08-28T14:31:06+02:00",
                "modified": ""
            }
        ]
    }'