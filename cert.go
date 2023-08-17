package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path"
)

var certPath = "/home/ubuntu/quic/certs/live/dev.cafewithbook.org"

func getCertificateKeyPaths() (string, string) {
	return path.Join(certPath, "cert.pem"), path.Join(certPath, "privkey.pem")
}

func GetTlsConfig(domain string) *tls.Config {
	c, k := getCertificateKeyPaths()
	cert, err := ioutil.ReadFile(c)
	if err != nil && !os.IsNotExist(err) {
		fmt.Printf("error: %s", "no cert file")
	}

	key, err := ioutil.ReadFile(k)
	if err != nil && !os.IsNotExist(err) {
		fmt.Printf("error: %s", "no key file")
	}

	l := make([]tls.Certificate, 0)
	certs := make(map[string]*tls.Certificate)
	if cert != nil && key != nil {
		certificate, err := tls.X509KeyPair(cert, key)
		if err != nil {
			fmt.Printf("error: X509KeyPair")
			return nil
		}
		certs[domain] = &certificate
		if net.ParseIP(domain) != nil {
			l = append(l, certificate)
		} else {
			fmt.Printf("error: %s", "IP parsing error")
		}
	}

	tc := &tls.Config{
		Certificates: l,
		GetCertificate: func(c *tls.ClientHelloInfo) (*tls.Certificate, error) {
			v, ok := certs[c.ServerName]
			if ok {
				return v, nil
			}
			fmt.Printf("error: %s%s", "Not found ", c.ServerName)
			return nil, errors.New("Not found " + c.ServerName)
		},
	}

	return tc
}
