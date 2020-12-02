package server

import "testing"

func TestAddHost(t *testing.T) {
	lb := NewLoadBalancer(LoadBalancerConfig{})
	h := GenericHost{
		endpoint: "example.com",
	}
	if err := lb.AddHost(h); err != nil {
		t.Error("error adding host: ", err.Error())
	}
}
