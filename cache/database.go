package cache

import (
	"github.com/axelspringer/swerve/database"
	"github.com/pkg/errors"
)

// GetRedirectByDomain returns a redirect entry from the local cache
func (c *Cache) GetRedirectByDomain(name string) (database.Redirect, error) {
	c.mapMutex.Lock()
	defer c.mapMutex.Unlock()
	if redirect, ok := c.redirectsMap[name]; ok {
		if redirect == nil {
			return database.Redirect{}, errors.New(ErrCacheInconsitent)
		}
		return *redirect, nil
	}
	redirect, err := c.DB.GetRedirectByDomain(name)
	if err == nil {
		c.redirectsMap[name] = &redirect
		return redirect, nil
	}
	return database.Redirect{}, errors.New(ErrRedirectNotFound)
}

// ExportRedirects updates the cache and returns all redirect entries from it
func (c *Cache) ExportRedirects() ([]database.Redirect, error) {
	c.Update()
	return c.GetAllRedirects()
}

// GetAllRedirects returns all redirect entries from the local cache
func (c *Cache) GetAllRedirects() ([]database.Redirect, error) {
	redirects := []database.Redirect{}
	c.mapMutex.RLock()
	defer c.mapMutex.RUnlock()
	if c.redirectsMap == nil {
		return redirects, nil
	}

	for k, v := range c.redirectsMap {
		if v == nil {
			return nil, errors.New(ErrCacheInconsitent)
		}
		redirects = append(redirects, *c.redirectsMap[k])
	}
	return redirects, nil
}

// GetRedirectsPaginated database wrapper
func (c *Cache) GetRedirectsPaginated(cursor *string) ([]database.Redirect, *string, error) {
	return c.DB.GetRedirectsPaginated(cursor)
}

// CreateRedirect database wrapper
func (c *Cache) CreateRedirect(redirect database.Redirect) error {
	err := c.DB.CreateRedirect(redirect)
	if err == nil {
		c.Update()
	}
	return err
}

// DeleteRedirectByDomain database wrapper
func (c *Cache) DeleteRedirectByDomain(name string) error {
	err := c.DB.DeleteRedirectByDomain(name)
	if err == nil {
		c.Update()
	}
	return err
}

// UpdateRedirectByDomain database wrapper
func (c *Cache) UpdateRedirectByDomain(name string, redirect database.Redirect) error {
	err := c.DB.UpdateRedirectByDomain(name, redirect)
	if err == nil {
		c.Update()
	}
	return err
}

// ImportRedirects database wrapper
func (c *Cache) ImportRedirects(redirects []database.Redirect) error {
	err := c.DB.ImportRedirects(redirects)
	if err == nil {
		c.Update()
	}
	return err
}

// UpdateCacheEntry database wrapper
func (c *Cache) UpdateCacheEntry(key string, data []byte) error {
	return c.DB.UpdateCacheEntry(key, data)
}

// GetCacheEntry database wrapper
func (c *Cache) GetCacheEntry(key string) ([]byte, error) {
	return c.DB.GetCacheEntry(key)
}

// DeleteCacheEntry database wrapper
func (c *Cache) DeleteCacheEntry(key string) error {
	return c.DB.DeleteCacheEntry(key)
}

// GetPwdHash database wrapper
func (c *Cache) GetPwdHash(username string) (string, error) {
	return c.DB.GetPwdHash(username)
}
