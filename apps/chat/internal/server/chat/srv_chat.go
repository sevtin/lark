package chat

import (
	"google.golang.org/grpc"
	"io"
	"lark/apps/chat/internal/config"
	"lark/apps/chat/internal/service"
	"lark/pkg/common/xgrpc"
	"lark/pkg/proto/pb_chat"
)

type ChatServer interface {
	Run()
}

type chatServer struct {
	pb_chat.UnimplementedChatServer
	cfg         *config.Config
	grpcServer  *xgrpc.GrpcServer
	chatService service.ChatService
}

func NewChatServer(cfg *config.Config, chatService service.ChatService) ChatServer {
	return &chatServer{cfg: cfg, chatService: chatService}
}

func (s *chatServer) Run() {
	var (
		srv    *grpc.Server
		closer io.Closer
	)
	srv, closer = xgrpc.NewServer(s.cfg.GrpcServer)
	defer func() {
		if closer != nil {
			closer.Close()
		}
	}()

	pb_chat.RegisterChatServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
