package template

var AppsGrpcServerTemplate = ParseTemplate(`
package {{.PackageName}}

import (
	"google.golang.org/grpc"
	"io"
	"lark/apps/{{.PackageName}}/internal/config"
	"lark/apps/{{.PackageName}}/internal/service"
	"lark/pkg/common/xgrpc"
	"lark/pkg/proto/pb_{{.PackageName}}"
)

type {{.UpperServiceName}}Server interface {
	Run()
}

type {{.LowerServiceName}}Server struct {
	pb_{{.PackageName}}.Unimplemented{{.UpperServiceName}}Server
	cfg         *config.Config
	{{.LowerServiceName}}Service service.{{.UpperServiceName}}Service
	grpcServer  *xgrpc.GrpcServer
}

func New{{.UpperServiceName}}Server(cfg *config.Config, {{.LowerServiceName}}Service service.{{.UpperServiceName}}Service) {{.UpperServiceName}}Server {
	return &{{.LowerServiceName}}Server{cfg: cfg, {{.LowerServiceName}}Service: {{.LowerServiceName}}Service}
}

func (s *{{.LowerServiceName}}Server) Run() {
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

	pb_{{.PackageName}}.Register{{.UpperServiceName}}Server(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
`)
