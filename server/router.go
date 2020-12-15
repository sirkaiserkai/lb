package server

import (
	"errors"
)

// Router maps a request to a given host, if the pattern exists.
type Router struct {
	hostManager *HostManager
}

// NewRouter returns a new Router instance for the given regex.
func NewRouter(hm *HostManager) (*Router, error) {
	return &Router{hostManager: hm}, nil
}

// GetHost returns the host URL for the given request string.
func (r Router) GetHost(request string) (Host, error) {
	for _, h := range r.hostManager.GetHosts() {
		if h.Route().Match(request) {
			return h, nil
		}
	}
	return nil, errors.New("no match found")
}
