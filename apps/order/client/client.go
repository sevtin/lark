package order_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_order"
)

type OrderClient interface {
	CreateRedEnvelopeOrder(req *pb_order.CreateRedEnvelopeOrderReq) (resp *pb_order.CreateRedEnvelopeOrderResp)
}

type orderClient struct {
	opt *xgrpc.ClientDialOption
}

func NewOrderClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) OrderClient {
	return &orderClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *orderClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *orderClient) CreateRedEnvelopeOrder(req *pb_order.CreateRedEnvelopeOrderReq) (resp *pb_order.CreateRedEnvelopeOrderResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_order.NewOrderClient(conn)
	var err error
	resp, err = client.CreateRedEnvelopeOrder(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
