//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"mini-redis-go/internal/app"
	"mini-redis-go/internal/service/server"
)

func InitializeServer() *app.ServerApp {
	wire.Build(GetServerConfig, server.NewServerService, app.NewServerApp)
	return &app.ServerApp{}
}
