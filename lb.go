package lb

import "github.com/sirkaiserkai/lb/server"

func NewLoadBalancer(port string) server.LoadBalancer {
	return server.NewLoadBalancer(server.LoadBalancerConfig{Port: port})
}
