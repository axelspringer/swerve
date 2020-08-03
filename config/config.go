package config

import (
	"github.com/axelspringer/swerve/acm"
	"github.com/axelspringer/swerve/api"
	"github.com/axelspringer/swerve/database"
)

// NewConfiguration creates a new instance
func NewConfiguration() *Configuration {
	return &Configuration{
		API: api.Config{
			Listener: defaultAPIListener,
		},
		ACM: acm.Config{
			PebbleCA: defaultPebbleCACert,
		},
		HTTPListenerPort:  defaultHTTPListener,
		HTTPSListenerPort: defaultHTTPSListener,
		LogLevel:          defaultLogLevel,
		LogFormatter:      defaultLogFormatter,
		Bootstrap:         defaultBootstrap,
		CacheInterval:     defaultCacheInterval,
		Database: database.Config{
			TableNamePrefix: defaultTableNamePrefix,
			Region:          defaultDBRegion,
			TableRedirects:  defaultTableRedirects,
			TableCertCache:  defaultTableCertCache,
			TableUsers:      defaultTableUsers,
		},
	}
}
