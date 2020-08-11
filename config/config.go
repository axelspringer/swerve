package config

import (
	"github.com/axelspringer/swerve/api"
	"github.com/axelspringer/swerve/database"
)

// NewConfiguration creates a new instance
func NewConfiguration() *Configuration {
	return &Configuration{
		API: api.Config{
			Listener: defaultAPIListener,
		},
		HTTPListenerPort:  defaultHTTPListener,
		HTTPSListenerPort: defaultHTTPSListener,
		LogLevel:          defaultLogLevel,
		LogFormatter:      defaultLogFormatter,
		Prod:              defaultProd,
		Bootstrap:         defaultBootstrap,
		CacheInterval:     defaultCacheInterval,
		PebbleCA:          defaultPebbleCACert,
		Database: database.Config{
			TableNamePrefix: defaultTableNamePrefix,
			Region:          defaultDBRegion,
			TableRedirects:  defaultTableRedirects,
			TableCertCache:  defaultTableCertCache,
			TableUsers:      defaultTableUsers,
		},
	}
}
