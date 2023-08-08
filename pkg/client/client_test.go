package client_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"mini-redis-go/pkg/client"
	"mini-redis-go/pkg/config"
	"mini-redis-go/pkg/core/redis"
	"mini-redis-go/pkg/server"
	"os"
	"path/filepath"
	"testing"
)

func createTempFolder() string {
	folder, err := os.MkdirTemp("", "mini-redis")
	if err != nil {
		log.Fatal("error creating temp folder", err)
	}
	return folder
}

func startServer(host, port, cacheFolder string) server.MiniRedisServer {
	s := server.NewServer(host, port, cacheFolder)
	go s.Start()
	return s
}

func connectToServer(host, port string) client.MiniRedisClient {
	c := client.NewClient(host, port, config.ClientPublicKeyFile, config.ClientPrivateKeyFile)
	c.Connect()
	return c
}

func set(t *testing.T, c client.MiniRedisClient, k string, v string) {
	if err := c.Set(k, v); err != nil {
		t.Error("Error set", err)
	}
}

func get(t *testing.T, c client.MiniRedisClient, k string) string {
	response, err := c.Get(k)
	if err != nil {
		t.Error("Error get", err)
	}
	return response
}

func write(t *testing.T, c client.MiniRedisClient, message []byte) {
	if err := c.Write(message); err != nil {
		t.Error("Error sending a message", err)
	}
}

func read(t *testing.T, c client.MiniRedisClient) string {
	response, err := c.Read()
	if err != nil {
		t.Error("Error reading from server", err)
	}
	return response
}

func TestPingPong(t *testing.T) {
	tempFolder := createTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	s := startServer(config.ConnectionHost, config.ConnectionPort, tempFolder)
	c := connectToServer(config.ConnectionHost, config.ConnectionPort)

	write(t, c, []byte("PING\n"))
	response := read(t, c)
	assert.Equal(t, "PONG\n", response)

	s.Stop()
}

func TestSetAndGet(t *testing.T) {
	tempFolder := createTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	s := startServer(config.ConnectionHost, config.ConnectionPort, tempFolder)
	c := connectToServer(config.ConnectionHost, config.ConnectionPort)

	set(t, c, "hello", "world")
	response1 := read(t, c)
	assert.Equal(t, "Set ok\n", response1)

	response2 := get(t, c, "hello")
	assert.Equal(t, "world\n", response2)

	s.Stop()
}

func TestCache(t *testing.T) {
	tempFolder := createTempFolder()
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempFolder)
	fmt.Println("tempFolder", tempFolder)
	createCacheFileWithData(tempFolder, "testKey", append([]byte{byte(redis.StringByteType)}, []byte("tesValue")...))
	s := startServer(config.ConnectionHost, config.ConnectionPort, tempFolder)
	c := connectToServer(config.ConnectionHost, config.ConnectionPort)

	response := get(t, c, "testKey")
	assert.Equal(t, "tesValue\n", response)

	s.Stop()
}

//func TestSubscriberAndPublish(t *testing.T) {
//	tempFolder := createTempFolder()
//	defer func(path string) {
//		_ = os.RemoveAll(path)
//	}(tempFolder)
//	s := startServer(config.ConnectionHost, config.ConnectionPort, tempFolder)
//	client1 := connectToServer(config.ConnectionHost, config.ConnectionPort)
//	client3 := connectToServer(config.ConnectionHost, config.ConnectionPort)
//
//	//write(t, conn, "set hello world\n")
//	//response1 := read(t, conn)
//	//assert.Equal(t, "Set ok\n", response1)
//	//
//	//write(t, conn, "get hello\n")
//	//response2 := read(t, conn)
//	//assert.Equal(t, "world\n", response2)
//
//	s.Stop()
//}

func createCacheFileWithData(folder string, k string, v []byte) {
	file, _ := os.OpenFile(filepath.Join(folder, k), os.O_CREATE|os.O_WRONLY, 0644)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, _ = file.Write(v)
}
