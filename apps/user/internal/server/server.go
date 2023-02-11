package server

import (
	"lark/apps/user/internal/server/user"
	"lark/pkg/commands"
)

type server struct {
	userServer user.UserServer
}

func NewServer(userServer user.UserServer) commands.MainInstance {
	return &server{userServer: userServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.userServer.Run()
}

func (s *server) Destroy() {

}
