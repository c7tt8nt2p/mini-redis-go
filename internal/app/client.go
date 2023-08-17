package app

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"mini-redis-go/internal/service/client"
	"mini-redis-go/internal/utils"
	"os"
)

// Client is an entrypoint when instantiating a new client
type Client interface {
	ConnectToServer() *tls.Conn
	// OnMessageReceivedFromServer register a function to handles messages from the server
	OnMessageReceivedFromServer(handlerFunc func(messageFromServer string))
	// OnMessageReceivedFromClient register a function to handles messages to the server
	OnMessageReceivedFromClient(handlerFunc func(messageFromClient string))
}

type ClientApp struct {
	clientService client.ClientService
}

func NewClientApp(clientService client.ClientService) *ClientApp {
	return &ClientApp{
		clientService: clientService,
	}
}

func (c *ClientApp) ConnectToServer() *tls.Conn {
	return c.clientService.Connect()
}

func (c *ClientApp) OnMessageReceivedFromServer(handlerFunc func(messageFromServer string)) {
	go c.handleMessageFromServer(handlerFunc)()
}

func (c *ClientApp) handleMessageFromServer(handlerFunc func(messageFromServer string)) func() {
	return func() {
		buffer := make([]byte, 1024)
		for {
			n, err := c.clientService.GetConnection().Read(buffer)
			if err != nil {
				if err == io.EOF || err == io.ErrUnexpectedEOF {
					_ = c.clientService.GetConnection().Close()
					os.Exit(0)
				}
				fmt.Println("Error reading server response:", err)
				break

			}
			message := string(buffer[:n])

			handlerFunc(message)
		}
	}
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
