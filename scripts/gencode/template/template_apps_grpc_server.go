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

type {{.UpperPackageName}}Server interface {
	Run()
}

type {{.LowerPackageName}}Server struct {
	pb_{{.PackageName}}.Unimplemented{{.UpperPackageName}}Server
	cfg         *config.Config
	{{.LowerPackageName}}Service service.{{.UpperPackageName}}Service
	grpcServer  *xgrpc.GrpcServer
}

func New{{.UpperPackageName}}Server(cfg *config.Config, {{.LowerPackageName}}Service service.{{.UpperPackageName}}Service) {{.UpperPackageName}}Server {
	return &{{.LowerPackageName}}Server{cfg: cfg, {{.LowerPackageName}}Service: {{.LowerPackageName}}Service}
}

func (s *{{.LowerPackageName}}Server) Run() {
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

	pb_{{.PackageName}}.Register{{.UpperPackageName}}Server(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
`)
