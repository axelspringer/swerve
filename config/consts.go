package config

const (
	envVarPrefix = "SWERVE_"
)

const (
	envStrAPIListenerPort   = "API_LISTENER"
	envStrAPIVersion        = "API_VERSION"
	envStrAPIUIURL          = "API_UI_URL"
	envStrAPIJWTSecret      = "API_JWT_SECRET"
	envStrHTTPListenerPort  = "HTTP_LISTENER"
	envStrHTTPSListenerPort = "HTTPS_LISTENER"
	envStrLogLevel          = "LOG_LEVEL"
	envStrLogFormatter      = "LOG_FORMATTER"
	envStrProd              = "PROD"
	envStrBootstrap         = "BOOTSTRAP"
	envStrCacheInterval     = "CACHE_INTERVAL"
	envStrTableNamePrefix   = "TABLE_PREFIX"
	envStrDBRegion          = "DB_REGION"
	envStrTableRedirects    = "TABLE_REDIRECTS"
	envStrTableCertCache    = "TABLE_CERTCACHE"
	envStrTableUsers        = "TABLE_USERS"
	envStrDBKey             = "DB_KEY"
	envStrDBSecret          = "DB_SECRET"
	envStrDBEndpoint        = "DB_ENDPOINT"
)

const (
	paramStrAPIListenerPort   = "api-listener"
	paramStrAPIVersion        = "api-version"
	paramStrAPIUIURL          = "api-ui-url"
	paramStrAPIJWTSecret      = "api-jwt-secret"
	paramStrHTTPListenerPort  = "http-listener"
	paramStrHTTPSListenerPort = "https-listener"
	paramStrLogLevel          = "log-level"
	paramStrLogFormatter      = "log-formatter"
	paramStrProd              = "prod"
	paramStrBootstrap         = "bootstrap"
	paramStrCacheInterval     = "cache-interval"
	paramStrTableNamePrefix   = "table-prefix"
	paramStrDBRegion          = "db-region"
	paramStrTableRedirects    = "table-redirects"
	paramStrTableCertCache    = "table-certcache"
	paramStrTableUsers        = "table-users"
	paramStrDBKey             = "db-key"
	paramStrDBSecret          = "db-secret"
	paramStrDBEndpoint        = "db-endpoint"
)

const (
	defaultAPIListener     = 8082
	defaultHTTPListener    = 8080
	defaultHTTPSListener   = 8081
	defaultLogLevel        = "debug"
	defaultLogFormatter    = "text"
	defaultProd            = false
	defaultBootstrap       = false
	defaultCacheInterval   = 10
	defaultTableNamePrefix = "Swerve"
	defaultDBRegion        = "eu-west-1"
	defaultTableRedirects  = "Redirects"
	defaultTableCertCache  = "CertCache"
	defaultTableUsers      = "Users"
)
