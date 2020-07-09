package https

import (
	"crypto/tls"
	"net/http"
	"strconv"

	"github.com/axelspringer/swerve/helper"

	"github.com/axelspringer/swerve/log"
)

// NewHTTPSServer creates a new instance
func NewHTTPSServer(getRedirect GetRedirect,
	getCertificate GetCertificate,
	listener int) *HTTPS {
	server := &HTTPS{
		getRedirect: getRedirect,
		listener:    listener,
	}

	addr := ":" + strconv.Itoa(listener)
	server.server = &http.Server{
		Addr: addr,
		TLSConfig: &tls.Config{
			GetCertificate: getCertificate,
		},
		Handler: helper.LoggingMiddleware(server.handler()),
	}

	return server
}

// Listen to https
func (h *HTTPS) Listen() error {
	log.Infof("HTTPS listening to %d", h.listener)
	return h.server.ListenAndServeTLS("", "")
}

func (h *HTTPS) handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})
}
