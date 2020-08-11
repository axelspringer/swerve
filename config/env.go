package config

import (
	"strconv"

	"github.com/pkg/errors"

	"github.com/axelspringer/swerve/log"
)

// FromEnv reads the configuration from the environment
func (c *Configuration) FromEnv() error {
	params := make(map[string]interface{})
	apiListenerPort := getPrefixedOSEnv(envStrAPIListenerPort)
	if apiListenerPort != "" {
		apiPortNumber, err := strconv.Atoi(apiListenerPort)
		if err != nil {
			return errors.New(ErrAPIPortInvalid)
		}
		params[envStrAPIListenerPort] = apiPortNumber
		c.API.Listener = apiPortNumber
	}

	usePebble := getPrefixedOSEnv(envStrUsePebble)
	if usePebble != "" {
		params[envStrUsePebble] = usePebble
		pebbleBoolean, err := strconv.ParseBool(usePebble)
		if err != nil {
			return errors.New(ErrPebbleValInvalid)
		}
		c.UsePebble = pebbleBoolean
	}

	pebbleCAUrl := getPrefixedOSEnv(envStrPebbleCAUrl)
	if pebbleCAUrl != "" {
		params[envStrPebbleCAUrl] = pebbleCAUrl
		c.PebbleCAUrl = pebbleCAUrl
	}

	useStage := getPrefixedOSEnv(envStrUseStage)
	if useStage != "" {
		params[envStrUseStage] = usePebble
		stageBoolean, err := strconv.ParseBool(useStage)
		if err != nil {
			return errors.New(ErrPebbleValInvalid)
		}
		c.UseStage = stageBoolean
	}

	letsencryptUrl := getPrefixedOSEnv(envStrLetsencryptUrl)
	if letsencryptUrl != "" {
		params[envStrLetsencryptUrl] = letsencryptUrl
		c.LetsencryptUrl = letsencryptUrl
	}

	apiVersion := getPrefixedOSEnv(envStrAPIVersion)
	if apiVersion != "" {
		params[envStrAPIVersion] = apiVersion
		c.API.Version = apiVersion
	}

	apiUIURL := getPrefixedOSEnv(envStrAPIUIURL)
	if apiUIURL != "" {
		params[envStrAPIUIURL] = apiUIURL
		c.API.COR = apiUIURL
	}

	apiJWTSecret := getPrefixedOSEnv(envStrAPIJWTSecret)
	if apiJWTSecret != "" {
		params[envStrAPIJWTSecret] = apiJWTSecret
		c.API.Secret = apiJWTSecret
	}

	httpListenerPort := getPrefixedOSEnv(envStrHTTPListenerPort)
	if httpListenerPort != "" {
		httpPortNumber, err := strconv.Atoi(httpListenerPort)
		if err != nil {
			return errors.New(ErrHTTPPortInvalid)
		}
		params[envStrHTTPListenerPort] = httpPortNumber
		c.HTTPListenerPort = httpPortNumber
	}

	httpsListenerPort := getPrefixedOSEnv(envStrHTTPSListenerPort)
	if httpsListenerPort != "" {
		httpsPortNumber, err := strconv.Atoi(httpsListenerPort)
		if err != nil {
			return errors.New(ErrHTTPSPortInvalid)
		}
		params[envStrHTTPSListenerPort] = httpsPortNumber
		c.HTTPSListenerPort = httpsPortNumber
	}

	logLevel := getPrefixedOSEnv(envStrLogLevel)
	if logLevel != "" {
		params[envStrLogLevel] = logLevel
		c.LogLevel = logLevel
	}

	logFormatter := getPrefixedOSEnv(envStrLogFormatter)
	if logFormatter != "" {
		params[envStrLogFormatter] = logFormatter
		c.LogFormatter = logFormatter
	}

	prod := getPrefixedOSEnv(envStrProd)
	if prod != "" {
		isProd, err := strconv.ParseBool(prod)
		if err != nil {
			return errors.New(ErrProdValInvalid)
		}
		params[envStrProd] = isProd
		c.Prod = isProd
	}

	bootstrap := getPrefixedOSEnv(envStrBootstrap)
	if bootstrap != "" {
		isBootstrap, err := strconv.ParseBool(bootstrap)
		if err != nil {
			return errors.New(ErrBoostrapValInvalid)
		}
		params[envStrBootstrap] = isBootstrap
		c.Bootstrap = isBootstrap
	}

	cacheInterval := getPrefixedOSEnv(envStrCacheInterval)
	if cacheInterval != "" {
		cacheIntervalNumber, err := strconv.Atoi(cacheInterval)
		if err != nil {
			return errors.New(ErrCacheIntervalInvalid)
		}
		params[envStrCacheInterval] = cacheIntervalNumber
		c.CacheInterval = cacheIntervalNumber
	}

	tableNamePrefix := getPrefixedOSEnv(envStrTableNamePrefix)
	if tableNamePrefix != "" {
		params[envStrTableNamePrefix] = tableNamePrefix
		c.Database.TableNamePrefix = tableNamePrefix
	}

	dbRegion := getPrefixedOSEnv(envStrDBRegion)
	if dbRegion != "" {
		params[envStrDBRegion] = dbRegion
		c.Database.Region = dbRegion
	}

	tableRedirects := getPrefixedOSEnv(envStrTableRedirects)
	if tableRedirects != "" {
		params[envStrTableRedirects] = tableRedirects
		c.Database.TableRedirects = tableRedirects
	}

	tableCertCache := getPrefixedOSEnv(envStrTableCertCache)
	if tableCertCache != "" {
		params[envStrTableCertCache] = tableCertCache
		c.Database.TableCertCache = tableCertCache
	}

	tableUsers := getPrefixedOSEnv(envStrTableUsers)
	if tableUsers != "" {
		params[envStrTableUsers] = tableUsers
		c.Database.TableUsers = tableUsers
	}

	dbKey := getPrefixedOSEnv(envStrDBKey)
	dbSecret := getPrefixedOSEnv(envStrDBSecret)
	if dbKey != "" {
		if dbSecret != "" {
			params[envStrDBKey] = dbKey
			params[envStrDBSecret] = dbSecret
			c.Database.Key = dbKey
			c.Database.Secret = dbSecret
		}
	}

	dbEndpoint := getPrefixedOSEnv(envStrDBEndpoint)
	if dbEndpoint != "" {
		params[envStrDBEndpoint] = dbEndpoint
		c.Database.Endpoint = dbEndpoint
	}

	log.Infof("Loading config from environment: %+v", params)

	return nil
}
