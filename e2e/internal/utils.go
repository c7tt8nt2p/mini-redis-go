package internal

import (
	"mini-redis-go/internal/config"
	"mini-redis-go/internal/service/client"
	"mini-redis-go/internal/service/server"
	"testing"
)

func StartServer(cacheFolder string) server.IServer {
	serverConfig := GetServerConfigTest(cacheFolder)
	s := server.NewServerService(serverConfig)
	go s.StartNonSecure()
	return s
}

func GetClientConfigTest() *config.ClientConfig {
	return &config.ClientConfig{
		Host: "localhost",
		Port: "6973",
	}
}

func GetServerConfigTest(cacheFolder string) *config.ServerConfig {
	return &config.ServerConfig{
		Host:        "localhost",
		Port:        "6973",
		CacheFolder: cacheFolder,
	}
}

func ConnectToServer(clientConfig *config.ClientConfig) client.ClientService {
	c := client.NewClientService(clientConfig)
	c.ConnectNonSecure()
	return c
}

func Set(t *testing.T, c client.ClientService, k string, v string) {
	if err := c.Set(k, v); err != nil {
		t.Error("Error set", err)
	}
}

func Get(t *testing.T, c client.ClientService, k string) string {
	response, err := c.Get(k)
	if err != nil {
		t.Error("Error get", err)
	}
	return response
}

func Write(t *testing.T, c client.ClientService, message []byte) {
	if err := c.Write(message); err != nil {
		t.Error("Error sending a message", err)
	}
}

func Read(t *testing.T, c client.ClientService) string {
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
