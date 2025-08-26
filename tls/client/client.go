package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type (
	Transport struct {
		Adress string
		Tls    *Tls
	}

	Tls struct {
		MinVersion             string
		CipherSuite            []uint16
		InsecureSkipVerify     bool
		HighSecureRejectDvCert bool
		RootCAs                *x509.CertPool
	}
)

const (
	VersionTLS10 = "TLS 1.0"
	VersionTLS11 = "TLS 1.1"
	VersionTLS12 = "TLS 1.2"
	VersionTLS13 = "TLS 1.3"
)

func (t *Tls) TLSClientConfig() *tls.Config {
	if t == nil {
		return &tls.Config{}
	}

	return &tls.Config{
		RootCAs:            t.RootCAs,
		InsecureSkipVerify: t.InsecureSkipVerify,
		MinVersion:         tlsVersionFromStr(t.MinVersion),
		CipherSuites:       t.cipherSuite(),
		VerifyConnection:   t.verifyConnection(),
	}
}

func (t *Tls) cipherSuite() []uint16 {
	if len(t.CipherSuite) == 0 {
		return nil
	}

	return t.CipherSuite
}

func tlsVersionFromStr(version string) uint16 {
	switch version {
	case "":
		return tls.VersionTLS12 //default for tls-conf in go
	case VersionTLS13:
		return tls.VersionTLS13
	case VersionTLS12:
		return tls.VersionTLS12
	case VersionTLS11:
		return tls.VersionTLS11
	case VersionTLS10:
		return tls.VersionTLS10
	default:
		return tls.VersionTLS12
	}
}

func (t *Tls) verifyConnection() func(cs tls.ConnectionState) error {
	return func(cs tls.ConnectionState) error {
		if err := t.validateTypeCert(cs); err != nil {
			return err
		}

		return nil
	}
}

func (t *Tls) validateTypeCert(cs tls.ConnectionState) error {
	if t.HighSecureRejectDvCert {
		if len(cs.PeerCertificates) == 0 {
			return fmt.Errorf("no trusted certificates provided for validation")
		}

		if len(cs.PeerCertificates[0].Subject.Organization) == 0 {
			return fmt.Errorf("detected DV-certificate. Expected OV or EV-certificate")
		}
	}

	return nil
}

func main() {
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	transportConfig := &Transport{
		Adress: "https://localhost:8080",
		Tls: &Tls{
			MinVersion:             "TLS 1.2",
			CipherSuite:            []uint16{},
			InsecureSkipVerify:     false,
			HighSecureRejectDvCert: true,
			RootCAs:                caCertPool,
		},
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: transportConfig.Tls.TLSClientConfig(),
		},
	}

	u, _ := url.Parse(transportConfig.Adress)
	log.Printf("→ %s", u.String())

	resp, err := client.Get(transportConfig.Adress)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer resp.Body.Close()

	fmt.Println("Status response:", resp.Status)
}
