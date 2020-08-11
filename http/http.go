package http

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/axelspringer/swerve/helper"
	"github.com/axelspringer/swerve/log"
)

// NewHTTPServer creates a new instance
func NewHTTPServer(getRedirect GetRedirect,
	acmHandler ACMHandler,
	listener int,
	wrapHandler func(string, http.Handler) http.Handler) *HTTP {
	server := &HTTP{
		ACMHandler:  acmHandler,
		getRedirect: getRedirect,
		listener:    listener,
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", wrapHandler("HTTP", helper.LoggingMiddleware(server.handler())))

	addr := ":" + strconv.Itoa(listener)
	server.server = &http.Server{
		Addr:    addr,
		Handler: mux,
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

	// regular domain lookup
	if err == nil {
		redirectURL, redirectCode := redirect.GetRedirect(r.URL)
		http.Redirect(w, r, redirectURL, redirectCode)
		// log.Response(r, redirectCode)
		return
	}

	http.NotFound(w, r)
	// log.Response(r, http.StatusNotFound)
}

func (h *HTTP) serve(fallback http.Handler, w http.ResponseWriter, r *http.Request) {
	h.ACMHandler(fallback).ServeHTTP(w, r)
}
