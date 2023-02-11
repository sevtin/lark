package server

import (
	"lark/apps/avatar/internal/server/avatar"
	"lark/pkg/commands"
)

type server struct {
	avatarServer avatar.AvatarServer
}

func NewServer(avatarServer avatar.AvatarServer) commands.MainInstance {
	return &server{avatarServer: avatarServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.avatarServer.Run()
}

func (s *server) Destroy() {

}
