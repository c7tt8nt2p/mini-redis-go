// Package client contains all business logic in server side
package client

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"io"
	"mini-redis-go/internal/config"
	"mini-redis-go/internal/utils"
	"net"
	"os"
)

// ClientService this is a client API interface that contains functions that use can do and interact with the Redis
type ClientService interface {
	Connect() io.ReadWriteCloser
	ConnectNonSecure() io.ReadWriteCloser
	GetConnection() io.ReadWriteCloser
	Set(k string, v string) error
	Get(k string) (string, error)
	Write(m []byte) error
	Read() (string, error)
	Subscribe(topic string) ISubscriber
}

type clientService struct {
	config         *config.ClientConfig
	addr           string
	publicKeyFile  string
	privateKeyFile string
	conn           io.ReadWriteCloser
}

func NewClientService(clientConfig *config.ClientConfig) ClientService {
	return &clientService{
		config:         clientConfig,
		addr:           clientConfig.Host + ":" + clientConfig.Port,
		publicKeyFile:  clientConfig.PublicKeyFile,
		privateKeyFile: clientConfig.PrivateKeyFile,
	}
}

func (c *clientService) Connect() io.ReadWriteCloser {
	cert := utils.LoadCertificate(c.config.PublicKeyFile, c.config.PrivateKeyFile)
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

func (c *clientService) ConnectNonSecure() io.ReadWriteCloser {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		fmt.Println("Error when connecting to a server:", err.Error())
		os.Exit(1)
	}
	c.conn = conn
	return conn
}

func (c *clientService) GetConnection() io.ReadWriteCloser {
	return c.conn
}

func (c *clientService) Set(k string, v string) error {
	msg := fmt.Sprintf("set %s %s\n", k, v)
	return c.Write([]byte(msg))
}

func (c *clientService) Get(k string) (string, error) {
	msg := fmt.Sprintf("get %s\n", k)
	if err := c.Write([]byte(msg)); err != nil {
		return "", err
	}
	return c.Read()
}

func (c *clientService) Write(m []byte) error {
	_, err := c.conn.Write(m)
	return err
}

func (c *clientService) Read() (string, error) {
	buf := make([]byte, 1024)
	n, err := c.conn.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func (c *clientService) Subscribe(topic string) ISubscriber {

	msg := fmt.Sprintf("SUBSCRIBE %s\n", topic)
	utils.WriteToServer(c.conn, msg)

	return &Subscriber{
		clientService: c,
	}
}
