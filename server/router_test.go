package server

import "testing"

// Should probably split this out into two structs. One for configuration another for the tests.RouterTest
// At this point, the structure misleading because a 'request' may match to a previous RoutePattern.
type RouterTest struct {
	endpoint string
	pattern  string
	request  string
	isMatch  bool
}

var tests []RouterTest = []RouterTest{
	{
		endpoint: "00z.com",
		pattern:  "^00[0-z].*",
		request:  "00U123123ABCaa",
		isMatch:  true,
	},
	{
		endpoint: "abc.com",
		pattern:  "^abc[0-z].*",
		request:  "12345678",
		isMatch:  false,
	},
}

func TestMatch(t *testing.T) {

	// Set up tests
	routes := make([]RouteMap, len(tests))
	for i, test := range tests {
		m, err := NewMap(test.pattern)
		if err != nil {
			t.Errorf(err.Error())
			return
		}
		routes[i] = *m
	}

	for i, test := range tests {
		if routes[i].Match(test.request) != test.isMatch {
			t.Errorf("Match was not equal to expected output.")
		}
	}
}

func TestGetHost(t *testing.T) {
	// Set up tests
	hm := NewHostManager()
	r, err := NewRouter(&hm)
	if err != nil {
		t.Errorf("Unexpected error: '%s'", err)
		return
	}
	for _, test := range tests {
		m, err := NewMap(test.pattern)
		if err != nil {
			t.Errorf("Unexpected error: '%s'", err)
			return
		}
		h := GenericHost{
			endpoint: test.endpoint,
			route:    *m,
		}
		hm.AddHost(h)
	}

	for _, test := range tests {
		host, err := r.GetHost(test.request)
		if err != nil {
			if test.isMatch {
				t.Errorf("Failed to map to host: '%s'", err)
			}
		}
		if test.isMatch && host.Endpoint() != test.endpoint {
			t.Errorf("Failed to map to correct host. Expected: '%s'. Received: '%s'", test.endpoint, host.Endpoint())
		}
	}
}
