package router

import (
	"errors"
	"regexp"
)

// Pattern is a Regex pattern.
type Pattern struct {
	RegexString string
}

// RoutePattern maps a pattern to a Host endpoint. Not, the patterns are concatenated onto one another.
type RoutePattern struct {
	p    Pattern
	Host string
}

// New returns an ew Router instance for the given regex.
func New(routes []RoutePattern) (*Router, error) {
	maps := make([]RouteMap, len(routes))
	for i, r := range routes {
		reg, err := regexp.Compile(r.p.RegexString)
		if err != nil {
			return nil, err
		}
		maps[i] = RouteMap{
			regex: reg,
			r:     r,
		}
	}
	return &Router{routeMaps: maps}, nil
}

// Router maps a request to a given host, if the pattern exists.
type Router struct {
	routeMaps []RouteMap
}

// GetHost returns the host URL for the given request string.
func (r Router) GetHost(request string) (string, error) {
	for _, m := range r.routeMaps {
		if m.Match(request) {
			return m.r.Host, nil
		}
	}
	return "", errors.New("No match found.")
}

// RouteMap contains the regex map
type RouteMap struct {
	regex *regexp.Regexp
	r     RoutePattern
}

// Match validates whether the given request matches the router's regex.
func (r RouteMap) Match(request string) bool {
	return r.regex.MatchString(request)
}
