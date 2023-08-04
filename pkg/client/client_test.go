package client_test

import (
	"github.com/stretchr/testify/assert"
	"mini-redis-go/pkg/client"
	"mini-redis-go/pkg/config"
	"mini-redis-go/pkg/server"
	"testing"
)

func TestPingPong(t *testing.T) {
	s := server.NewServer(config.ConnectionHost, config.ConnectionPort)
	go s.Start()

	c := client.NewClient(config.ConnectionHost, config.ConnectionPort)
	connection := c.Connect()

	if _, err1 := (*connection).Write([]byte("PING\n")); err1 != nil {
		t.Error("Error sending a message", err1)
	}

	buf := make([]byte, 1024)
	n, err2 := (*connection).Read(buf)
	if err2 != nil {
		t.Error("Error reading response", err2)
	}

	response := string(buf[:n])
	assert.Equal(t, "PONG\n", response)
}
