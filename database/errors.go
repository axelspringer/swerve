package database

// Error string constants - Session
const (
	ErrSessionCreate = "Session could not be created"
)

// Error string constants - CertCache
const (
	ErrCertCacheFetch   = "Cert cache entry could not be fetched"
	ErrCertCacheUpdate  = "Cert cache entry could not be updated"
	ErrCertCacheDelete  = "Cert cache entry could not be deleted"
	ErrCertCacheMarshal = "Cert cache entry could not be marshaled"
)

// Error string constants - Redirects
const (
	ErrRedirectMarshal      = "Redirect entry could not be marshaled"
	ErrRedirectCreate       = "Redirect entry could not be created"
	ErrRedirectDelete       = "Redirect entry could not be deleted"
	ErrRedirectUpdate       = "Redirect entry could not be updated"
	ErrRedirectUpdatePartly = "Redirect entry could only be updated partly"
	ErrRedirectNotExist     = "Redirect entry using this domain does not exist"
	ErrRedirectsFetch       = "Redirect entries could not be fetched"
	ErrRedirectsUnmarshal   = "Redirect entries could not be unmarshaled"
	ErrRedirectsScan        = "Table could not be scanned"
	ErrCursorDecode         = "Cursor could not be decoded"
	ErrRedirectListEmpty    = "Paginated redirect entry list ist empty"
	ErrRedirectExists       = "Redirect entry alerady exists"
	ErrRedirectNotFound     = "Redirect entry could not be found"
)

// Error string constants - Auth
const (
	ErrUserNotFound = "User could not be found"
	ErrUserFetch    = "User could not be fetched"
)

// Error string contants - Tables
const (
	ErrfTableCreate = "Table '%s' could not be created"
)
