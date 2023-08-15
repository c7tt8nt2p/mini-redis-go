package utils

import (
	"mini-redis-go/internal/config"
	"mini-redis-go/internal/service/client"
	"mini-redis-go/internal/service/server"
	"testing"
)

func StartServer(host, port, cacheFolder string) server.IServer {
	s := server.NewServerService(host, port, cacheFolder)
	go s.Start()
	return s
}

func GetClientConfigTest() *config.ClientConfig {
	return &config.ClientConfig{
		ClientPublicKeyFile:  "/Users/chantapat.t/GolandProjects/mini-redis-go/internal/config/ssl/client/client.pem",
		ClientPrivateKeyFile: "/Users/chantapat.t/GolandProjects/mini-redis-go/internal/config/ssl/client/client.key",
		ConnectionHost:       "localhost",
		ConnectionPort:       "6973",
	}
}

func ConnectToServer(clientConfig *config.ClientConfig) client.IClient {
	c := client.NewClientService(clientConfig)
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
