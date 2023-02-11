package msg_hot

import "lark/apps/msg_hot/internal/service"

type MessageHotServer interface {
	Run()
}

type messageHotServer struct {
	messageHotService service.MessageHotService
}

func NewMessageHotServer(messageHotService service.MessageHotService) MessageHotServer {
	return &messageHotServer{messageHotService: messageHotService}
}

func (s *messageHotServer) Run() {

}
