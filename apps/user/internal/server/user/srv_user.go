package user

import (
	"google.golang.org/grpc"
	"io"
	"lark/apps/user/internal/config"
	"lark/apps/user/internal/service"
	"lark/pkg/common/xgrpc"
	"lark/pkg/proto/pb_user"
)

type UserServer interface {
	Run()
}

type userServer struct {
	pb_user.UnimplementedUserServer
	cfg         *config.Config
	grpcServer  *xgrpc.GrpcServer
	userService service.UserService
}

func NewUserServer(cfg *config.Config, userService service.UserService) UserServer {
	return &userServer{cfg: cfg, userService: userService}
}

func (s *userServer) Run() {
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

	pb_user.RegisterUserServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
