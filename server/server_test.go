package server

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestAddHost(t *testing.T) {
	lb := NewLoadBalancer(LoadBalancerConfig{})
	h := GenericHost{
		endpoint: "example.com",
	}
	if err := lb.AddHost(h); err != nil {
		t.Error("error adding host: ", err.Error())
	}
}

func TestRemoveHost(t *testing.T) {
	// Set up test using sample endpoints.
	r := rand.New(rand.NewSource(99))
	size := 100
	randHosts := make([]Host, size)
	for i := 0; i < size; i++ {
		endpoint := fmt.Sprintf("%d", r.Intn(size))
		randHosts[i] = NewHost(endpoint)
	}
	lb := NewLoadBalancer(LoadBalancerConfig{})
	lb.hosts = randHosts

	// Perform test.
	randomSample := randHosts[rand.Intn(len(randHosts))]
	if err := lb.RemoveHost(randomSample.Endpoint()); err != nil {
		t.Error("Unexpected error: ", err.Error())
	}
}
