package server

import (
	"lark/apps/msg_hot/internal/server/msg_hot"
	"lark/pkg/commands"
)

type server struct {
	messageHotServer msg_hot.MessageHotServer
}

func NewServer(messageHotServer msg_hot.MessageHotServer) commands.MainInstance {
	return &server{messageHotServer: messageHotServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.messageHotServer.Run()
}

func (s *server) Destroy() {

}
