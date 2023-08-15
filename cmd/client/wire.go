//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"mini-redis-go/internal/app"
	"mini-redis-go/internal/service/client"
)

func InitializeClient() *app.ClientApp {
	wire.Build(GetClientConfig, client.NewClientService, app.NewClientApp)
	return &app.ClientApp{}
}
