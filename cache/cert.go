package cache

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/acme/autocert"
)

// Get is used by the acm to retrieve a certificate
func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	var (
		data []byte
		err  error
	)
	done := make(chan struct{})

	go func() {
		defer close(done)
		data, err = c.DB.GetCacheEntry(key)
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-done:
	case <-time.After(time.Minute):
		return nil, errors.New(ErrCacheManagerTimeout)
	}

	if err == nil && data != nil && len(data) > 0 {
		return data, nil
	}

	if err != nil {
		return nil, errors.WithMessage(err, autocert.ErrCacheMiss.Error())
	}

	return nil, autocert.ErrCacheMiss
}

// Put is used by the acm to put a certificate
func (c *Cache) Put(ctx context.Context, key string, data []byte) error {
	var err error
	done := make(chan struct{})

	go func() {
		defer close(done)
		err = c.DB.UpdateCacheEntry(key, data)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
	case <-time.After(time.Minute):
		return errors.New(ErrCacheManagerTimeout)
	}

	return err
}

// Delete is used by the acm to delete a certificate
func (c *Cache) Delete(ctx context.Context, key string) error {
	var err error
	done := make(chan struct{})

	go func() {
		defer close(done)
		err = c.DB.DeleteCacheEntry(key)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
	case <-time.After(time.Minute):
		return errors.New(ErrCacheManagerTimeout)
	}

	return err
}
