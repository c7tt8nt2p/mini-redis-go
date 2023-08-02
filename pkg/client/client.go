package client

import (
	"fmt"
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

func (c *Client) Connect() *net.Conn {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		fmt.Println("Error when connecting to a server:", err.Error())
		os.Exit(1)
	}
	return &conn
}
