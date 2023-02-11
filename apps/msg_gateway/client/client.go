package gw_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_gw"
)

type MessageGatewayClient interface {
	HealthCheck(req *pb_gw.HealthCheckReq) (resp *pb_gw.HealthCheckResp)
	SendTopicMessage(req *pb_gw.SendTopicMessageReq) (resp *pb_gw.SendTopicMessageResp)
}

type messageGatewayClient struct {
	opt *xgrpc.ClientDialOption
}

func NewMsgGwClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) MessageGatewayClient {
	return &messageGatewayClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *messageGatewayClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *messageGatewayClient) HealthCheck(req *pb_gw.HealthCheckReq) (resp *pb_gw.HealthCheckResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_gw.NewMessageGatewayClient(conn)
	var err error
	resp, err = client.HealthCheck(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *messageGatewayClient) SendTopicMessage(req *pb_gw.SendTopicMessageReq) (resp *pb_gw.SendTopicMessageResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_gw.NewMessageGatewayClient(conn)
	var err error
	resp, err = client.SendTopicMessage(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
