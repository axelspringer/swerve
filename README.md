# Evade
A scalable redirecting service with integrated path mapping functionality and automatic certificate issuing and renewal using a persistent certificate cache
## Step 1

### Configuration
Is pulled form variables
* API_LISTENER - The port used for the API server
* API_VERSION 		- The api version
* API_UI_URL 		- The URL of the UI for the API (import for CORS)
* API_JWT_SECRET	- The secret used to sign the JWT for authentication
* HTTP_LISTENER 	- The port used for the redirecting server (http)
* HTTPS_LISTENER 	- The port used for the redirecting server (https)
* LOG_LEVEL 		- Logrus standard log level (info,debug,warning,error,fatal,panic)
* LOG_FORMATTER 	- Logrus log formatter, accepts "json" or "text"
* PRODUCTION 		- Do you want to run your app in production mode?
* BOOTSTRAP 		- Should the tables be created on init?
* CACHE_INTERVAL 	- Interval in minutes in which the cache is updated
* TABLE_PREFIX 		- The prefix used for the DynamoDB table names
* DB_REGION 		- The AWS Region the DynamoDB tables are located in
* TABLE_REDIRECTS 	- The table name of the table storing the redirect records
* TABLE_CERTCACHE 	- The table name of the table storing the certificates
* TABLE_USERS 		- The table name of the table storing the user credentials
* DB_KEY 			- AWS credentials for database access 
* DB_SECRET 		- AWS credentials for database access 
* DB_ENDPOINT 		- DynamoDB endpoint URL

#### Default configuration
```
APIListener     = 8082
HTTPListener    = 8080
HTTPSListener   = 8081
LogLevel        = "debug"
LogFormatter    = "text"
Prod            = false
Bootstrap		= false
CacheInterval	= 10
TableNamePrefix = "Evade"
DBRegion        = "eu-west-1"
TableRedirects  = "Domains"
TableCertCache  = "CertCache"
```
## Step 2
Run evade and insert your redirect records via API ([API doc](https://app.swaggerhub.com/apis-docs/TetsuyaXD/evade/1.0.0))
## Step 3
Let your domains point to the HTTP/HTTPS Server