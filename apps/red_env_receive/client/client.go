package red_env_receive_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_red_env_receive"
)

// 暂不开放
type RedEnvReceiveClient interface {
	GrabRedEnvelope(req *pb_red_env_receive.GrabRedEnvelopeReq) (resp *pb_red_env_receive.GrabRedEnvelopeResp)
	OpenRedEnvelope(req *pb_red_env_receive.OpenRedEnvelopeReq) (resp *pb_red_env_receive.OpenRedEnvelopeResp)
}

type redEnvReceiveClient struct {
	opt *xgrpc.ClientDialOption
}

func NewRedEnvReceiveClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) RedEnvReceiveClient {
	return &redEnvReceiveClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *redEnvReceiveClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *redEnvReceiveClient) GrabRedEnvelope(req *pb_red_env_receive.GrabRedEnvelopeReq) (resp *pb_red_env_receive.GrabRedEnvelopeResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_red_env_receive.NewRedEnvReceiveClient(conn)
	var err error
	resp, err = client.GrabRedEnvelope(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *redEnvReceiveClient) OpenRedEnvelope(req *pb_red_env_receive.OpenRedEnvelopeReq) (resp *pb_red_env_receive.OpenRedEnvelopeResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_red_env_receive.NewRedEnvReceiveClient(conn)
	var err error
	resp, err = client.OpenRedEnvelope(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
