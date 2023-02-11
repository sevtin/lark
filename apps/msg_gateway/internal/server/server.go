package server

import (
	"lark/apps/msg_gateway/internal/server/gateway"
	"lark/pkg/commands"
)

type server struct {
	gatewayServer gateway.GatewayServer
}

func NewServer(gatewayServer gateway.GatewayServer) commands.MainInstance {
	return &server{gatewayServer: gatewayServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.gatewayServer.Run()
}

func (s *server) Destroy() {

}
