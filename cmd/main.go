package main

import (
	"flag"
	"log"

	"github.com/sirkaiserkai/lb"
	"github.com/sirkaiserkai/lb/server"
)

// "github.com/sirkaiserkai/lb"

type serviceArgs struct {
	serviceType string
	port        string
	lb          string
}

func flags() serviceArgs {
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

func main() {
	// a := server.RegexAlphabet
	// fmt.Println(a[len(a)/2-1 : len(a)/2])

	args := flags()
	switch args.serviceType {
	case "lb":
		lb := lb.NewLoadBalancer(args.port)
		lb.Run()
	case "dummy":
		dummy := server.NewDummy(args.port, args.lb)
		dummy.Run()
	default:
		log.Fatal("Undefined service type parameter.")
	}
}
