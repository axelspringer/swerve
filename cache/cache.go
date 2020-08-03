package cache

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/axelspringer/swerve/log"

	"github.com/axelspringer/swerve/database"
	"github.com/pkg/errors"
)

// NewCache creates a new instance
func NewCache(d DatabaseAdapter) *Cache {
	return &Cache{
		DB:           d,
		Observing:    false,
		mapMutex:     &sync.RWMutex{},
		redirectsMap: map[string]*database.Redirect{},
	}
}

// Observe updates the certificate cache at the given interval (in minutes)
func (c *Cache) Observe(pollInterval int) error {
	if c.Observing {
		return errors.New(ErrObserverRunning)
	}
	c.mapMutex.Lock()
	c.closer = make(chan struct{})
	c.mapMutex.Unlock()
	ticker := time.NewTicker(time.Minute * time.Duration(pollInterval))
	go func() {
		for c.Observing {
			select {
			case <-ticker.C:
				c.Update()
			case <-c.closer:
				c.mapMutex.Lock()
				c.Observing = false
				c.mapMutex.Unlock()
				ticker.Stop()
				return
			case <-time.After(time.Minute):
				log.Warn("Observer timed out")
				ticker.Stop()
				return
			}
		}
	}()
	return nil
}

// ObserverCloser returns the observers closer channel to get a signal when the observer is closed
func (c *Cache) ObserverCloser() chan struct{} {
	return c.closer
}

// CloseObserver closes the observer closer channel, closing the observer
func (c *Cache) CloseObserver() {
	close(c.closer)
}

// Update updates the local redirect entry cache
func (c *Cache) Update() {
	redirects, err := c.DB.ExportRedirects()
	if err != nil {
		log.Warn(errors.WithMessage(err, "Redirect entries could not be fetched").Error())
		return
	}

	c.mapMutex.Lock()
	defer c.mapMutex.Unlock()

	c.redirectsMap = map[string]*database.Redirect{}

	for i, redirect := range redirects {
		c.redirectsMap[redirect.RedirectFrom] = &redirects[i]
	}
}

// AllowHostPolicy decides which host shall pass
func (c *Cache) AllowHostPolicy(_ context.Context, host string) error {
	if _, err := c.GetRedirectByDomain(host); err != nil {
		return errors.WithMessage(err, fmt.Sprintf(ErrHostNotConfigured, host))
	}
	return nil
}
