package acm

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"

	"github.com/axelspringer/swerve/log"

	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

// NewACM creates a new instance
func NewACM(hostPolicy autocert.HostPolicy, cache autocert.Cache, c Config) *autocert.Manager {
	client := &acme.Client{
		HTTPClient: createHTTPClient(c),
	}
	if c.UsePebble {
		client.DirectoryURL = c.LetsEncryptURL
	} else if c.UseStage {
		client.DirectoryURL = LetsEncryptStagingURL
	} else {
		client.DirectoryURL = acme.LetsEncryptURL
	}
	return &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: hostPolicy,
		Cache:      cache,
		Client:     client,
	}
}

func createHTTPClient(c Config) *http.Client {
	if c.UsePebble {
		cpool := x509.NewCertPool()
		cpool.AppendCertsFromPEM([]byte(c.PebbleCA))

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		resp, err := client.Get(c.PebbleCAURL)
		if err != nil {
			log.Fatal(err)
		}
		cert, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		cpool.AppendCertsFromPEM(cert)

		trWithValidation := &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: cpool,
			},
		}

		httpclient := &http.Client{
			Transport: trWithValidation,
		}
		return httpclient
	}
	return nil
}
