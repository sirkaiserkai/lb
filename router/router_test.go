package router

import "testing"

// Should probably split this out into two structs. One for configuration another for the tests.RouterTest
// At this point, it's misleading because a 'request' may match to a previous RoutePattern.
type RouterTest struct {
	pattern RoutePattern
	request string
	isMatch bool
}

var tests []RouterTest = []RouterTest{
	{
		pattern: RoutePattern{
			p: Pattern{
				RegexString: "^00[0-z].*",
			},
			Host: "xyz.com",
		},
		request: "00U123123ABCaa",
		isMatch: true,
	},
	{
		pattern: RoutePattern{
			p: Pattern{
				RegexString: "^00[0-z].*",
			},
			Host: "xyz.com",
		},
		request: "12345678",
		isMatch: false,
	},
}

func TestMatch(t *testing.T) {

	// Set up tests
	routes := make([]RoutePattern, len(tests))
	for i, test := range tests {
		routes[i] = test.pattern
	}

	r, err := New(routes)
	if err != nil {
		t.Errorf("Unexpected error: '%s'", err)
		return
	}

	for i, test := range tests {
		if r.routeMaps[i].Match(test.request) != test.isMatch {
			t.Errorf("Match was not equal to expected output.")
		}
	}
}

func TestGetHost(t *testing.T) {
	// Set up tests
	routes := make([]RoutePattern, len(tests))
	for i, test := range tests {
		routes[i] = test.pattern
	}

	r, err := New(routes)
	if err != nil {
		t.Errorf("Unexpected error: '%s'", err)
		return
	}
	for _, test := range tests {
		host, err := r.GetHost(test.request)
		if err != nil {
			if test.isMatch {
				t.Errorf("Failed to map to host: '%s'", err)
			}
		}
		if test.isMatch && host != test.pattern.Host {
			t.Errorf("Failed to map to correct host. Expected: '%s'. Received: '%s'", test.pattern.Host, host)
		}
	}
}
