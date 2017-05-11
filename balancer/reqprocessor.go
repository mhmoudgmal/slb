package balancer

import (
	"fmt"
	"net/http"
)

var (
	requestChannel = make(chan SLBRequest)
)

// NewRequestProcessor ....
func NewRequestProcessor() struct{ Ch chan SLBRequest } {
	reqChannel := struct {
		Ch chan SLBRequest
	}{requestChannel}

	return reqChannel
}

// HandleRequests ...
func HandleRequests() {
	for {
		select {
		case reqCh := <-requestChannel:
			fmt.Println(reqCh.Host)
			reqCh.Done <- true
		}
	}
}

/* -------------------------------------------------------------------------- */
// Structs

// SLBRequest Encapsulates the request
type SLBRequest struct {
	Host string
	http.ResponseWriter
	*http.Request
	Done chan bool
}
