// Copyright 2018 Axel Springer SE
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDB model
type DynamoDB struct {
	Session *session.Session
	Service *dynamodb.DynamoDB
}

// DynamoConnection model
type DynamoConnection struct {
	Endpoint  string
	Key       string
	Secret    string
	TableName string
	Region    string
}

// DomainList db entry
type DomainList struct {
	Domains []Domain `json:"domains"`
}

// PathMappingEntry model
type PathMappingEntry struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// PathList model
type PathList []PathMappingEntry

// TLSCacheEntry model
type TLSCacheEntry struct {
	Key   string `json:"cacheKey"`
	Value string `json:"cacheValue"`
}

// Domain struct as it is received via the request body entry
type Domain struct {
	ID           string    `json:"id"`
	Name         string    `json:"domain"`
	PathMapping  *PathList `json:"paths"`
	Redirect     string    `json:"redirect"`
	Promotable   bool      `json:"promotable"`
	Wildcard     bool      `json:"wildcard"`
	Certificate  string    `json:"certificate"`
	RedirectCode int       `json:"code"`
	Description  string    `json:"description"`
	Created      string    `json:"created"`
	Modified     string    `json:"modified"`
}

// DomainDB entry
type DomainDB struct {
	Domain
	BinPathMapping *[]byte `json:"bin_paths"`
}

// ExportDomains model
type ExportDomains struct {
	Domains []Domain `json:"domains"`
}

// User model
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
