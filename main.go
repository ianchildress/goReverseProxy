// goReverseProxy is a simple reverse proxy that will forward requests to the configured
// target server and return the unmodified response to the client.
package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:4567",
	})

	http.Handle("/", proxy)
	http.ListenAndServe(":80", nil)
}
