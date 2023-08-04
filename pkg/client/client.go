package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"mini-redis-go/pkg/config"
	"net"
	"os"
)

type MiniRedisClient interface {
	Connect() *net.Conn
}

type Client struct {
	addr string
}

func NewClient(host, port string) *Client {
	c := Client{
		addr: host + ":" + port,
	}
	return &c
}

func (c *Client) Connect() *tls.Conn {
	certPool := loadCert()
	tlsConfig := &tls.Config{
		RootCAs:            certPool,
		InsecureSkipVerify: true, // Set to false for production use, when you have a valid certificate.
	}

	conn, err := tls.Dial("tcp", c.addr, tlsConfig)
	if err != nil {
		fmt.Println("Error when connecting to a server:", err.Error())
		os.Exit(1)
	}
	return conn
}

func loadCert() *x509.CertPool {
	cert, err := os.ReadFile(config.PublicKeyFile)
	if err != nil {
		panic(fmt.Sprintf("Error loading certificate: %s", err))
	}

	certPool := x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM(cert)
	if !ok {
		panic("Failed to parse server certificate")
	}

	return certPool
}
