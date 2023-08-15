package main

import (
	"mini-redis-go/internal/config"
)

func GetServerConfig() *config.ServerConfig {
	return &config.ServerConfig{
		ServerPublicKeyFile:  "/Users/chantapat.t/GolandProjects/mini-redis-go/internal/config/ssl/server/server.pem",
		ServerPrivateKeyFile: "/Users/chantapat.t/GolandProjects/mini-redis-go/internal/config/ssl/server/server.key",
		ConnectionHost:       "localhost",
		ConnectionPort:       "6973",
		CacheFolder:          "/Users/chantapat.t/GolandProjects/mini-redis-go/cache",
	}
}

func main() {
	myApp := InitializeServer()
	myApp.StartServer()
}
