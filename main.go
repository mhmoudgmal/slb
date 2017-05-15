package main

import (
	"fmt"
	"net/http"

	"github.com/mhmoudgmal/slb/balancer"
	"github.com/mhmoudgmal/slb/registry"
)

var (
	reg = &registry.Registry{
		Hosts: []string{
			"localhost:5000",
			"localhost:7000",
		},
	}
)

func main() {
	// Handle all incoming requests
	http.HandleFunc("/", balancer.Controller{Registry: reg}.Handle)

	go balancer.HandleRequests()

	go http.ListenAndServe(":3000", nil)
	go http.ListenAndServe(":3001", registry.Controller{Registry: reg})

	// to exit...
	var exit string
	fmt.Scanln(&exit)
}
