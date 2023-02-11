package server

import (
	"lark/apps/chat_member/internal/server/chat_member"
	"lark/pkg/commands"
)

type server struct {
	chatMemberServer chat_member.ChatMemberServer
}

func NewServer(chatMemberServer chat_member.ChatMemberServer) commands.MainInstance {
	return &server{chatMemberServer: chatMemberServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.chatMemberServer.Run()
}

func (s *server) Destroy() {

}
