package http

import (
	"net/http"
	"strconv"

	"github.com/axelspringer/swerve/helper"
	"github.com/axelspringer/swerve/log"
)

// NewHTTPServer creates a new instance
func NewHTTPServer(getRedirect GetRedirect,
	acmHandler ACMHandler,
	listener int) *HTTP {
	server := &HTTP{
		ACMHandler:  acmHandler,
		getRedirect: getRedirect,
		listener:    listener,
	}

	addr := ":" + strconv.Itoa(listener)
	server.server = &http.Server{
		Addr:    addr,
		Handler: helper.LoggingMiddleware(server.handler()),
	}

	return server
}

// Listen to http
func (h *HTTP) Listen() error {
	log.Infof("HTTP listening to %d", h.listener)
	return h.server.ListenAndServe()
}

func (h *HTTP) handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.serve(http.HandlerFunc(h.handleRedirect), w, r)
	})
}

func (h *HTTP) handleRedirect(w http.ResponseWriter, r *http.Request) {
	hostHeader := r.Host
	redirect, err := h.getRedirect(hostHeader)
	msg := "Response with status code %d"

	// regular domain lookup
	if err == nil {
		redirectURL, redirectCode := redirect.GetRedirect(r.URL)
		http.Redirect(w, r, redirectURL, redirectCode)
		log.Infof(msg, redirectCode)
		return
	}

	log.Infof(msg, http.StatusNotFound)
	http.NotFound(w, r)
}

func (h *HTTP) serve(fallback http.Handler, w http.ResponseWriter, r *http.Request) {
	h.ACMHandler(fallback).ServeHTTP(w, r)
}
