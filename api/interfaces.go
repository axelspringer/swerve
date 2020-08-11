package api

// ModelAdapter is the interface for the business logic
type ModelAdapter interface {
	CreateRedirectFromJSON(jsonStr []byte) error
	UpdateRedirectByDomainWithJSON(domain string, jsonStr []byte) error
	DeleteRedirectByDomain(domain string) error
	GetRedirectByDomainAsJSON(domain string) ([]byte, error)
	ImportRedirectsFromJSON(jsonStr []byte) error
	ExportRedirectsAsJSON() ([]byte, error)
	GetRedirectsPaginatedAsJSON(cursor string) ([]byte, string, error)
	CheckPasswordFromJSON(jsonStr []byte, secret string) (string, int64, error)
	CheckToken(token string, secret string) bool
}
