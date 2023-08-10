package main

import (
	"crypto/tls"
	"fmt"
	"mini-redis-go/internal/app"
	"mini-redis-go/internal/config"
	"mini-redis-go/internal/service/client"
	"os"
	"strings"
)

func main() {
	clientService := client.NewClientService(config.ConnectionHost, config.ConnectionPort, config.ClientPublicKeyFile, config.ClientPrivateKeyFile)

	myApp := app.NewClientApp(clientService)
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
