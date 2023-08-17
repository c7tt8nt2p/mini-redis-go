package utils

import (
	"crypto/tls"
	"log"
)

func LoadCertificate(publicKeyFile, privateKeyFile string) *tls.Certificate {
	cert, err := tls.LoadX509KeyPair(publicKeyFile, privateKeyFile)
	if err != nil {
		log.Printf("error loading certificate: %s", err)
	}
	return &cert
}
