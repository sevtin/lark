package auth

import (
	"google.golang.org/grpc"
	"io"
	"lark/apps/auth/internal/config"
	"lark/apps/auth/internal/service"
	"lark/pkg/common/xgrpc"
	"lark/pkg/proto/pb_auth"
)

type AuthServer interface {
	Run()
}

type authServer struct {
	pb_auth.UnimplementedAuthServer
	cfg         *config.Config
	authService service.AuthService
	grpcServer  *xgrpc.GrpcServer
}

func NewAuthServer(cfg *config.Config, authService service.AuthService) AuthServer {
	return &authServer{cfg: cfg, authService: authService}
}

func (s *authServer) Run() {
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

	pb_auth.RegisterAuthServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
