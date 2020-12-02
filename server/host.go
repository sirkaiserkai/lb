package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Host is an interface implemented by endpoint using the loadbalancer as a proxy.
type Host interface {
	Endpoint() string
	Health() (*HealthStatusResponse, error)
}

// GenericHost is the default struct that encapsulates a host.
type GenericHost struct {
	endpoint string
}

// Health returns the health of the service.
func (h GenericHost) Health() (*HealthStatusResponse, error) {
	resp, err := http.Get(h.Endpoint() + "/health")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	hs := HealthStatusResponse{}
	if err := json.Unmarshal(body, &hs); err != nil {
		return nil, err
	}
	return &hs, nil
}

func (h GenericHost) isHealthy() bool {
	return false
}

// Endpoint returns the endpoint string.
func (h GenericHost) Endpoint() string {
	return h.endpoint
}

// NewHostForAddHostRequest creates a host for a AddHostRequest
func NewHostForAddHostRequest(request AddHostRequest) GenericHost {
	return GenericHost{
		endpoint: request.Endpoint,
	}
}
