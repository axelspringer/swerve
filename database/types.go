package database

// Database is the API to the database
type Database struct {
	Service DynamoDBAdapter
	Config  Config
}

// Config contains the databases config
type Config struct {
	TableNamePrefix string
	Region          string
	TableRedirects  string
	TableCertCache  string
	TableUsers      string
	Key             string
	Secret          string
	Endpoint        string
}

// Redirect is the redirect entry model
type Redirect struct {
	RedirectFrom string  `json:"redirect_from"`
	Description  string  `json:"description"`
	RedirectTo   string  `json:"redirect_to"`
	Domain       string  `json:"domain"`
	Promotable   bool    `json:"promotable"`
	Code         int     `json:"code"`
	Created      int     `json:"created"`
	Modified     int     `json:"modified"`
	CPathMaps    *[]byte `json:"cpath-map,omitempty"`
}

// CertCacheEntry contains a certificate for the domain
type CertCacheEntry struct {
	Key   string `json:"domain"`
	Value string `json:"cert"`
}

// User contains a users cerdentials
type User struct {
	Name string `json:"username"`
	Pwd  string `json:"pwd"`
}
