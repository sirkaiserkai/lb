package main

import (
	"log"

	"github.com/sirkaiserkai/lb"
	"github.com/sirkaiserkai/lb/server"
)

func main() {
	args := parseFlags()
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
