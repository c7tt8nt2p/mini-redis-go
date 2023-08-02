package main

import (
	"mini-redis-go/pkg/config"
	"mini-redis-go/pkg/server"
)

func main() {
	s := server.NewServer(config.ConnectionHost, config.ConnectionPort)
	s.Start()
}
