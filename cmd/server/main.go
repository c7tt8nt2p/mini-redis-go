package main

import (
	"bufio"
	"fmt"
	"io"
	"mini-redis-go/pkg/config"
	"net"
	"os"
	"strings"
)

func main() {
	// initialize a listener
	listener, err := net.Listen("tcp", config.ConnectionHost+":"+config.ConnectionPort)
	if err != nil {
		fmt.Println("Error when initialize a connection:", err.Error())
		os.Exit(1)
	}
	// to close the listener
	defer func(listener net.Listener) {
		fmt.Println("Closing the listener...")
		err := listener.Close()
		if err != nil {
			fmt.Println("Error when closing a listener:", err.Error())
		}
	}(listener)
	// listener started
	fmt.Println("Server started...", config.ConnectionHost+":"+config.ConnectionPort)
	for {
		// incoming connection.
		connection, err := listener.Accept()
		fmt.Println("Incoming connection from:", connection.RemoteAddr())
		if err != nil {
			fmt.Println("Error when accepting a new connection: ", err.Error())
			os.Exit(1)
		}
		// handler
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			fmt.Println("Error when closing a connection:", err.Error())
		}
	}(connection)

	reader := bufio.NewReader(connection)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Goodbye", connection.RemoteAddr())
				break
			}
			fmt.Println("Error reading:", err.Error())
		}
		if strings.TrimSpace(message) == "exit" {
			fmt.Println("Bye", connection.RemoteAddr())
			break
		}
		fmt.Print("\t", connection.RemoteAddr().String()+" : ", message)
	}
}
