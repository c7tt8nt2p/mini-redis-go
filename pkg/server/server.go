package server

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"mini-redis-go/pkg/config"
	"mini-redis-go/pkg/core"
	"net"
	"os"
)

type MiniRedisServer interface {
	Start() *net.Listener
}

type Server struct {
	Addr        string
	CacheFolder string
}

func NewServer(host, port, cacheFolder string) *Server {
	s := Server{
		Addr:        host + ":" + port,
		CacheFolder: cacheFolder,
	}
	return &s

}

func (s *Server) Start() {
	cert := s.loadCert()
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*cert},
		ClientAuth:   tls.RequireAnyClientCert,
	}
	listener, err := tls.Listen("tcp", s.Addr, tlsConfig)
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
	core.InitOnce(s.CacheFolder, config.CacheFileName)
	readCache(s.CacheFolder, config.CacheFileName)
	fmt.Println("Server started...", s.Addr)

	for {
		// incoming connection.
		conn := acceptANewConnection(&listener)
		go handleConnection(*s, *conn)
	}
}

func (s *Server) loadCert() *tls.Certificate {
	cert, err := tls.LoadX509KeyPair(config.PublicKeyFile, config.PrivateKeyFile)
	if err != nil {
		panic(fmt.Sprintf("Error loading certificate: %s", err))
	}
	return &cert
}

func acceptANewConnection(listener *net.Listener) *net.Conn {
	conn, err := (*listener).Accept()
	fmt.Println("Incoming connection from:", conn.RemoteAddr())
	if err != nil {
		fmt.Println("Error when accepting a new connection: ", err.Error())
		os.Exit(1)
	}
	return &conn
}

func handleConnection(s Server, conn net.Conn) {
	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			fmt.Println("Error when closing a connection:", err.Error())
		}
	}(conn)

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				fmt.Println("Goodbye", conn.RemoteAddr())
			} else {
				fmt.Println("Error reading:", err.Error())
			}
			break
		}

		cmdType := parse(message)
		switch cmdType {
		case exitCmd:
			fmt.Println("Bye", conn.RemoteAddr())
			break
		case pingCmd:
			_, err = conn.Write([]byte("PONG\n"))
			if err != nil {
				fmt.Println("Error sending response to pingCmd:", err)
				break
			}
			continue
		case setCmd:
			k, v := extractSetCli(message)
			myRedis := core.GetMyRedis()

			if myRedis.Exists(k) {
				myRedis.Set(k, v)
				cacheRewrite(&myRedis, s.CacheFolder)
			} else {
				myRedis.Set(k, v)
				cacheAppend(s.CacheFolder, k, v)
			}
			_, err = conn.Write([]byte("Set ok" + "\n"))
			if err != nil {
				fmt.Println("Error sending response to setCmd:", err)
				_ = conn.Close()
			}
		case getCmd:
			k := extractGetCli(message)
			myRedis := core.GetMyRedis()
			v := myRedis.Get(k)

			_, err = conn.Write([]byte(v + "\n"))
			if err != nil {
				fmt.Println("Error sending response to getCmd:", err)
				break
			}
		case otherCmd:
			_, err = conn.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Println("Error sending response to otherCmd:", err)
				break
			}

		}
		fmt.Print("\t", conn.RemoteAddr().String()+" : ", message)
	}
}
