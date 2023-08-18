package main

import (
	"flag"
	"fmt"
	"io"
	"mini-redis-go/internal/config"
	"os"
	"strings"
)

const (
	defaultPublicKeyFile  = "/Users/chantapat.t/GolandProjects/mini-redis-go/internal/config/ssl/client/client.pem"
	defaultPrivateKeyFile = "/Users/chantapat.t/GolandProjects/mini-redis-go/internal/config/ssl/client/client.key"
	defaultHost           = "localhost"
	defaultPort           = "6973"
)

func GetClientConfig() *config.ClientConfig {
	publicKeyFile := flag.String("publickey", defaultPublicKeyFile, "a client public key file")
	privateKeyFile := flag.String("privatekey", defaultPrivateKeyFile, "a client private key file")
	hostFlag := flag.String("host", defaultHost, "a connection host to connect to")
	portFlag := flag.String("port", defaultPort, "a connection port to connect to")
	flag.Parse()

	return &config.ClientConfig{
		PublicKeyFile:  *publicKeyFile,
		PrivateKeyFile: *privateKeyFile,
		Host:           *hostFlag,
		Port:           *portFlag,
	}
}

func main() {
	myApp := InitializeClient()
	conn := myApp.ConnectToServer()
	defer func(connection io.ReadWriteCloser) {
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
