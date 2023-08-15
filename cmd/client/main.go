package main

import (
	"crypto/tls"
	"fmt"
	"mini-redis-go/internal/config"
	"os"
	"strings"
)

func GetClientConfig() *config.ClientConfig {
	return &config.ClientConfig{
		ClientPublicKeyFile:  "/Users/chantapat.t/GolandProjects/mini-redis-go/internal/config/ssl/client/client.pem",
		ClientPrivateKeyFile: "/Users/chantapat.t/GolandProjects/mini-redis-go/internal/config/ssl/client/client.key",
		ConnectionHost:       "localhost",
		ConnectionPort:       "6973",
	}
}

func main() {
	myApp := InitializeClient()
	conn := myApp.ConnectToServer()
	defer func(connection *tls.Conn) {
		err := connection.Close()
		if err != nil {
			fmt.Println("Error when closing a connection:", err.Error())
		}
	}(conn)
	fmt.Println("Connected to the server")

	myApp.OnMessageReceivedFromServer(handleMessagesFromServer)
	myApp.OnMessageReceivedFromClient(handleMessagesFromClient)
}

func handleMessagesFromServer(messageFromServer string) {
	fmt.Println("Server says: ", messageFromServer)
}

func handleMessagesFromClient(messageFromClient string) {
	if len(messageFromClient) > 0 {
		if strings.TrimSpace(messageFromClient) == "exit" {
			fmt.Println("Disconnected.")
			os.Exit(0)
		}
	}
}
