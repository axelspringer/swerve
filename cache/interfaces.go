package cache

import "github.com/axelspringer/swerve/database"

// DatabaseAdapter is the interface required for the database interacion
type DatabaseAdapter interface {
	CreateRedirect(redirect database.Redirect) error
	GetRedirectByDomain(name string) (database.Redirect, error)
	DeleteRedirectByDomain(name string) error
	UpdateRedirectByDomain(name string, redirect database.Redirect) error
	GetRedirectsPaginated(cursor *string) ([]database.Redirect, *string, error)
	ImportRedirects(redirects []database.Redirect) error
	ExportRedirects() ([]database.Redirect, error)
	UpdateCacheEntry(key string, data []byte) error
	GetCacheEntry(key string) ([]byte, error)
	DeleteCacheEntry(key string) error
	GetPwdHash(username string) (string, error)
}
