package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

func setupHandler(root string) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(root)))

	mux.HandleFunc("/1/", func(w http.ResponseWriter, r *http.Request) {

	})

	mux.HandleFunc("/2/", func(w http.ResponseWriter, r *http.Request) {

	})

	mux.HandleFunc("/3/", func(w http.ResponseWriter, r *http.Request) {

	})

	mux.HandleFunc("/4/", func(w http.ResponseWriter, r *http.Request) {

	})
}

func main() {
	port := flag.String("port", "443", "bind to port")
	domain := flag.String("domain", "dev.cafewithbook.org", "domain name")
	// default: www == ./html
	www := flag.String("www", "./html", "web root")
	flag.Parse()

	handler := setupHandler(*www)
	tc := GetTlsConfig(*domain)

	h2 := &http.Server{
		Addr:           ":" + *port,
		ReadTimeout:    time.Duration(0) * time.Second,
		WriteTimeout:   time.Duration(0) * time.Second,
		IdleTimeout:    time.Duration(0) * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        handler,
		ErrorLog:       log.New(&tlserr{}, "", log.LstdFlags),
		TLSConfig:      tc,
	}

	h3 := &http.Server{
		Addr:      ":" + *port,
		TLSConfig: tc,
		Handler:   handler,
	}
}
