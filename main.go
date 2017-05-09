package main

import (
	"net/http"

	"github.com/mhmoudgmal/slb/registery"
)

var (
	r = registery.Registery{
		Hosts: []string{
			"http://localhost:5000",
			"http://localhost:7000",
		},
	}

	currentHostIndex = 0
)

func main() {
	http.HandleFunc("/register", registerHost)
	http.HandleFunc("/unregister", unregisterHost)

	http.HandleFunc("/favicon.ico", nil)

	// Delegate all incoming requests to one of the registered servers round-robin
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		client := http.DefaultClient
		rw, rErr := client.Get(r.Hosts[currentHostIndex] + "/" + req.URL.Path)

		currentHostIndex++
		if currentHostIndex == len(r.Hosts) {
			currentHostIndex = 0
		}

		if rErr != nil {
			w.WriteHeader(400)
			w.Write([]byte(rErr.Error()))
		} else {
			w.WriteHeader(rw.StatusCode)
		}
	})

	http.ListenAndServe(":3000", nil)
}

/* -------------------------------------------------------------------------- */
// callbacks

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
}
