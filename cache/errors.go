package cache

// Error string constants - acme/autocert interface
const (
	ErrCacheManagerTimeout = "acme/autocert: cache manager timed out"
	ErrHostNotConfigured   = "acme/autocert: host '%s' not configured"
)

// Error string constants - Cache
const (
	ErrRedirectNotFound = "Redirect entry could not be found"
	ErrObserverRunning  = "Observer is already running"
	ErrCacheInconsitent = "Cache is inconsistent"
)
