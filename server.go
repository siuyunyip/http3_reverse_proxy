package main

import (
	"fmt"
	"github.com/quic-go/quic-go/http3"
	"github.com/txthinking/runnergroup"
	"net/http"
)

// func NewServer(domain string, port string, root string) (*http.Server, *http.Server, *http3.Server, runnergroup.RunnerGroup) {
func NewServer(domain string, port string, root string) (*http.Server, runnergroup.RunnerGroup) {
	g, pool := InitWorker()
	handler := SetupHandler(root, pool.Workers)
	//tc := GetTlsConfig(domain)

	h := &http.Server{
		Addr:    ":80",
		Handler: handler,
	}

	go func() {
		//certFile, keyFile := testdata.GetCertificatePaths()
		certFile, keyFile := getCertificateKeyPaths()
		err := http3.ListenAndServe(":6101", certFile, keyFile, handler)
		if err != nil {
			fmt.Println(err)
		}
	}()

	//h2 := &http.Server{
	//	Addr:           ":" + port,
	//	ReadTimeout:    time.Duration(0) * time.Second,
	//	WriteTimeout:   time.Duration(0) * time.Second,
	//	IdleTimeout:    time.Duration(0) * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//	Handler:        handler,
	//	ErrorLog:       log.New(&tlserr{}, "", log.LstdFlags),
	//	TLSConfig:      tc,
	//}
	//
	//h3 := &http3.Server{
	//	Addr:      ":" + port,
	//	TLSConfig: tc,
	//	Handler:   handler,
	//}

	//return h, h2, h3, g
	return h, g
}
