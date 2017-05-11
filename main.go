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

func registerHost(w http.ResponseWriter, req *http.Request) {
	if err := r.Register(req.Host); err == nil {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
		w.Write([]byte(http.StatusText(500)))
	}
}

func unregisterHost(w http.ResponseWriter, req *http.Request) {
	if err := r.Unregister(req.Host); err == nil {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
		w.Write([]byte(http.StatusText(500)))
	}
	// to exit...
	var exit string
	fmt.Scanln(&exit)
}
