package config

import (
	"github.com/axelspringer/swerve/acm"
	"github.com/axelspringer/swerve/api"
	"github.com/axelspringer/swerve/database"
)

// Configuration contains the config for the app
type Configuration struct {
	Database          database.Config
	API               api.Config
	ACM               acm.Config
	HTTPListenerPort  int
	HTTPSListenerPort int
	LogLevel          string
	LogFormatter      string
	Bootstrap         bool
	CacheInterval     int
}
