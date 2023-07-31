package main

import (
	"bufio"
	"fmt"
	"mini-redis-go/pkg/config"
	"net"
	"os"
)

func main() {
	// initialize a connection
	connection, err := net.Dial("tcp", config.ConnectionHost+":"+config.ConnectionPort)
	if err != nil {
		fmt.Println("Error when connecting to a server:", err.Error())
		os.Exit(1)
	}
	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			fmt.Println("Error when closing a connection:", err.Error())
		}
	}(connection)

	fmt.Println("Connected to the server.")
	reader := bufio.NewReader(os.Stdin)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading a string:", err.Error())
			os.Exit(1)
		}
		if len(message) > 0 {
			_, err := connection.Write([]byte(message))
			if err != nil {
				fmt.Println("Error sending message to server:", err.Error())
				os.Exit(1)
			}
		}
	}
}
