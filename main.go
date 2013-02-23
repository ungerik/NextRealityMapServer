package main

import (
	_ "github.com/ungerik/NextRealityMapServer/mapserver"
	"github.com/ungerik/go-rest"
)

func main() {

	rest.RunServer("0.0.0.0:8080", nil)
}
