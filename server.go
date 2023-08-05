package main

import (
	"github.com/quic-go/quic-go/http3"
	"log"
	"net/http"
	"time"
)

func NewServer(domain string, port string, root string) (*http.Server, *http.Server, *http3.Server, error) {
	handler := SetupHandler(root)
	tc := GetTlsConfig(domain)

	h := &http.Server{
		Addr:    ":80",
		Handler: handler,
	}

	h2 := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    time.Duration(0) * time.Second,
		WriteTimeout:   time.Duration(0) * time.Second,
		IdleTimeout:    time.Duration(0) * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        handler,
		ErrorLog:       log.New(&tlserr{}, "", log.LstdFlags),
		TLSConfig:      tc,
	}

	h3 := &http3.Server{
		Addr:      ":" + port,
		TLSConfig: tc,
		Handler:   handler,
	}

	return h, h2, h3, nil

}