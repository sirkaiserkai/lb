package main

import "flag"

type serviceArgs struct {
	serviceType string
	port        string
	lb          string
}

func parseFlags() serviceArgs {
	serviceType := flag.String("type", "lb", "The service type to launch.")
	port := flag.String("port", "8080", "The port to launch the service on.")
	lb := flag.String("lb", "http://localhost:8080", "The endpoint for the load balancer (only used by storage nodes).")

	flag.Parse()
	return serviceArgs{
		serviceType: *serviceType,
		port:        *port,
		lb:          *lb,
	}
}
