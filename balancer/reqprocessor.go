package balancer

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	requestChannel = make(chan *SLBRequest)
	client         = http.Client{}
)

// NewRequestProcessor ....
func NewRequestProcessor() struct{ Ch chan *SLBRequest } {
	reqChannel := struct {
		Ch chan *SLBRequest
	}{requestChannel}

	return reqChannel
}

// HandleRequests ...
func HandleRequests() {
	for {
		select {
		case req := <-requestChannel:
			handle(req.Host, req)
		}
	}
}

// TODO: refactor and simplify for readability and maintainability
func handle(host string, req *SLBRequest) {
	hostURL, _ := url.Parse(req.Request.URL.String())
	hostURL.Scheme = "http"
	hostURL.Host = host

	newRequest, newReqErr := http.NewRequest(
		req.Request.Method,
		hostURL.String(),
		req.Request.Body)

	if newReqErr != nil {
		fmt.Println("Failed creating new request")
		fmt.Println(newReqErr)
		req.ResponseWriter.WriteHeader(500)
		req.ResponseWriter.Write([]byte(http.StatusText(500)))
		req.Done <- true
		return
	}

	// Copy all headers for the new request from the origianl request
	for k, v := range req.Request.Header {

		headerVals := ""
		for _, headerVal := range v {
			headerVals += headerVal + " "
		}

		newRequest.Header.Add(k, headerVals)
	}

	res, err := client.Do(newRequest)
	fmt.Printf("Request will be done by %s\n", newRequest.Host)

	if err != nil {
		// TODO: handle response properly, not every error is 500!
		req.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		req.Done <- true
		return
	}

	// Copy all the headers from the response to the original response writer.
	for k, v := range res.Header {

		headerVals := ""
		for _, headerVal := range v {
			headerVals += headerVal + " "
		}

		req.ResponseWriter.Header().Add(k, headerVals)
	}

	io.Copy(req.ResponseWriter, res.Body)

	req.Done <- true
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
