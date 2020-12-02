package lb

import "github.com/sirkaiserkai/lb/server"

func NewLoadBalancer() server.LoadBalancer {
	return server.NewLoadBalancer(server.LoadBalancerConfig{})
}
