package server

import (
	"lark/apps/msg_history/internal/server/msg_history"
	"lark/pkg/commands"
)

type server struct {
	messageStoreServer msg_history.MessageHistoryServer
}

func NewServer(messageStoreServer msg_history.MessageHistoryServer) commands.MainInstance {
	return &server{messageStoreServer: messageStoreServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.messageStoreServer.Run()
}

func (s *server) Destroy() {

}
