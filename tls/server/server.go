package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
)

func main() {
	certFileName, keyFileName, err := generateSelfSignedCert()
	if err != nil {
		log.Fatalf("failed to generate self-signed certificate: %v", err)
	}

	server := &http.Server{
		Addr: ":8080",
		TLSConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("get request with tls-cipher: 0x%x", r.TLS.CipherSuite)
	})

	log.Println("Server launched on port: ", server.Addr)
	log.Fatal(server.ListenAndServeTLS(certFileName, keyFileName))
}

func generateSelfSignedCert() (string, string, error) {
	certTemplate, err := certTemplate()
	if err != nil {
		return "", "", err
	}

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate key: %v", err)
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, certTemplate, certTemplate, &priv.PublicKey, priv)
	if err != nil {
		return "", "", fmt.Errorf("failed to create certificate: %v", err)
	}

	certFile, err := os.Create("cert.pem")
	if err != nil {
		return "", "", fmt.Errorf("failed to create cert.pem: %v", err)
	}
	defer certFile.Close()

	if err := pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return "", "", fmt.Errorf("failed to write cert.pem: %v", err)
	}

	keyFile, err := os.Create("key.pem")
	if err != nil {
		return "", "", fmt.Errorf("failed to create key.pem: %v", err)
	}
	defer keyFile.Close()

	keyBytes := x509.MarshalPKCS1PrivateKey(priv)
	if err := pem.Encode(keyFile, &pem.Block{Type: "PRIVATE KEY", Bytes: keyBytes}); err != nil {
		return "", "", fmt.Errorf("failed to write key.pem: %v", err)
	}

	return "cert.pem", "key.pem", nil
}

func certTemplate() (*x509.Certificate, error) {

	serialNumber, err := rand.Int(rand.Reader, big.NewInt(1<<62))
	if err != nil {
		return nil, fmt.Errorf("failed to generate serial number: %v", err)
	}

	return &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"go-developer-company"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{"localhost123"},
	}, nil
}
