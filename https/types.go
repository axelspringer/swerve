package https

import (
	"crypto/tls"
	"net/http"

	"github.com/axelspringer/swerve/model"
)

// HTTPS is an HTTPS server
type HTTPS struct {
	getRedirect func(name string) (model.Redirect, error)
	server      *http.Server
	listener    int
}

// GetRedirect function to retrieve redirect entry
type GetRedirect func(name string) (model.Redirect, error)

// GetCertificate tls client hello function to retrieving certificate
type GetCertificate func(hello *tls.ClientHelloInfo) (*tls.Certificate, error)
