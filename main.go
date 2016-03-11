// goReverseProxy is a simple reverse proxy that will forward requests to the configured
// target server and return the unmodified response to the client.
package main

/*
import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	defaultPort   = "80"
	defaultTarget = "127.0.0.1:8080"
)

var (
	hostPort  string
	targetURL string
	debug     bool
)

func main() {

	flag.StringVar(&hostPort, "port", defaultPort, "port the proxy will listen on. Example: -port 80")
	flag.StringVar(&targetURL, "target", defaultTarget, "url to forward the request to. Example: -target 127.0.0.1:8080")
	flag.BoolVar(&debug, "debug", false, "Turn debug information on/off. Example: -debug true")
	flag.Parse()

	// Validate port and target formats here
	fmt.Println("Debug flag set to", debug)

	newURL := new(url.URL)
	newURL.Scheme = "http"
	newURL.Host = targetURL

	proxy := NewReverseProxy(newURL)
	log.Fatal(http.ListenAndServe(":"+hostPort, proxy))

}

func NewReverseProxy(target *url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}
	return &httputil.ReverseProxy{Director: director}
}
*/

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func report(r *http.Request) {
	r.Host = "stackoverflow.com"
	r.URL.Host = r.Host
	r.URL.Scheme = "http"
}

func main() {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:4567",
	})
	proxy.Director = report
	http.Handle("/", proxy)
	http.ListenAndServe(":80", nil)
}
