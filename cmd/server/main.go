package main

import (
	"mini-redis-go/internal/config"
	"mini-redis-go/internal/service/server"
)

func main() {
	s := server.NewServer(config.ConnectionHost, config.ConnectionPort, config.CacheFolder)
	s.Start()
}
