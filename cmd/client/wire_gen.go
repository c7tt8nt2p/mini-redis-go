// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"mini-redis-go/internal/app"
	"mini-redis-go/internal/service/client"
)

// Injectors from wire.go:

func InitializeClient() *app.ClientApp {
	clientConfig := GetClientConfig()
	clientService := client.NewClientService(clientConfig)
	clientApp := app.NewClientApp(clientService)
	return clientApp
}
