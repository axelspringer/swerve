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
	envStrLetsEncryptURL    = "LETSENCRYPT_URL"
	envStrUsePebble         = "USE_PEBBLE"
	envStrPebbleCAURL       = "PEBBLE_CA_URL"
	envStrUseStage          = "USE_STAGE"
	envStrLogLevel          = "LOG_LEVEL"
	envStrLogFormatter      = "LOG_FORMATTER"
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
	paramStrLetsEncryptURL    = "letsencrypt-url"
	paramStrUsePebble         = "use-pebble"
	paramStrPebbleCAURL       = "pebble-ca-url"
	paramStrUseStage          = "use-stage"
	paramStrLogLevel          = "log-level"
	paramStrLogFormatter      = "log-formatter"
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
	defaultBootstrap       = false
	defaultCacheInterval   = 10
	defaultTableNamePrefix = ""
	defaultDBRegion        = "eu-west-1"
	defaultTableRedirects  = "Swerve_Redirects"
	defaultTableCertCache  = "Swerve_CertCache"
	defaultTableUsers      = "Swerve_Users"
	defaultPebbleCACert    = `-----BEGIN CERTIFICATE-----
MIIDCTCCAfGgAwIBAgIIJOLbes8sTr4wDQYJKoZIhvcNAQELBQAwIDEeMBwGA1UE
AxMVbWluaWNhIHJvb3QgY2EgMjRlMmRiMCAXDTE3MTIwNjE5NDIxMFoYDzIxMTcx
MjA2MTk0MjEwWjAgMR4wHAYDVQQDExVtaW5pY2Egcm9vdCBjYSAyNGUyZGIwggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC5WgZNoVJandj43kkLyU50vzCZ
alozvdRo3OFiKoDtmqKPNWRNO2hC9AUNxTDJco51Yc42u/WV3fPbbhSznTiOOVtn
Ajm6iq4I5nZYltGGZetGDOQWr78y2gWY+SG078MuOO2hyDIiKtVc3xiXYA+8Hluu
9F8KbqSS1h55yxZ9b87eKR+B0zu2ahzBCIHKmKWgc6N13l7aDxxY3D6uq8gtJRU0
toumyLbdzGcupVvjbjDP11nl07RESDWBLG1/g3ktJvqIa4BWgU2HMh4rND6y8OD3
Hy3H8MY6CElL+MOCbFJjWqhtOxeFyZZV9q3kYnk9CAuQJKMEGuN4GU6tzhW1AgMB
AAGjRTBDMA4GA1UdDwEB/wQEAwIChDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYB
BQUHAwIwEgYDVR0TAQH/BAgwBgEB/wIBADANBgkqhkiG9w0BAQsFAAOCAQEAF85v
d40HK1ouDAtWeO1PbnWfGEmC5Xa478s9ddOd9Clvp2McYzNlAFfM7kdcj6xeiNhF
WPIfaGAi/QdURSL/6C1KsVDqlFBlTs9zYfh2g0UXGvJtj1maeih7zxFLvet+fqll
xseM4P9EVJaQxwuK/F78YBt0tCNfivC6JNZMgxKF59h0FBpH70ytUSHXdz7FKwix
Mfn3qEb9BXSk0Q3prNV5sOV3vgjEtB4THfDxSz9z3+DepVnW3vbbqwEbkXdk3j82
2muVldgOUgTwK8eT+XdofVdntzU/kzygSAtAQwLJfn51fS1GvEcYGBc1bDryIqmF
p9BI7gVKtWSZYegicA==
-----END CERTIFICATE-----`
)
