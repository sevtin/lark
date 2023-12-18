package template

var AppsClientTemplate = ParseTemplate(`
package {{.PackageName}}_client

import (
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/conf"
)

type {{.UpperPackageName}}Client interface {

}

type {{.LowerPackageName}}Client struct {
	opt *xgrpc.ClientDialOption
}

func New{{.UpperPackageName}}Client(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) {{.UpperPackageName}}Client {
	return &{{.LowerPackageName}}Client{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *{{.LowerPackageName}}Client) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}
`)
