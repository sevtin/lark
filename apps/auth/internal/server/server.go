package server

import (
	"lark/apps/auth/internal/server/auth"
	"lark/pkg/commands"
)

type server struct {
	authServer auth.AuthServer
}

func NewServer(authServer auth.AuthServer) commands.MainInstance {
	return &server{authServer: authServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.authServer.Run()
}

func (s *server) Destroy() {

}
