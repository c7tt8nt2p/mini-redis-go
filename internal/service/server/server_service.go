// Package server contains all business logic in server side
package server

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"mini-redis-go/internal/config"
	"mini-redis-go/internal/service/server/core"
	"mini-redis-go/internal/service/server/handler"
	"mini-redis-go/internal/service/server/parser"
	"mini-redis-go/internal/utils"
	"net"
	"sync"
)

var serverServiceInstance *ServerService
var serverServiceMutex = &sync.Mutex{}

type IServer interface {
	Start()
	Stop()
	GetCacheFolder() string
	SetCacheFolder(cacheFolder string)
}

type ServerService struct {
	config            *config.ServerConfig
	redisService      core.IRedis
	brokerService     core.IBroker
	cmdHandlerService handler.ICmdHandler
	listener          *net.Listener
	Addr              string
	cacheFolder       string
	stopSignal        chan bool
}

func NewServerService(serverConfig *config.ServerConfig) IServer {
	if serverServiceInstance == nil {
		serverServiceMutex.Lock()
		defer serverServiceMutex.Unlock()
		if serverServiceInstance == nil {
			instance := &ServerService{
				config:            serverConfig,
				redisService:      core.NewRedisService(),
				brokerService:     core.NewBrokerService(),
				cmdHandlerService: handler.NewCmdHandlerService(),
				Addr:              serverConfig.Host + ":" + serverConfig.Port,
				cacheFolder:       serverConfig.CacheFolder,
				stopSignal:        make(chan bool, 1),
			}
			serverServiceInstance = instance
		}
	}
	return serverServiceInstance
}

func (s *ServerService) Start() {
	cert := utils.LoadCertificate(s.config.PublicKeyFile, s.config.PrivateKeyFile)
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*cert},
		ClientAuth:   tls.RequireAnyClientCert,
		MinVersion:   tls.VersionTLS12,
		MaxVersion:   tls.VersionTLS13,
	}
	listener, err := tls.Listen("tcp", s.Addr, tlsConfig)
	if err != nil {
		log.Panic("error when initialize a connection: ", err)
	}
	s.listener = &listener
	defer func(listener net.Listener) {
		_ = listener.Close()
	}(listener)

	s.listen(listener)
}

func (s *ServerService) listen(listener net.Listener) {
	log.Println("===============================================================================================")
	s.redisService.ReadCache(s.cacheFolder)
	log.Println("================================================================================================")
	log.Println("server is ready...")
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

func (s *ServerService) Stop() {
	s.stopSignal <- true
}

func (s *ServerService) GetCacheFolder() string {
	return s.cacheFolder
}

func (s *ServerService) SetCacheFolder(cacheFolder string) {
	s.cacheFolder = cacheFolder
}

func acceptANewConnection(listener *net.Listener) (net.Conn, error) {
	conn, err := (*listener).Accept()
	if err != nil {
		return nil, err
	}
	log.Println("incoming connection from: ", conn.RemoteAddr())
	return conn, nil
}

func handleConnection(serverService *ServerService, conn net.Conn) {
	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			log.Println("error when closing a connection: ", err.Error())
		}
	}(conn)

	reader := bufio.NewReader(conn)
	for {
		message, err := readMessage(serverService, reader, conn)
		if err != nil {
			break
		}

		subscriptionConnection := serverService.brokerService.IsSubscriptionConnection(conn)
		if subscriptionConnection {
			handleSubscriptionConnection(serverService, conn, message)
		} else {
			handleNonSubscriptionConnection(serverService, conn, message)
			fmt.Printf("\t[%s]: %s", conn.RemoteAddr().String(), message)
		}
	}
}

func handleSubscriptionConnection(serverService *ServerService, conn net.Conn, message string) {
	cmdType := parser.ParseSubscriptionCommand(message)
	switch cmdType {
	case parser.UnsubscribeCmd:
		serverService.cmdHandlerService.UnsubscribeCmdHandler(conn)
	case parser.PublishCmd:
		serverService.cmdHandlerService.PublishCmdHandler(conn, message)
	}
}

func handleNonSubscriptionConnection(serverService *ServerService, conn net.Conn, message string) {
	cmdType := parser.ParseNonSubscriptionCommand(message)

	switch cmdType {
	case parser.ExitCmd:
		serverService.cmdHandlerService.ExitCmdHandler(conn.RemoteAddr().String())
	case parser.PingCmd:
		if err := serverService.cmdHandlerService.PingCmdHandler(conn); err != nil {
			log.Println("error sending response to pingCmd: ", err)
		}
	case parser.SetCmd:
		err := serverService.cmdHandlerService.SetCmdHandler(conn, serverService.GetCacheFolder(), message)
		if err != nil {
			log.Println("error sending response to setCmd: ", err)
			_ = conn.Close()
		}
	case parser.GetCmd:
		err := serverService.cmdHandlerService.GetCmdHandler(conn, message)
		if err != nil {
			log.Println("Error sending response to getCmd: ", err)
		}
	case parser.SubscribeCmd:
		fmt.Println("ok SubscribeCmd")
		err := serverService.cmdHandlerService.SubscribeCmdHandler(conn, message)
		if err != nil {
			log.Println("Error subscribing response to getCmd: ", err)
		}
	case parser.OtherCmd:
		err := serverService.cmdHandlerService.OtherCmdHandler(conn, message)
		if err != nil {
			log.Println("Error sending response to otherCmd: ", err)
		}
	}
}

func readMessage(serverService *ServerService, reader *bufio.Reader, conn net.Conn) (string, error) {
	message, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			serverService.cmdHandlerService.UnsubscribeCmdHandler(conn)
			fmt.Println("goodbye", conn.RemoteAddr())
		} else {
			fmt.Println("error reading message from client:", err.Error())
		}
		return "", err
	}
	return message, nil
}
