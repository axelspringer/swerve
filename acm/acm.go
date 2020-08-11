package acm

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"

	"github.com/axelspringer/swerve/log"

	"github.com/axelspringer/swerve/config"

	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

// NewACM creates a new instance
func NewACM(hostPolicy autocert.HostPolicy, cache autocert.Cache, cfg *config.Configuration) *autocert.Manager {
	client := &acme.Client{
		HTTPClient: createHttpClient(cfg),
	}
	if cfg.UsePebble {
		client.DirectoryURL = cfg.LetsencryptUrl
	} else if cfg.UseStage {
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

func createHttpClient(cfg *config.Configuration) *http.Client {
	if cfg.UsePebble {
		cpool := x509.NewCertPool()
		cpool.AppendCertsFromPEM([]byte(cfg.PebbleCA))

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		resp, err := client.Get(cfg.PebbleCAUrl)
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
