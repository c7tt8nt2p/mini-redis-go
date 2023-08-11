package app

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"mini-redis-go/internal/service/client"
	"mini-redis-go/internal/utils"
	"os"
)

// IClientApp is an entrypoint when instantiating a new client
type IClientApp interface {
	ConnectToServer() *tls.Conn
	// OnMessageReceivedFromServer register a function to handles messages from the server
	OnMessageReceivedFromServer(handlerFunc func(messageFromServer string))
	// OnMessageReceivedFromClient register a function to handles messages to the server
	OnMessageReceivedFromClient(handlerFunc func(messageFromClient string))
}

type ClientApp struct {
	clientService client.IClient
}

func NewClientApp(clientService client.IClient) *ClientApp {
	return &ClientApp{
		clientService: clientService,
	}
}

func (c *ClientApp) ConnectToServer() *tls.Conn {
	return c.clientService.Connect()
}

func (c *ClientApp) OnMessageReceivedFromServer(handlerFunc func(messageFromServer string)) {
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, err := c.clientService.GetConnection().Read(buffer)
			if err != nil {
				fmt.Println("Error reading server response:", err)
				break
			}
			message := string(buffer[:n])

			handlerFunc(message)
		}
	}()
}

func (c *ClientApp) OnMessageReceivedFromClient(handlerFunc func(messageFromClient string)) {
	reader := bufio.NewReader(os.Stdin)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading a string:", err.Error())
			os.Exit(1)
		}

		handlerFunc(message)

		utils.WriteToServer(c.clientService.GetConnection(), message)
	}
}
