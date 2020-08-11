package api

import (
	"net/http"
)

// API is an API server
type API struct {
	Model  ModelAdapter
	server *http.Server
	Config Config
}

// Config contains the API config
type Config struct {
	COR      string
	Version  string
	Listener int
	Secret   string
}
