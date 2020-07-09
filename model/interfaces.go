package model

import "github.com/axelspringer/swerve/database"

// DatabaseAdapter is the interface required for the database interacion via cache
type DatabaseAdapter interface {
	CreateRedirect(redirect database.Redirect) error
	GetRedirectByDomain(name string) (database.Redirect, error)
	DeleteRedirectByDomain(name string) error
	UpdateRedirectByDomain(name string, redirect database.Redirect) error
	GetRedirectsPaginated(cursor *string) ([]database.Redirect, *string, error)
	ImportRedirects(redirects []database.Redirect) error
	ExportRedirects() ([]database.Redirect, error)
	GetPwdHash(username string) (string, error)
}
