package config

import (
	"flag"

	"github.com/axelspringer/swerve/log"
)

// FromParameter reads the configuration from application parameters
func (c *Configuration) FromParameter() {
	params := make(map[string]interface{})
	apiListenerPort := flag.Int(paramStrAPIListenerPort, 0, "Set the API listener port")
	apiVersion := flag.String(paramStrAPIVersion, "", "Set the used API version")
	apiUIURL := flag.String(paramStrAPIUIURL, "", "Set the API Cross Origin Resource UI URL")
	apiJWTSecret := flag.String(paramStrAPIJWTSecret, "", "Set the API JWT secret")
	httpListenerPort := flag.Int(paramStrHTTPListenerPort, 0, "Set the http listener port")
	httpsListenerPort := flag.Int(paramStrHTTPSListenerPort, 0, "Set the https listener port")

	usePebble := flag.Bool(paramStrUsePebble, false, "Use pebble?")
	pebbleCAURL := flag.String(paramStrPebbleCAURL, "", "Set the the pebble CA URL")
	useStage := flag.Bool(paramStrUseStage, false, "Use the LetsEncrypt stage URL?")
	letsEncryptURL := flag.String(paramStrLetsEncryptURL, "", "Set the LetsEncrypt URL")

	logLevel := flag.String(paramStrLogLevel, "", "Set the log level (info,debug,warning,error,fatal,panic)")
	logFormatter := flag.String(paramStrLogFormatter, "", "Set the log formatter (text,json)")

	boostrap := flag.Bool(paramStrBootstrap, false, "Is bootstrap?")
	cacheInterval := flag.Int(paramStrCacheInterval, 0, "Set cache interval in minutes")

	tableNamePrefix := flag.String(paramStrTableNamePrefix, "", "DynamoDB table name prefix")
	dbRegion := flag.String(paramStrDBRegion, "", "AWS region of the database")
	tableRedirects := flag.String(paramStrTableRedirects, "", "Table name of the redirect table")
	tableCertCache := flag.String(paramStrTableCertCache, "", "Table name of the cert cache table")
	tableUsers := flag.String(paramStrTableUsers, "", "Table name of the user table")
	dbKey := flag.String(paramStrDBKey, "", "AWS Key")
	dbSecret := flag.String(paramStrDBSecret, "", "AWS Secret")
	dbEndpoint := flag.String(paramStrDBEndpoint, "", "DynamoDB endpoint")

	flag.Parse()

	if apiListenerPort != nil && *apiListenerPort != 0 {
		params[paramStrAPIListenerPort] = *apiListenerPort
		c.API.Listener = *apiListenerPort
	}

	if apiVersion != nil && *apiVersion != "" {
		params[paramStrAPIVersion] = *apiVersion
		c.API.Version = *apiVersion
	}

	if apiUIURL != nil && *apiUIURL != "" {
		params[paramStrAPIUIURL] = *apiUIURL
		c.API.COR = *apiUIURL
	}

	if apiJWTSecret != nil && *apiJWTSecret != "" {
		params[paramStrAPIJWTSecret] = *apiJWTSecret
		c.API.COR = *apiJWTSecret
	}

	if httpListenerPort != nil && *httpListenerPort != 0 {
		params[paramStrHTTPListenerPort] = *httpListenerPort
		c.HTTPListenerPort = *httpListenerPort
	}

	if httpsListenerPort != nil && *httpsListenerPort != 0 {
		params[paramStrHTTPSListenerPort] = *httpsListenerPort
		c.HTTPSListenerPort = *httpsListenerPort
	}

	if usePebble != nil && *usePebble {
		params[paramStrUsePebble] = *usePebble
		c.ACM.UsePebble = *usePebble
	}

	if pebbleCAURL != nil && *pebbleCAURL != "" {
		params[paramStrPebbleCAURL] = *pebbleCAURL
		c.ACM.PebbleCAURL = *pebbleCAURL
	}

	if useStage != nil && *useStage {
		params[paramStrUseStage] = *useStage
		c.ACM.UseStage = *useStage
	}

	if letsEncryptURL != nil && *letsEncryptURL != "" {
		params[paramStrLetsEncryptURL] = *letsEncryptURL
		c.ACM.LetsEncryptURL = *letsEncryptURL
	}

	if logLevel != nil && *logLevel != "" {
		params[paramStrLogLevel] = *logLevel
		c.LogLevel = *logLevel
	}

	if logFormatter != nil && *logFormatter != "" {
		params[paramStrLogFormatter] = *logFormatter
		c.LogFormatter = *logFormatter
	}

	if boostrap != nil && *boostrap {
		params[paramStrBootstrap] = true
		c.Bootstrap = true
	}

	if cacheInterval != nil && *cacheInterval != 0 {
		params[paramStrCacheInterval] = *cacheInterval
		c.CacheInterval = *cacheInterval
	}

	if tableNamePrefix != nil && *tableNamePrefix != "" {
		params[paramStrTableNamePrefix] = *tableNamePrefix
		c.Database.TableNamePrefix = *tableNamePrefix
	}

	if dbRegion != nil && *dbRegion != "" {
		params[paramStrDBRegion] = *dbRegion
		c.Database.Region = *dbRegion
	}

	if tableRedirects != nil && *tableRedirects != "" {
		params[paramStrTableRedirects] = *tableRedirects
		c.Database.TableRedirects = *tableRedirects
	}

	if tableCertCache != nil && *tableCertCache != "" {
		params[paramStrTableCertCache] = *tableCertCache
		c.Database.TableCertCache = *tableCertCache
	}

	if tableUsers != nil && *tableUsers != "" {
		params[paramStrTableUsers] = *tableUsers
		c.Database.TableUsers = *tableUsers
	}

	if dbKey != nil && *dbKey != "" {
		if dbSecret != nil && *dbSecret != "" {
			params[paramStrDBKey] = *dbKey
			params[paramStrDBSecret] = *dbSecret
			c.Database.Key = *dbKey
			c.Database.Secret = *dbSecret
		}
	}

	if dbEndpoint != nil && *dbEndpoint != "" {
		params[paramStrDBEndpoint] = *dbEndpoint
		c.Database.Endpoint = *dbEndpoint
	}

	log.Infof("Loading config from parameters: %+v", params)
}
