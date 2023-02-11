package avatar

import (
	"google.golang.org/grpc"
	"io"
	"lark/apps/avatar/internal/config"
	"lark/apps/avatar/internal/service"
	"lark/pkg/common/xgrpc"
	"lark/pkg/proto/pb_avatar"
)

type AvatarServer interface {
	Run()
}

type avatarServer struct {
	pb_avatar.UnimplementedAvatarServer
	cfg           *config.Config
	grpcServer    *xgrpc.GrpcServer
	avatarService service.AvatarService
}

func NewAvatarServer(cfg *config.Config, avatarService service.AvatarService) AvatarServer {
	return &avatarServer{cfg: cfg, avatarService: avatarService}
}

func (s *avatarServer) Run() {
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

	pb_avatar.RegisterAvatarServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
