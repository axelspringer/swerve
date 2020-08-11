package config

import (
	"github.com/axelspringer/swerve/api"
	"github.com/axelspringer/swerve/database"
)

// Configuration contains the config for the app
type Configuration struct {
	Database          database.Config
	API               api.Config
	LetsencryptUrl    string
	PebbleCA          string
	PebbleCAUrl       string
	UsePebble         bool
	UseStage          bool
	HTTPListenerPort  int
	HTTPSListenerPort int
	LogLevel          string
	LogFormatter      string
	Prod              bool
	Bootstrap         bool
	CacheInterval     int
}
