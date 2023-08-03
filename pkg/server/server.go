package server

import (
	"bufio"
	"fmt"
	"io"
	"mini-redis-go/pkg/core"
	"net"
	"os"
)

type MiniRedisServer interface {
	Start() *net.Listener
}

type Server struct {
	Addr string
}

func NewServer(host, port string) *Server {
	s := Server{
		Addr: host + ":" + port,
	}
	return &s

}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		fmt.Println("Error when initialize a connection:", err.Error())
		os.Exit(1)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("Error when closing a listener:", err.Error())
		}
	}(listener)
	fmt.Println("Server started...", s.Addr)

	for {
		// incoming connection.
		connection := acceptANewConnection(&listener)
		go handleConnection(*connection)
	}
}

func acceptANewConnection(listener *net.Listener) *net.Conn {
	connection, err := (*listener).Accept()
	fmt.Println("Incoming connection from:", connection.RemoteAddr())
	if err != nil {
		fmt.Println("Error when accepting a new connection: ", err.Error())
		os.Exit(1)
	}
	return &connection
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

		cmdType := parse(message)
		switch cmdType {
		case exitCmd:
			fmt.Println("Bye", connection.RemoteAddr())
			break
		case pingCmd:
			_, err = connection.Write([]byte("PONG\n"))
			if err != nil {
				fmt.Println("Error sending response:", err)
				break
			}
			continue
		case setCmd:
			k, v := extractSetCli(message)
			myRedis := core.NewMyRedis()
			myRedis.Set(k, v)

			_, err = connection.Write([]byte("Set ok" + "\n"))
			if err != nil {
				fmt.Println("Error sending response:", err)
				break
			}
		case getCmd:
			k := extractGetCli(message)
			myRedis := core.NewMyRedis()
			v := myRedis.Get(k)

			_, err = connection.Write([]byte(v + "\n"))
			if err != nil {
				fmt.Println("Error sending response:", err)
				break
			}
		case otherCmd:
			_, err = connection.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Println("Error sending response:", err)
				break
			}

		}
		fmt.Print("\t", connection.RemoteAddr().String()+" : ", message)
	}
}
