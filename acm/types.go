package acm

// Config contains the ACM config
type Config struct {
	UsePebble      bool
	UseStage       bool
	PebbleCA       string
	PebbleCAURL    string
	LetsEncryptURL string
}
