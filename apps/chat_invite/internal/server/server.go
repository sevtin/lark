package server

import (
	"lark/apps/chat_invite/internal/server/chat_invite"
	"lark/pkg/commands"
)

type server struct {
	chatInviteServer chat_invite.ChatInviteServer
}

func NewServer(chatInviteServer chat_invite.ChatInviteServer) commands.MainInstance {
	return &server{chatInviteServer: chatInviteServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.chatInviteServer.Run()
}

func (s *server) Destroy() {

}
