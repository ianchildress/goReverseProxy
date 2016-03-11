// goReverseProxy is a simple reverse proxy that will forward requests to the configured
// target server and return the unmodified response to the client.
package main

import (
	"flag"
	"net/http"
)

func main() {

	const (
		defaultPort   = "80"
		defaultTarget = "http://127.0.0.1:8080"
	)
	var (
		targetPort string
		targetURL  string
		debug      bool
	)

	flag.StringVar(&targetPort, "port", defaultPort, "port the proxy will listen on. Example: -port 80")
	flag.StringVar(&targetURL, "target", defaultTarget, "url to forward the request to. Example: -target http://127.0.0.1:8080")
	flag.BoolVar(&debug, "debug", false, "Turn debug information on/off. Example: -debug true")
	flag.Parse()

	// Validate port and target formats here

	// Create the reverse proxy
	proxy := new(SingleProxy)
	proxy.New(targetURL, debug)

	http.HandleFunc("/", proxy.handle)
	http.ListenAndServe(":"+targetPort, nil)
}
