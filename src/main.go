package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: "127.0.0.1:9001", Path: "/"})
	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", proxy)
	if err != nil {
		log.Fatalln(err)
	}
}
