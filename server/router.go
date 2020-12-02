package server

import (
	"errors"
)

// Router maps a request to a given host, if the pattern exists.
type Router struct {
	routeMaps []*RouteMap
}

// RoutePattern maps a pattern to a Host endpoint. Not, the patterns are concatenated onto one another.
type RoutePattern struct {
	p    Pattern
	Host string
}

// Pattern is a Regex pattern.
type Pattern struct {
	RegexString string
}

// New returns an ew Router instance for the given regex.
func New(routes []RoutePattern) (*Router, error) {
	maps := make([]*RouteMap, len(routes))
	for i, rp := range routes {
		routeMap, err := NewMap(rp)
		if err != nil {
			return nil, err
		}
		maps[i] = routeMap
	}
	return &Router{routeMaps: maps}, nil
}

// GetHost returns the host URL for the given request string.
func (r Router) GetHost(request string) (string, error) {
	for _, m := range r.routeMaps {
		if m.Match(request) {
			return m.r.Host, nil
		}
	}
	return "", errors.New("no match found")
}
