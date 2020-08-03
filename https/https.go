package https

import (
	"crypto/tls"
	"net/http"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/axelspringer/swerve/helper"

	"github.com/axelspringer/swerve/log"
)

// NewHTTPSServer creates a new instance
func NewHTTPSServer(getRedirect GetRedirect,
	getCertificate GetCertificate,
	listener int,
	wrapHandler func(string, http.Handler) http.Handler) *HTTPS {
	server := &HTTPS{
		getRedirect: getRedirect,
		listener:    listener,
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", wrapHandler("HTTPS", helper.LoggingMiddleware(server.handler())))

	addr := ":" + strconv.Itoa(listener)
	server.server = &http.Server{
		Addr: addr,
		TLSConfig: &tls.Config{
			GetCertificate: getCertificate,
		},
		Handler: mux,
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
		redirect, err := h.getRedirect(strings.Split(r.Host, ":")[0])

		// regular domain lookup
		if err == nil {
			redirectURL, redirectCode := redirect.GetRedirect(r.URL)
			http.Redirect(w, r, redirectURL, redirectCode)
			return
		}

		http.NotFound(w, r)
	})
}
