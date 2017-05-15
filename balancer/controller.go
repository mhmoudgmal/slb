package balancer

import (
	"net/http"
	"sync"

	"github.com/mhmoudgmal/slb/registry"
)

var (
	lock sync.RWMutex

	currentHostIndex = 0
	reqProcessor     = NewRequestProcessor()
)

// Controller is responsible for wrapping every request with more information -
// needed for the communication between the channels
// - [ done:            for informing back when the request has finished]
// - [ reqProcessor.Ch: for accepting requests messages to route it to the desired server]
type Controller struct {
	Registry *registry.Registry
}

// Handle takes over the job from from (/) handler.
func (c Controller) Handle(w http.ResponseWriter, req *http.Request) {
	if c.Registry.IsEmpty() {
		w.WriteHeader(503)
		w.Write([]byte(http.StatusText(503)))
		return
	}

	done := make(chan bool)
	slbRequest := &SLBRequest{
		ResponseWriter: w,
		Request:        req,
		Done:           done,
		Host:           c.Registry.Host(currentHostIndex),
	}

	c.nextHost()

	// send request message to the channel
	reqProcessor.Ch <- slbRequest

	// blocks until the request is done
	<-done
}

/* -------------------------------------------------------------------------- */
// Helper functions

func (c Controller) nextHost() {
	lock.RLock()
	currentHostIndex++
	if currentHostIndex == len(c.Registry.Hosts) {
		currentHostIndex = 0
	}
	lock.RUnlock()
}
