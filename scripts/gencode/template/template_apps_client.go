package template

var AppsClientTemplate = ParseTemplate(`
package {{.PackageName}}_client

import (
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/conf"
)

type {{.UpperServiceName}}Client interface {

}

type {{.LowerServiceName}}Client struct {
	opt *xgrpc.ClientDialOption
}

func New{{.UpperServiceName}}Client(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) {{.UpperServiceName}}Client {
	return &{{.LowerServiceName}}Client{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *{{.LowerServiceName}}Client) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}
`)
