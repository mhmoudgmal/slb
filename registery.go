package slb

import (
	"errors"
	"fmt"
)

// Register a given host to the list of hosts to balance the load between them.
func (r *Registery) Register(host string) error {
	for _, h := range r.Hosts {
		if h == host {
			fmt.Println("[INFO:Register] - host already registered")
			return errors.New("Host already registered")
		}
	}

	r.Hosts = append(r.Hosts, host)
	fmt.Println("[INFO:Reigtser] - registered host - ", host)
	return nil
}

// Unregister a given host from the list of hosts that used tio balance the load.
func (r *Registery) Unregister(host string) error {
	for i, h := range r.Hosts {
		if h == host {
			r.Hosts = append(r.Hosts[:i], r.Hosts[i+1:]...)
			return nil
		}
	}

	fmt.Println("[INFO::Unregister] - host does not exist")
	return errors.New("Host dies not exist")
}

// Registery maintains the hosts list
type Registery struct {
	Hosts []string
}
