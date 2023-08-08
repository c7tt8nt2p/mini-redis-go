package client

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"mini-redis-go/pkg/config"
	"mini-redis-go/pkg/utils"
	"os"
)

type MiniRedisClient interface {
	Connect() *tls.Conn
	Subscribe(topic string) MiniRedisSubscriber
	Publish(s string)
}

type Client struct {
	addr           string
	publicKeyFile  string
	privateKeyFile string
	conn           *tls.Conn
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
	c.conn = conn
	return conn
}

func (c *Client) Subscribe(topic string) MiniRedisSubscriber {
	msg := fmt.Sprintf("SUBSCRIBE %s\n", topic)
	utils.WriteToServer(c.conn, msg)

	return Subscriber{}
}

func (c *Client) Publish(s string) {
	//TODO implement me
	panic("implement me")
}

func loadCert() *tls.Certificate {
	cert, err := tls.LoadX509KeyPair(config.ClientPublicKeyFile, config.ClientPrivateKeyFile)
	if err != nil {
		panic(fmt.Sprintf("Error loading certificate: %s", err))
	}
	return &cert
}
