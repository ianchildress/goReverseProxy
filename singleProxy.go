package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// SingleProxy contains our *httputil.ReverseProxy settings.
type SingleProxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
	debug  bool
}

// New returns a simpleProxy
func (p *SingleProxy) New(target string, debug bool) *SingleProxy {
	url, _ := url.Parse(target)
	proxy := &SingleProxy{target: url, proxy: httputil.NewSingleHostReverseProxy(url), debug: debug}
	if proxy.proxy == nil {
		log.Panic("Returned nil proxy")
	}
	return proxy
}

func (p *SingleProxy) handle(w http.ResponseWriter, r *http.Request) {
	if p.debug == true {
		log.Println("Incoming request from ", r.RemoteAddr, "for uri", r.RequestURI)
	}
	w.Header().Set("X-goReverseProxy", "goReverseProxy")
	p.proxy.ServeHTTP(w, r)
}
