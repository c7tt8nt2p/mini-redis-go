package app

import "mini-redis-go/internal/service/server"

// IServerApp is an entrypoint when instantiating a new server
type IServerApp interface {
	StartServer()
}

type ServerApp struct {
	serverService server.IServer
}

func NewServerApp(serverService server.IServer) *ServerApp {
	return &ServerApp{
		serverService: serverService,
	}
}

func (s *ServerApp) StartServer() {
	s.serverService.Start()
}
