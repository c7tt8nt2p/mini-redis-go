package client

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"mini-redis-go/internal/config"
	"mini-redis-go/internal/utils"
	"os"
)

type IClient interface {
	Connect() *tls.Conn
	GetConnection() *tls.Conn
	Set(k string, v string) error
	Get(k string) (string, error)
	Write(m []byte) error
	Read() (string, error)
	Subscribe(topic string) ISubscriber
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

func (c *Client) GetConnection() *tls.Conn {
	return c.conn
}

func (c *Client) Set(k string, v string) error {
	msg := fmt.Sprintf("set %s %s\n", k, v)
	return c.Write([]byte(msg))
}

func (c *Client) Get(k string) (string, error) {
	msg := fmt.Sprintf("get %s\n", k)
	if err := c.Write([]byte(msg)); err != nil {
		return "", err
	}
	return c.Read()
}

func (c *Client) Write(m []byte) error {
	_, err := (*c.conn).Write(m)
	return err
}

func (c *Client) Read() (string, error) {
	buf := make([]byte, 1024)
	n, err := (*c.conn).Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func (c *Client) Subscribe(topic string) ISubscriber {
	msg := fmt.Sprintf("SUBSCRIBE %s\n", topic)
	utils.WriteToServer(c.conn, msg)

	return &Subscriber{
		c: c,
	}
}

func loadCert() *tls.Certificate {
	cert, err := tls.LoadX509KeyPair(config.ClientPublicKeyFile, config.ClientPrivateKeyFile)
	if err != nil {
		panic(fmt.Sprintf("Error loading certificate: %s", err))
	}
	return &cert
}
