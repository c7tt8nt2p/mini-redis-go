package client

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"mini-redis-go/pkg/config"
	"net"
	"os"
)

type MiniRedisClient interface {
	Connect() *net.Conn
}

type Client struct {
	addr           string
	publicKeyFile  string
	privateKeyFile string
}

func NewClient(host, port, publicKeyFile, privateKeyFile string) *Client {
	c := Client{
		addr:           host + ":" + port,
		publicKeyFile:  publicKeyFile,
		privateKeyFile: privateKeyFile,
	}
	return &c
}

func (c *Client) Connect() *tls.Conn {
	cert := loadCert()
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{*cert}, InsecureSkipVerify: true}
	tlsConfig.Rand = rand.Reader

	conn, err := tls.Dial("tcp", c.addr, tlsConfig)
	if err != nil {
		fmt.Println("Error when connecting to a server:", err.Error())
		os.Exit(1)
	}
	return conn
}

func loadCert() *tls.Certificate {
	cert, err := tls.LoadX509KeyPair(config.ClientPublicKeyFile, config.ClientPrivateKeyFile)
	if err != nil {
		panic(fmt.Sprintf("Error loading certificate: %s", err))
	}
	return &cert
}
