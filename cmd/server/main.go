package main

import (
	"mini-redis-go/internal/app"
	"mini-redis-go/internal/config"
	"mini-redis-go/internal/service/server"
)

func main() {
	serverService := server.NewServerService(config.ConnectionHost, config.ConnectionPort, config.CacheFolder)

	myApp := app.NewServerApp(serverService)
	myApp.StartServer()
}
