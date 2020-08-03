package config

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Validate validates the configuration
func (c *Configuration) Validate() error {
	errs := make(map[string]string)

	if c.API.Version == "" {
		errs[paramStrAPIVersion] = "No API Version specified"
	}

	if c.API.Listener > 65535 || c.API.Listener < 1 {
		errs[paramStrAPIListenerPort] = "Port invalid - Valid range: 1-65535"
	}

	if c.HTTPListenerPort > 65535 || c.HTTPListenerPort < 1 {
		errs[paramStrHTTPListenerPort] = "Port invalid - Valid range: 1-65535"
	}

	if c.HTTPSListenerPort > 65535 || c.HTTPSListenerPort < 1 {
		errs[paramStrHTTPSListenerPort] = "Port invalid - valid range: 1-65535"
	}
	listener := []int{c.API.Listener, c.HTTPListenerPort, c.HTTPSListenerPort}
	if hasDuplicates(listener) {
		errs["listener"] = "Same port used multiple times"
	}

	_, err := url.ParseRequestURI(c.API.COR)
	if c.API.COR == "" || err != nil {
		errs[paramStrAPIUIURL] = "Invalid URL"
	}

	if _, err := logrus.ParseLevel(c.LogLevel); err != nil {
		errs[paramStrLogLevel] = "Invalid logrus log level - valid values: info, debug, warning, error, fatal, panic"
	}

	formatter := strings.ToLower(c.LogFormatter)
	if formatter != "text" && formatter != "json" {
		errs[paramStrLogFormatter] = "Invalid logrus log formatter - valid values: text, json"
	}

	if c.CacheInterval == 0 {
		errs[paramStrCacheInterval] = "Can not be zero"
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9_.-]*$`)
	match := re.Match([]byte(c.Database.TableNamePrefix))
	if !match {
		errs[paramStrTableNamePrefix] = "Contains invalid characters"
	}

	if c.Database.TableRedirects == "" {
		errs[paramStrTableRedirects] = "Can not be empty"
	}
	match = re.Match([]byte(c.Database.TableRedirects))
	if !match {
		errs[paramStrTableRedirects] = "Contains invalid characters"
	}

	if c.Database.TableCertCache == "" {
		errs[paramStrTableCertCache] = "Can not be empty"
	}
	match = re.Match([]byte(c.Database.TableCertCache))
	if !match {
		errs[paramStrTableCertCache] = "Contains invalid characters"
	}

	_, err = url.ParseRequestURI(c.Database.Endpoint)
	if c.API.COR == "" || err != nil {
		errs[paramStrAPIUIURL] = "Invalid URL"
	}

	if c.ACM.UsePebble && c.ACM.LetsEncryptURL == "" {
		errs[envStrUsePebble] = "When using pebble, you must provide a custom LetsEncrypt URL"
	}

	_, err = url.ParseRequestURI(c.ACM.PebbleCAURL)
	if c.ACM.UsePebble && c.ACM.PebbleCAURL != "" && err != nil {
		errs[paramStrPebbleCAURL] = "Invalid pebble CA URL"
	}

	_, err = url.ParseRequestURI(c.ACM.LetsEncryptURL)
	if !c.ACM.UsePebble &&
		!c.ACM.UseStage &&
		c.ACM.LetsEncryptURL != "" &&
		err != nil {
		errs[paramStrLetsEncryptURL] = "Invalid LetsEncrypt URL"
	}

	if len(errs) > 0 {
		return errors.New(fmt.Sprintf("%+v", errs))
	}
	return nil
}
