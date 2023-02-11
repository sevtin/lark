package convo_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_convo"
)

type ConvoClient interface {
	ConvoList(req *pb_convo.ConvoListReq) (resp *pb_convo.ConvoListResp)
	ConvoChatSeqList(req *pb_convo.ConvoChatSeqListReq) (resp *pb_convo.ConvoChatSeqListResp)
}

type convoClient struct {
	opt *xgrpc.ClientDialOption
}

func NewConvoClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) ConvoClient {
	return &convoClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *convoClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *convoClient) ConvoList(req *pb_convo.ConvoListReq) (resp *pb_convo.ConvoListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_convo.NewConvoClient(conn)
	var err error
	resp, err = client.ConvoList(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *convoClient) ConvoChatSeqList(req *pb_convo.ConvoChatSeqListReq) (resp *pb_convo.ConvoChatSeqListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_convo.NewConvoClient(conn)
	var err error
	resp, err = client.ConvoChatSeqList(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
