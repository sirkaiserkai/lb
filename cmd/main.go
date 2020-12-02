package main

import (
	"fmt"
	"time"

	"github.com/sirkaiserkai/lb/server"
)

func main() {
	go doNothing()
	a := server.RegexAlphabet
	fmt.Println(a[len(a)/2-1 : len(a)/2])
	time.Sleep(5 * time.Second)
}

func doNothing() {
	time.Sleep(1 * time.Second)
	fmt.Println("Done!")
}
