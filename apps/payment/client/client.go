package payment_client

import (
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/conf"
)

type PaymentClient interface {
}

type paymentClient struct {
	opt *xgrpc.ClientDialOption
}

func NewPaymentClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) PaymentClient {
	return &paymentClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *paymentClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}
