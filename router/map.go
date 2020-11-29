package router

import "regexp"

// NewMap returns a new RouteMap instance for the given pattern.
func NewMap(rp RoutePattern) (*RouteMap, error) {
	reg, err := regexp.Compile(rp.p.RegexString)
	if err != nil {
		return nil, err
	}
	return &RouteMap{
		regex: reg,
		r:     rp,
	}, nil
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
