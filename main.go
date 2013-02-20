package main

import (
	"github.com/ungerik/go-rest"
)

func isPow2(x int) bool {
	return x&(x-1) == 0
}

func main() {

	rest.RunServer("0.0.0.0:8080", nil)
}
