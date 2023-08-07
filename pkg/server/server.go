package server

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"mini-redis-go/pkg/config"
	"mini-redis-go/pkg/core"
	"net"
)

type MiniRedisServer interface {
	Start()
}

type Server struct {
	listener    *net.Listener
	Addr        string
	CacheFolder string
	stopSignal  chan bool
}

func NewServer(host, port, cacheFolder string) *Server {
	s := Server{
		Addr:        host + ":" + port,
		CacheFolder: cacheFolder,
		stopSignal:  make(chan bool, 1),
	}
	return &s

}

func (s *Server) Start() {
	cert := loadCert()
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*cert},
		ClientAuth:   tls.RequireAnyClientCert,
	}
	listener, err := tls.Listen("tcp", s.Addr, tlsConfig)
	if err != nil {
		log.Panic("Error when initialize a connection: ", err)
	}
	s.listener = &listener
	defer func(listener net.Listener) {
		_ = listener.Close()
	}(listener)

	core.InitMyRedis(s.CacheFolder, config.CacheFileName)
	readCache(s.CacheFolder, config.CacheFileName)
	fmt.Println("Server started...", s.Addr)

	go func() {
		stop := <-s.stopSignal
		if stop {
			fmt.Println("Stopping the server...")
			_ = (*s.listener).Close()
		}
	}()

	for {
		// incoming connection.
		conn, err := acceptANewConnection(&listener)
		if err != nil {
			break
		}
		go handleConnection(s, conn)
	}

}

func (s *Server) Stop() {
	s.stopSignal <- true
}

func loadCert() *tls.Certificate {
	cert, err := tls.LoadX509KeyPair(config.PublicKeyFile, config.PrivateKeyFile)
	if err != nil {
		log.Fatal("error loading certificate: ", err)
	}
	return &cert
}

func acceptANewConnection(listener *net.Listener) (*net.Conn, error) {
	conn, err := (*listener).Accept()
	if err != nil {
		return nil, err
	}
	fmt.Println("Incoming connection from:", conn.RemoteAddr())
	return &conn, nil
}

func handleConnection(server *Server, conn *net.Conn) {
	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			log.Println("Error when closing a connection: ", err.Error())
		}
	}(*conn)

	reader := bufio.NewReader(*conn)
	for {
		message, err := readMessage(reader, conn)
		if err != nil {
			break
		}

		cmdType := parse(message)

		switch cmdType {
		case exitCmd:
			exitCmdHandler((*conn).RemoteAddr().String())
		case pingCmd:
			if err := pingCmdHandler(conn); err != nil {
				log.Println("error sending response to pingCmd: ", err)
			}
		case setCmd:
			err := setCmdHandler(conn, server, message)
			if err != nil {
				log.Println("error sending response to setCmd: ", err)
				_ = (*conn).Close()
			}
		case getCmd:
			err := getCmdHandler(conn, message)
			if err != nil {
				log.Println("Error sending response to getCmd: ", err)
			}
		case otherCmd:
			err := otherCmdHandler(conn, message)
			if err != nil {
				log.Println("Error sending response to otherCmd: ", err)
			}
		}
		fmt.Print("\t", (*conn).RemoteAddr().String()+" : ", message)
	}
}

func readMessage(reader *bufio.Reader, conn *net.Conn) (string, error) {
	message, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			fmt.Println("goodbye", (*conn).RemoteAddr())
		} else {
			fmt.Println("error reading:", err.Error())
		}
		return "", err
	}
	return message, nil
}
