package model

import jwt "github.com/dgrijalva/jwt-go"

// Model is the business logic API
type Model struct {
	DB DatabaseAdapter
}

// Redirect is the redirect entry model
type Redirect struct {
	RedirectFrom string    `json:"redirect_from"`
	Description  string    `json:"description"`
	RedirectTo   string    `json:"redirect_to"`
	Promotable   bool      `json:"promotable"`
	Code         int       `json:"code"`
	Created      int       `json:"created"`
	Modified     int       `json:"modified"`
	PathMaps     []PathMap `json:"path-map,omitempty"`
}

// PathMap contains a condition
type PathMap struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type paginatedRedirects struct {
	Data   []Redirect `json:"data"`
	Cursor string     `json:"next_cursor,omitempty"`
}

// User contains a users cerdentials
type User struct {
	Name string `json:"username"`
	Pwd  string `json:"pwd"`
}

type claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
