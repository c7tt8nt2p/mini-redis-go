package main

import (
	"bufio"
	"fmt"
	"mini-redis-go/pkg/client"
	"mini-redis-go/pkg/config"
	"net"
	"os"
	"strings"
)

func main() {
	c := client.NewClient(config.ConnectionHost, config.ConnectionPort)
	conn := c.Connect()
	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			fmt.Println("Error when closing a connection:", err.Error())
		}
	}(*conn)
	fmt.Println("Connected to the server")
	go handleMessagesFromServer(conn)
	handleMessagesFromClient(conn)
}

func handleMessagesFromServer(connection *net.Conn) {
	buffer := make([]byte, 1024)

	for {
		n, err := (*connection).Read(buffer)
		if err != nil {
			fmt.Println("Error reading server response:", err)
			break
		}

		message := string(buffer[:n])
		fmt.Println("Server says: ", message)
	}
}

func handleMessagesFromClient(connection *net.Conn) {
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
			_, err := (*connection).Write([]byte(message))
			if err != nil {
				fmt.Println("Error sending message to server:", err.Error())
				os.Exit(1)
			}
		}
	}
}
