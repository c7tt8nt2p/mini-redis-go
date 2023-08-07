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

	if myRedis.ExistsByKey(k) {
		myRedis.Set(k, v)
		cacheRewriteAll(&myRedis, server.CacheFolder)
	} else {
		myRedis.Set(k, v)
		cacheAppend(server.CacheFolder, k, v)
	}
	_, err := (*conn).Write([]byte("Set ok" + "\n"))
	return err
}

func getCmdHandler(conn *net.Conn, message string) error {
	k := extractGetCli(message)
	myRedis := core.GetMyRedis()
	v := myRedis.Get(k)

	_, err := (*conn).Write([]byte(v + "\n"))
	return err
}

func otherCmdHandler(conn *net.Conn, message string) error {
	_, err := (*conn).Write([]byte(message))
	return err
}
