package msg_history

import (
	"lark/apps/msg_history/internal/service"
)

type MessageHistoryServer interface {
	Run()
}

type messageHistoryServer struct {
	messageHistoryService service.MessageHistoryService
}

func NewMessageHistoryServer(messageHistoryService service.MessageHistoryService) MessageHistoryServer {
	return &messageHistoryServer{messageHistoryService: messageHistoryService}
}

func (s *messageHistoryServer) Run() {

}
