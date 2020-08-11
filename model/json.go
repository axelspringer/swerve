package model

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// CreateRedirectFromJSON takes JSON and creates the corresponding new redirect entry in the database
func (c *Model) CreateRedirectFromJSON(jsonStr []byte) error {
	var redirect Redirect
	if err := json.Unmarshal(jsonStr, &redirect); err != nil {
		return errors.WithMessage(err, ErrBodyUnmarshal)
	}
	return c.CreateRedirect(redirect)
}

// UpdateRedirectByDomainWithJSON takes JSON and updates the corresponding redirect entry in the database by its domain name
func (c *Model) UpdateRedirectByDomainWithJSON(domain string, jsonStr []byte) error {
	var redirect Redirect
	if err := json.Unmarshal(jsonStr, &redirect); err != nil {
		return errors.WithMessage(err, ErrBodyUnmarshal)
	}
	return c.UpdateRedirectByDomain(domain, redirect)
}

// GetRedirectByDomainAsJSON returns a redirect entry in JSON format by its domain name
func (c *Model) GetRedirectByDomainAsJSON(domain string) ([]byte, error) {
	redirect, err := c.GetRedirectByDomain(domain)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(redirect)
	if err != nil {
		return nil, errors.WithMessage(err, ErrRedirectMarshal)
	}
	return data, nil
}

// GetRedirectsPaginatedAsJSON - using the provided cursor it returns a limited set of redirect entries in JSON format and a new cursor
func (c *Model) GetRedirectsPaginatedAsJSON(cursor string) ([]byte, string, error) {
	var usedCursor *string
	usedCursor = &cursor
	if *usedCursor == "" {
		usedCursor = nil
	}
	redirects, newCursor, err := c.GetRedirectsPaginated(usedCursor)
	if err != nil {
		return nil, "", err
	}
	data, err := json.Marshal(redirects)
	if err != nil {
		return nil, "", errors.WithMessage(err, ErrRedirectsMarshal)
	}
	if newCursor == nil {
		empty := ""
		newCursor = &empty
	}
	return data, *newCursor, nil
}

// ImportRedirectsFromJSON takes JSON, truncates the database and creates the corresponding new redirect entries in the database
func (c *Model) ImportRedirectsFromJSON(jsonStr []byte) error {
	var redirects []Redirect
	if err := json.Unmarshal(jsonStr, &redirects); err != nil {
		return errors.WithMessage(err, ErrBodyUnmarshal)
	}
	return c.ImportRedirects(redirects)
}

// ExportRedirectsAsJSON returns the database contents in JSON format
func (c *Model) ExportRedirectsAsJSON() ([]byte, error) {
	redirects, err := c.ExportRedirects()
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(redirects)
	if err != nil {
		return nil, errors.WithMessage(err, ErrRedirectsMarshal)
	}
	return data, nil
}

// CheckPasswordFromJSON takes JSON and a secret for the JWT and checks pwd hash on db against entered plain pwd, returns nil if it is correct
func (c *Model) CheckPasswordFromJSON(jsonStr []byte, secret string) (string, int64, error) {
	var user User
	if err := json.Unmarshal(jsonStr, &user); err != nil {
		return "", 0, errors.WithMessage(err, ErrBodyUnmarshal)
	}
	if err := c.CheckPassword(user.Name, user.Pwd); err != nil {
		return "", 0, err
	}
	tokenString, expirationTime, err := c.CreateJWT(user.Name, secret)
	if err != nil {
		return "", 0, err
	}
	return tokenString, expirationTime, err
}
