package chat_gpt

import (
	"lark/apps/chat_gpt/internal/config"
	"lark/apps/chat_gpt/internal/service"
	"lark/pkg/common/xgrpc"
)

type ChatGptServer interface {
	Run()
}

type chatGptServer struct {
	cfg            *config.Config
	grpcServer     *xgrpc.GrpcServer
	chatGptService service.ChatGptService
}

func NewChatGptServer(cfg *config.Config, grpcServer *xgrpc.GrpcServer, chatGptService service.ChatGptService) ChatGptServer {
	return &chatGptServer{
		cfg:            cfg,
		grpcServer:     grpcServer,
		chatGptService: chatGptService,
	}
}

func (s *chatGptServer) Run() {

}
