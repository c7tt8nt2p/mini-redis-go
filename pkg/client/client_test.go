package client_test

import (
	"github.com/stretchr/testify/assert"
	"mini-redis-go/pkg/client"
	"mini-redis-go/pkg/server"
	"testing"
)

func TestPingPong(t *testing.T) {
	go server.StartServer()

	connection := client.StartConnection()
	_, err1 := (*connection).Write([]byte("PING\n"))
	if err1 != nil {
		t.Error("Error sending a message")
	}

	buf := make([]byte, 1024)
	n, err2 := (*connection).Read(buf)
	if err2 != nil {
		t.Fatalf("Error reading response: %v", err2)
	}

	response := string(buf[:n])
	assert.Equal(t, "PONG\n", response)
}
