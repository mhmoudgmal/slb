package main

import (
	"fmt"
	"net/http"

	"github.com/mhmoudgmal/slb/balancer"
	"github.com/mhmoudgmal/slb/registery"
)

var (
	reg = &registery.Registery{
		Hosts: []string{
			"http://localhost:5000",
			"http://localhost:7000",
		},
	}
)

func main() {
	// Handle all incoming requests
	http.HandleFunc("/", balancer.Controller{Registery: reg}.Handle)

	go balancer.HandleRequests()

	go http.ListenAndServe(":3000", nil)
	go http.ListenAndServe(":3001", registery.Controller{Registery: reg})

	// to exit...
	var exit string
	fmt.Scanln(&exit)
}
