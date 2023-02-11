package dist

import (
	"google.golang.org/grpc"
	"io"
	"lark/apps/dist/internal/config"
	"lark/apps/dist/internal/service"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xmonitor"
	"lark/pkg/proto/pb_dist"
)

type DistServer interface {
	Run()
}

type distServer struct {
	pb_dist.UnimplementedDistServer
	cfg         *config.Config
	distService service.DistService
	grpcServer  *xgrpc.GrpcServer
}

func NewDistServer(cfg *config.Config, distService service.DistService) DistServer {
	srv := &distServer{cfg: cfg, distService: distService}
	return srv
}

func (s *distServer) Run() {
	var (
		srv    *grpc.Server
		closer io.Closer
	)
	xmonitor.RunMonitor(s.cfg.Monitor.Port)

	srv, closer = xgrpc.NewServer(s.cfg.GrpcServer)
	defer func() {
		if closer != nil {
			closer.Close()
		}
	}()

	pb_dist.RegisterDistServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
