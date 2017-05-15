package registry

import "net/http"

// Controller as a Handler to handle the requests for register/unregister
// implaemets http.Handler
type Controller struct {
	*Registry
	http.Handler
}

func (rc Controller) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	switch path {
	case "/register":
		registerHost(rc.Registry, w, req)
	case "/unregister":
		unregisterHost(rc.Registry, w, req)
	default:
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
	}
}

/* -------------------------------------------------------------------------- */
// callbacks

func registerHost(r *Registry, w http.ResponseWriter, req *http.Request) {
	if err := r.Register(req.Host); err == nil {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
		w.Write([]byte(http.StatusText(500)))
	}
}

func unregisterHost(r *Registry, w http.ResponseWriter, req *http.Request) {
	if err := r.Unregister(req.Host); err == nil {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
		w.Write([]byte(http.StatusText(500)))
	}
}
