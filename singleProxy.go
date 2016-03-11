package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type SingleProxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

// New returns a simpleProxy
func (p *SingleProxy) New(target string) *SingleProxy {
	url, _ := url.Parse(target)
	return &SingleProxy{target: url, proxy: httputil.NewSingleHostReverseProxy(url)}
}

func (p *SingleProxy) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-goReverseProxy", "goReverseProxy")
	p.proxy.ServeHTTP(w, r)
}
