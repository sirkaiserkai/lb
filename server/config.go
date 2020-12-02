package server

const (
	// RegexAlphabet is the available alphabet for regular expression patterns.
	RegexAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// Config is the configuration object for the server.
type Config struct {
	Patterns []Pattern
	Routes   []RoutePattern
}
