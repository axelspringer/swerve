package http

import (
	"net/http"

	"github.com/axelspringer/swerve/model"
)

// HTTP is an HTTP server
type HTTP struct {
	ACMHandler  ACMHandler
	getRedirect func(name string) (model.Redirect, error)
	server      *http.Server
	listener    int
}

// ACMHandler HTTPHandler the autocert Manager for let's encrypt challenge
type ACMHandler func(fallback http.Handler) http.Handler

// GetRedirect function to retrieve redirect entry
type GetRedirect func(name string) (model.Redirect, error)
