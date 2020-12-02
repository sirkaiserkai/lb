package main

import (
	"fmt"

	"github.com/sirkaiserkai/lb/server"
)

func main() {
	a := server.RegexAlphabet
	fmt.Println(a[len(a)/2-1 : len(a)/2])
}
