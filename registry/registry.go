package registry

import "fmt"

// Register a given host to the list of hosts to balance the load between them.
func (r *Registry) Register(host string) error {
	for _, h := range r.Hosts {
		if h == host {
			fmt.Println("[INFO:Register] - host already registered - ", host)
			return fmt.Errorf("Host already registered %s", host)
		}
	}

	r.Hosts = append(r.Hosts, host)
	fmt.Println("[INFO:Reigtser] - registered host - ", host)
	return nil
}

// Unregister a given host from the list of hosts that used tio balance the load.
func (r *Registry) Unregister(host string) error {
	for i, h := range r.Hosts {
		if h == host {
			r.Hosts = append(r.Hosts[:i], r.Hosts[i+1:]...)
			return nil
		}
	}

	fmt.Println("[INFO:Unregister] - host does not exist - ", host)
	return fmt.Errorf("Host does not exist - %s", host)
}

// IsEmpty checks if the registry is empty, hence can't handle the request
func (r *Registry) IsEmpty() bool {
	return len(r.Hosts) == 0
}

// Host returns the host at the specified index
func (r *Registry) Host(idx int) string {
	return r.Hosts[idx]
}

// Registry maintains the hosts list
type Registry struct {
	Hosts []string
}

// IRegistry interface
type IRegistry interface {
	Register(host string) error
	Unregister(host string) error
	IsEmpty() bool
	Host(idx int) string
}
