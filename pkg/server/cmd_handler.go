package server

import (
	"fmt"
	"mini-redis-go/pkg/core/broker"
	"mini-redis-go/pkg/core/redis"
	"mini-redis-go/pkg/server/conversion"
	"net"
)

func exitCmdHandler(addr string) {
	fmt.Println("bye", addr)
}

func pingCmdHandler(conn *net.Conn) error {
	_, err := (*conn).Write([]byte("PONG\n"))
	return err
}

func setCmdHandler(conn *net.Conn, server *Server, message string) error {
	k, v := extractSetCmd(message)
	myRedis := redis.GetMyRedis()

	ba, _ := conversion.ToByteArray(v)
	appendedBa := appendByteTypeToFront(ba, redis.StringByteType)
	err := cacheAsFile(server.CacheFolder, k, appendedBa)
	if err != nil {
		_, _ = (*conn).Write([]byte("Set failed" + "\n"))
		return err
	} else {
		_, _ = (*conn).Write([]byte("Set ok" + "\n"))
		myRedis.SetString(k, v)
		return nil
	}
}

func getCmdHandler(conn *net.Conn, message string) error {
	k := extractGetCmd(message)
	myRedis := redis.GetMyRedis()
	v := myRedis.Get(k)

	_, err := (*conn).Write(append(v, []byte("\n")...))
	return err
}

func subscribeCmdHandler(conn *net.Conn, message string) error {
	topic := extractSubscribeCmd(message)
	b := broker.GetMyBroker()
	b.Subscribe(conn, topic)

	joinedMsg := fmt.Sprintf("%s has joined us", (*conn).RemoteAddr())
	b.Publish(conn, topic, joinedMsg)

	_, err := (*conn).Write([]byte("Subscribed\n"))
	return err
}

func otherCmdHandler(conn *net.Conn, message string) error {
	_, err := (*conn).Write([]byte(message))
	return err
}

func appendByteTypeToFront(originalByteArray []byte, byteType redis.ByteType) []byte {
	bt := byte(byteType)
	newByteArray := append([]byte{bt}, originalByteArray...)
	return newByteArray
}

func unsubscribeCmdHandler(conn *net.Conn) {
	b := broker.GetMyBroker()

	topic, exists := b.GetTopicFromConnection(conn)
	if exists {
		b.Unsubscribe(conn)

		leftMsg := fmt.Sprintf("%s has left", (*conn).RemoteAddr())
		b.Publish(conn, topic, leftMsg)
	}
}

func publishCmdHandler(conn *net.Conn, message string) {
	b := broker.GetMyBroker()
	topic, exists := b.GetTopicFromConnection(conn)
	if exists {
		b.Publish(conn, topic, message)
	}
}
