package server

import (
	"lark/apps/dist/internal/server/dist"
	"lark/pkg/commands"
)

type server struct {
	distServer dist.DistServer
}

func NewServer(distServer dist.DistServer) commands.MainInstance {
	return &server{distServer: distServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.distServer.Run()
}

func (s *server) Destroy() {

}
