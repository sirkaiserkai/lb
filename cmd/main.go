package main

import "github.com/sirkaiserkai/lb/server"

// "github.com/sirkaiserkai/lb"

func main() {
	// a := server.RegexAlphabet
	// fmt.Println(a[len(a)/2-1 : len(a)/2])
	// lb := lb.NewLoadBalancer()
	// lb.Run()

	dummy := server.NewDummy()
	dummy.Run()
}
