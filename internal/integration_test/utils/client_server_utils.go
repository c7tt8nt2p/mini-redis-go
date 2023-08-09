package utils

import (
	"log"
	"mini-redis-go/internal/config"
	"mini-redis-go/internal/service/client"
	"mini-redis-go/internal/service/server"
	"os"
	"testing"
)

func CreateTempFolder() string {
	folder, err := os.MkdirTemp("", "mini-redis")
	if err != nil {
		log.Fatal("error creating temp folder", err)
	}
	return folder
}

func StartServer(host, port, cacheFolder string) server.IServer {
	s := server.NewServer(host, port, cacheFolder)
	go s.Start()
	return s
}

func ConnectToServer(host, port string) client.IClient {
	c := client.NewClient(host, port, config.ClientPublicKeyFile, config.ClientPrivateKeyFile)
	c.Connect()
	return c
}

func Set(t *testing.T, c client.IClient, k string, v string) {
	if err := c.Set(k, v); err != nil {
		t.Error("Error set", err)
	}
}

func Get(t *testing.T, c client.IClient, k string) string {
	response, err := c.Get(k)
	if err != nil {
		t.Error("Error get", err)
	}
	return response
}

func Write(t *testing.T, c client.IClient, message []byte) {
	if err := c.Write(message); err != nil {
		t.Error("Error sending a message", err)
	}
}

func Read(t *testing.T, c client.IClient) string {
	response, err := c.Read()
	if err != nil {
		t.Error("Error reading from server", err)
	}
	return response
}

func NextMessage(t *testing.T, s client.ISubscriber) string {
	msg, err := s.NextMessage()
	if err != nil {
		t.Error("Error next message from subscriber", err)
	}
	return msg
}

func Publish(t *testing.T, s client.ISubscriber, msg string) {
	err := s.Publish(msg)
	if err != nil {
		t.Error("Error publishing a message", err)
	}
}

func Unsubscribe(t *testing.T, s client.ISubscriber) {
	err := s.Unsubscribe()
	if err != nil {
		t.Error("Error unsubscribing", err)
	}
}
