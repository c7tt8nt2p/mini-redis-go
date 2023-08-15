// Package client contains all business logic in server side
package client

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"mini-redis-go/internal/config"
	"mini-redis-go/internal/utils"
	"os"
)

// IClient this is a client API interface that contains functions that use can do and interact with the Redis
type IClient interface {
	Connect() *tls.Conn
	GetConnection() *tls.Conn
	Set(k string, v string) error
	Get(k string) (string, error)
	Write(m []byte) error
	Read() (string, error)
	Subscribe(topic string) ISubscriber
}

type ClientService struct {
	addr           string
	publicKeyFile  string
	privateKeyFile string
	conn           *tls.Conn
}

func NewClientService(clientConfig *config.ClientConfig) IClient {
	return &ClientService{
		addr:           clientConfig.ConnectionHost + ":" + clientConfig.ConnectionPort,
		publicKeyFile:  clientConfig.ClientPublicKeyFile,
		privateKeyFile: clientConfig.ClientPrivateKeyFile,
	}
}

func (c *ClientService) Connect() *tls.Conn {
	cert := utils.LoadCertificate(config.ClientPublicKeyFile, config.ClientPrivateKeyFile)
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{*cert},
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS12,
		MaxVersion:         0,
	}
	tlsConfig.Rand = rand.Reader

	conn, err := tls.Dial("tcp", c.addr, tlsConfig)
	if err != nil {
		fmt.Println("Error when connecting to a server:", err.Error())
		os.Exit(1)
	}
	c.conn = conn
	return conn
}

func (c *ClientService) GetConnection() *tls.Conn {
	return c.conn
}

func (c *ClientService) Set(k string, v string) error {
	msg := fmt.Sprintf("set %s %s\n", k, v)
	return c.Write([]byte(msg))
}

func (c *ClientService) Get(k string) (string, error) {
	msg := fmt.Sprintf("get %s\n", k)
	if err := c.Write([]byte(msg)); err != nil {
		return "", err
	}
	return c.Read()
}

func (c *ClientService) Write(m []byte) error {
	_, err := (*c.conn).Write(m)
	return err
}

func (c *ClientService) Read() (string, error) {
	buf := make([]byte, 1024)
	n, err := (*c.conn).Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func (c *ClientService) Subscribe(topic string) ISubscriber {

	msg := fmt.Sprintf("SUBSCRIBE %s\n", topic)
	utils.WriteToServer(c.conn, msg)

	return &Subscriber{
		c: c,
	}
}
