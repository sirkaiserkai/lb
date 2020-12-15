package server

import "regexp"

// NewMap returns a new RouteMap instance for the given pattern.
func NewMap(pattern string) (*RouteMap, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &RouteMap{
		regex:   reg,
		pattern: pattern,
	}, nil
}

// RouteMap contains the regex map
type RouteMap struct {
	regex   *regexp.Regexp
	pattern string
}

// Match validates whether the given request matches the router's regex.
func (r RouteMap) Match(request string) bool {
	return r.regex.MatchString(request)
}
