package model

import (
	"github.com/pkg/errors"
)

// NewModel creates a new instance
func NewModel(d DatabaseAdapter) *Model {
	return &Model{
		DB: d,
	}
}

// CreateRedirect creates the corresponding new redirect entry in the database
func (c *Model) CreateRedirect(redirect Redirect) error {
	err := redirect.Validate()
	if err != nil {
		return err
	}
	dbRedirect, err := compress(redirect)
	if err != nil {
		return err
	}
	return c.DB.CreateRedirect(dbRedirect)
}

// UpdateRedirectByDomain updates the corresponding redirect entry in the database by its domain name
func (c *Model) UpdateRedirectByDomain(domain string, redirect Redirect) error {
	err := redirect.Validate()
	if err != nil {
		return err
	}
	dbRedirect, err := compress(redirect)
	if err != nil {
		return errors.WithMessage(err, ErrRedirectCompress)
	}
	return c.DB.UpdateRedirectByDomain(domain, dbRedirect)
}

// GetRedirectByDomain returns a redirect entry by its domain name
func (c *Model) GetRedirectByDomain(domain string) (Redirect, error) {
	redirect := Redirect{}
	dbRedirect, err := c.DB.GetRedirectByDomain(domain)
	if err != nil {
		return redirect, err
	}
	redirect, err = decompress(dbRedirect)
	if err != nil {
		return redirect, errors.WithMessage(err, ErrRedirectsDecompress)
	}
	return redirect, nil
}

// GetRedirectsPaginated - using the provided cursor it returns a limited set of redirect entries and a new cursor
func (c *Model) GetRedirectsPaginated(cursor *string) ([]Redirect, *string, error) {
	dbRedirects, newCursor, err := c.DB.GetRedirectsPaginated(cursor)
	if err != nil {
		return nil, nil, err
	}
	redirects, err := multiDecompress(dbRedirects)
	if err != nil {
		return nil, nil, err
	}
	return redirects, newCursor, nil
}

// ImportRedirects truncates the database and creates the corresponding new redirect entries in the database
func (c *Model) ImportRedirects(redirects []Redirect) error {
	dbRedirects, err := multiCompress(redirects)
	if err != nil {
		return err
	}
	return c.DB.ImportRedirects(dbRedirects)
}

// ExportRedirects updates the cache and returns all contents from it
func (c *Model) ExportRedirects() ([]Redirect, error) {
	dbRedirects, err := c.DB.ExportRedirects()
	if err != nil {
		return nil, err
	}
	redirects, err := multiDecompress(dbRedirects)
	if err != nil {
		return nil, err
	}
	return redirects, nil
}

// DeleteRedirectByDomain deletes a redirect entry in the database by its domain name
func (c *Model) DeleteRedirectByDomain(domain string) error {
	return c.DB.DeleteRedirectByDomain(domain)
}
