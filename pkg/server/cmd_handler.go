package server

import (
	"fmt"
	"mini-redis-go/pkg/core"
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
	k, v := extractSetCli(message)
	myRedis := core.GetMyRedis()

	err := cacheAsFile(server.CacheFolder, k, v)
	if err != nil {
		_, _ = (*conn).Write([]byte("Set failed" + "\n"))
		return err
	} else {
		_, _ = (*conn).Write([]byte("Set ok" + "\n"))
		myRedis.Set(k, v)
		return nil
	}
}

func getCmdHandler(conn *net.Conn, message string) error {
	k := extractGetCli(message)
	myRedis := core.GetMyRedis()
	v := myRedis.Get(k)

	_, err := (*conn).Write(append(v, []byte("\n")...))
	return err
}

func otherCmdHandler(conn *net.Conn, message string) error {
	_, err := (*conn).Write([]byte(message))
	return err
}
