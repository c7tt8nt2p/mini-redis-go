package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"mini-redis-go/pkg/client"
	"mini-redis-go/pkg/config"
	"mini-redis-go/pkg/utils"
	"os"
	"strings"
)

func main() {
	c := client.NewClient(config.ConnectionHost, config.ConnectionPort, config.ClientPublicKeyFile, config.ClientPrivateKeyFile)
	conn := c.Connect()
	defer func(connection *tls.Conn) {
		err := connection.Close()
		if err != nil {
			fmt.Println("Error when closing a connection:", err.Error())
		}
	}(conn)
	fmt.Println("Connected to the server")
	go handleMessagesFromServer(conn)
	handleMessagesFromClient(conn)
}

func handleMessagesFromServer(conn *tls.Conn) {
	buffer := make([]byte, 1024)

	for {
		n, err := (*conn).Read(buffer)
		if err != nil {
			fmt.Println("Error reading server response:", err)
			break
		}

		message := string(buffer[:n])
		fmt.Println("Server says: ", message)
	}
}

func handleMessagesFromClient(conn *tls.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading a string:", err.Error())
			os.Exit(1)
		}
		if len(message) > 0 {
			if strings.TrimSpace(message) == "exit" {
				fmt.Println("Disconnected.")
				break
			}
			utils.WriteToServer(conn, message)
		}
	}
}
