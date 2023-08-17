// Package handler contains business logic of commands sent from a client
package handler

import (
	"fmt"
	"mini-redis-go/internal/model"
	"mini-redis-go/internal/service/server/core"
	"mini-redis-go/internal/service/server/parser"
	"mini-redis-go/internal/utils"
	"net"
)

type CmdHandlerService interface {
	ExitCmdHandler(addr string)
	PingCmdHandler(net.Conn) error
	SetCmdHandler(conn net.Conn, cacheFolder string, message string) error
	GetCmdHandler(conn net.Conn, message string) error
	SubscribeCmdHandler(conn net.Conn, message string) error
	OtherCmdHandler(conn net.Conn, message string) error

	UnsubscribeCmdHandler(net.Conn)
	PublishCmdHandler(conn net.Conn, message string)
}

type cmdHandlerService struct {
	redisService  core.RedisService
	brokerService core.BrokerService
}

func NewCmdHandlerService(redisService core.RedisService, brokerService core.BrokerService) *cmdHandlerService {
	return &cmdHandlerService{
		redisService:  redisService,
		brokerService: brokerService,
	}
}

func (*cmdHandlerService) ExitCmdHandler(addr string) {
	fmt.Println("bye", addr)
}

func (*cmdHandlerService) PingCmdHandler(conn net.Conn) error {
	_, err := conn.Write([]byte("PONG\n"))
	return err
}

func (s *cmdHandlerService) SetCmdHandler(conn net.Conn, cacheFolder string, message string) error {
	k, v := parser.ExtractSetCmd(message)

	ba, _ := utils.ToByteArray(v)
	appendedBa := appendByteTypeToFront(ba, model.StringByteType)
	err := s.redisService.WriteCache(cacheFolder, k, appendedBa)
	if err != nil {
		_, _ = conn.Write([]byte("Set failed" + "\n"))
		return err
	} else {
		_, _ = conn.Write([]byte("Set ok" + "\n"))
		s.redisService.Set(k, ba)
		return nil
	}
}

func (s *cmdHandlerService) GetCmdHandler(conn net.Conn, message string) error {
	k := parser.ExtractGetCmd(message)
	v := s.redisService.Get(k)

	_, err := conn.Write(append(v, []byte("\n")...))
	return err
}

func (s *cmdHandlerService) SubscribeCmdHandler(conn net.Conn, message string) error {
	topic := parser.ExtractSubscribeCmd(message)
	s.brokerService.Subscribe(conn, topic)

	_, err := conn.Write([]byte("Subscribed\n"))
	return err
}

func (*cmdHandlerService) OtherCmdHandler(conn net.Conn, message string) error {
	_, err := conn.Write([]byte(message))
	return err
}

func appendByteTypeToFront(originalByteArray []byte, byteType model.ByteType) []byte {
	bt := byte(byteType)
	newByteArray := append([]byte{bt}, originalByteArray...)
	return newByteArray
}

func (s *cmdHandlerService) UnsubscribeCmdHandler(conn net.Conn) {
	_, exists := s.brokerService.GetTopicFromConnection(conn)
	if exists {
		s.brokerService.Unsubscribe(conn)
	}
}

func (s *cmdHandlerService) PublishCmdHandler(conn net.Conn, message string) {
	topic, exists := s.brokerService.GetTopicFromConnection(conn)
	if exists {
		s.brokerService.Publish(conn, topic, message)
	}
}
