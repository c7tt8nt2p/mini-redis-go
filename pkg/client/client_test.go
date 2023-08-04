package client_test

import (
	"github.com/stretchr/testify/assert"
	"log"
	"mini-redis-go/pkg/client"
	"mini-redis-go/pkg/config"
	"mini-redis-go/pkg/server"
	"net"
	"os"
	"testing"
)

func createTempFolder() string {
	folder, err := os.MkdirTemp("", "mini-redis")
	if err != nil {
		log.Fatal("Error creating temp folder", err)
	}
	return folder
}

func startServer(host, port, cacheFolder string) {
	s := server.NewServer(host, port, cacheFolder)
	go s.Start()
}

func connectToServer(host, port string) *net.Conn {
	c := client.NewClient(host, port)
	conn := c.Connect()
	return conn
}

func write(t *testing.T, conn *net.Conn, s string) {
	if _, err1 := (*conn).Write([]byte(s)); err1 != nil {
		t.Error("Error sending a message", err1)
	}
}

func read(t *testing.T, conn *net.Conn) string {
	buf := make([]byte, 1024)
	n, err2 := (*conn).Read(buf)
	if err2 != nil {
		t.Error("Error reading response", err2)
	}
	return string(buf[:n])
}

func TestPingPong(t *testing.T) {
	tempFolder := createTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	startServer(config.ConnectionHost, config.ConnectionPort, tempFolder)
	conn := connectToServer(config.ConnectionHost, config.ConnectionPort)

	write(t, conn, "PING\n")
	response := read(t, conn)

	assert.Equal(t, "PONG\n", response)
}

func TestSet(t *testing.T) {
	tempFolder := createTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	startServer(config.ConnectionHost, config.ConnectionPort, tempFolder)
	conn := connectToServer(config.ConnectionHost, config.ConnectionPort)

	write(t, conn, "set hello world\n")
	response := read(t, conn)

	assert.Equal(t, "Set ok\n", response)
}

func TestSetAndGet(t *testing.T) {
	tempFolder := createTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	startServer(config.ConnectionHost, config.ConnectionPort, tempFolder)
	conn := connectToServer(config.ConnectionHost, config.ConnectionPort)

	write(t, conn, "set hello world\n")
	response1 := read(t, conn)
	assert.Equal(t, "Set ok\n", response1)

	write(t, conn, "get hello\n")
	response2 := read(t, conn)
	assert.Equal(t, "world\n", response2)
}
