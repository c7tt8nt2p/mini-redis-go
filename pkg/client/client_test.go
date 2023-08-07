package client_test

import (
	"crypto/tls"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"mini-redis-go/pkg/client"
	"mini-redis-go/pkg/config"
	"mini-redis-go/pkg/server"
	"os"
	"path/filepath"
	"testing"
)

func createTempFolder() string {
	folder, err := os.MkdirTemp("", "mini-redis")
	if err != nil {
		log.Fatal("Error creating temp folder", err)
	}
	return folder
}

func startServer(host, port, cacheFolder string) *server.Server {
	s := server.NewServer(host, port, cacheFolder)
	go s.Start()
	return s
}

func connectToServer(host, port string) *tls.Conn {
	c := client.NewClient(host, port, config.ClientPublicKeyFile, config.ClientPrivateKeyFile)
	conn := c.Connect()
	return conn
}

func write(t *testing.T, conn *tls.Conn, s string) {
	if _, err1 := (*conn).Write([]byte(s)); err1 != nil {
		t.Error("Error sending a message", err1)
	}
}

func read(t *testing.T, conn *tls.Conn) string {
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
	s := startServer(config.ConnectionHost, config.ConnectionPort, tempFolder)
	conn := connectToServer(config.ConnectionHost, config.ConnectionPort)

	write(t, conn, "PING\n")
	response := read(t, conn)

	assert.Equal(t, "PONG\n", response)

	s.Stop()
}

func TestSetAndGet(t *testing.T) {
	tempFolder := createTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	s := startServer(config.ConnectionHost, config.ConnectionPort, tempFolder)
	conn := connectToServer(config.ConnectionHost, config.ConnectionPort)

	write(t, conn, "set hello world\n")
	response1 := read(t, conn)
	assert.Equal(t, "Set ok\n", response1)

	write(t, conn, "get hello\n")
	response2 := read(t, conn)
	assert.Equal(t, "world\n", response2)

	s.Stop()
}

func TestCache(t *testing.T) {
	tempFolder := createTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	fmt.Println("tempFolder", tempFolder)
	createCacheFileWithData(tempFolder, "testKey", []byte("tesValue"))
	s := startServer(config.ConnectionHost, config.ConnectionPort, tempFolder)
	conn := connectToServer(config.ConnectionHost, config.ConnectionPort)

	write(t, conn, "get testKey\n")
	response2 := read(t, conn)
	assert.Equal(t, "tesValue\n", response2)

	s.Stop()
}

func createCacheFileWithData(folder string, k string, v []byte) {
	file, _ := os.OpenFile(filepath.Join(folder, k), os.O_CREATE|os.O_WRONLY, 0644)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, _ = file.Write(v)
}
