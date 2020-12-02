package server

const (
	// RegexAlphabet is the available alphabet for regular expression patterns.
	RegexAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// RouterConfig is the configuration object for the server.
type RouterConfig struct {
	Patterns []Pattern
	Routes   []RoutePattern
}

// LoadBalancerConfig is the configuration object for a loadbalancer.
type LoadBalancerConfig struct {
}
