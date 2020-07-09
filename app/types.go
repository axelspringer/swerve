package app

import (
	"github.com/axelspringer/swerve/api"
	"github.com/axelspringer/swerve/cache"
	"github.com/axelspringer/swerve/config"
	"github.com/axelspringer/swerve/http"
	"github.com/axelspringer/swerve/https"
)

// Application is swerves app model
type Application struct {
	Config      *config.Configuration
	APIServer   *api.API
	HTTPServer  *http.HTTP
	HTTPSServer *https.HTTPS
	Cache       *cache.Cache
}
