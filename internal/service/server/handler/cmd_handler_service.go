package handler

import (
	"fmt"
	"mini-redis-go/internal/model"
	"mini-redis-go/internal/service/server/core"
	"mini-redis-go/internal/service/server/parser"
	"mini-redis-go/internal/utils"
	"net"
	"sync"
)

var serverCmdHandlerServiceInstance *CmdHandlerService
var serverCmdHandlerServiceMutex = &sync.Mutex{}

type ICmdHandler interface {
	ExitCmdHandler(addr string)
	PingCmdHandler(net.Conn) error
	SetCmdHandler(conn net.Conn, cacheFolder string, message string) error
	GetCmdHandler(conn net.Conn, message string) error
	SubscribeCmdHandler(conn net.Conn, message string) error
	OtherCmdHandler(conn net.Conn, message string) error

	UnsubscribeCmdHandler(net.Conn)
	PublishCmdHandler(conn net.Conn, message string)
}

type CmdHandlerService struct {
	redisService  core.IRedis
	brokerService core.IBroker
}

func NewCmdHandlerService() *CmdHandlerService {
	if serverCmdHandlerServiceInstance == nil {
		serverCmdHandlerServiceMutex.Lock()
		defer serverCmdHandlerServiceMutex.Unlock()
		if serverCmdHandlerServiceInstance == nil {
			instance := &CmdHandlerService{
				redisService:  core.NewRedisService(),
				brokerService: core.NewBrokerService(),
			}
			serverCmdHandlerServiceInstance = instance
		}
	}
	return serverCmdHandlerServiceInstance
}

func (s *CmdHandlerService) ExitCmdHandler(addr string) {
	fmt.Println("bye", addr)
}

func (s *CmdHandlerService) PingCmdHandler(conn net.Conn) error {
	_, err := conn.Write([]byte("PONG\n"))
	return err
}

func (s *CmdHandlerService) SetCmdHandler(conn net.Conn, cacheFolder string, message string) error {
	k, v := parser.ExtractSetCmd(message)

	ba, _ := utils.ToByteArray(v)
	appendedBa := appendByteTypeToFront(ba, model.StringByteType)
	err := s.redisService.WriteCache(cacheFolder, k, appendedBa)
	if err != nil {
		_, _ = conn.Write([]byte("Set failed" + "\n"))
		return err
	} else {
		_, _ = conn.Write([]byte("Set ok" + "\n"))
		s.redisService.SetString(k, v)
		return nil
	}
}

func (s *CmdHandlerService) GetCmdHandler(conn net.Conn, message string) error {
	k := parser.ExtractGetCmd(message)
	v := s.redisService.Get(k)

	_, err := conn.Write(append(v, []byte("\n")...))
	return err
}

func (s *CmdHandlerService) SubscribeCmdHandler(conn net.Conn, message string) error {
	topic := parser.ExtractSubscribeCmd(message)
	s.brokerService.Subscribe(conn, topic)

	joinedMsg := fmt.Sprintf("%s has joined us.", conn.RemoteAddr())
	s.brokerService.Publish(conn, topic, joinedMsg)

	_, err := conn.Write([]byte("Subscribed\n"))
	return err
}

func (s *CmdHandlerService) OtherCmdHandler(conn net.Conn, message string) error {
	_, err := conn.Write([]byte(message))
	return err
}

func appendByteTypeToFront(originalByteArray []byte, byteType model.ByteType) []byte {
	bt := byte(byteType)
	newByteArray := append([]byte{bt}, originalByteArray...)
	return newByteArray
}

func (s *CmdHandlerService) UnsubscribeCmdHandler(conn net.Conn) {
	topic, exists := s.brokerService.GetTopicFromConnection(conn)
	if exists {
		s.brokerService.Unsubscribe(conn)

		leftMsg := fmt.Sprintf("%s has left.", conn.RemoteAddr())
		s.brokerService.Publish(conn, topic, leftMsg)
	}
}

func (s *CmdHandlerService) PublishCmdHandler(conn net.Conn, message string) {
	topic, exists := s.brokerService.GetTopicFromConnection(conn)
	if exists {
		s.brokerService.Publish(conn, topic, message)
	}
}
