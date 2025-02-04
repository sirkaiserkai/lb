package server

import (
	"math/rand"
	"testing"
)

func TestAddHost(t *testing.T) {
	hostManager := NewHostManager()
	h := GenericHost{
		endpoint: "example.com",
	}
	// TODO: Test adding an existing host which is unhealthy.
	if err := hostManager.AddHost(h); err != nil {
		t.Error("error adding host: ", err.Error())
	}
}

// RemoveHostTest is used to set up
type HostTest struct {
	host Host
	err  error
}

func TestRemoveHost(t *testing.T) {
	// Set up test using sample endpoints.
	endpoints := []string{"example.com", "bomb.com", "xyz.com"}
	randHosts := make([]Host, len(endpoints))
	for i := 0; i < len(endpoints); i++ {
		randHosts[i] = NewHost(endpoints[i])
	}
	hostManager := NewHostManager()
	hostManager.SetHosts(randHosts)
	randomSample := randHosts[rand.Intn(len(randHosts))]
	tests := []HostTest{
		{
			host: randomSample,
			err:  nil,
		},
		{
			host: NewHost("notfound.com"),
			err:  ErrHostNotFound,
		},
	}

	for _, test := range tests {
		if err := hostManager.RemoveHost(test.host); err != nil {
			if test.err == nil {
				t.Errorf("Unexpected error. Received error: '%s'", err)
			} else if err != test.err {
				t.Errorf("Unexpected error. Expected error: '%s', received error: '%s'", test.err, err)
			}
		} else {
			// Verify host is gone
		}
	}
}
