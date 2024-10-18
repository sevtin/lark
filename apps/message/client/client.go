package msg_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_msg"
)

type MsgClient interface {
	SendChatMessage(req *pb_msg.SendChatMessageReq) (resp *pb_msg.SendChatMessageResp)
	// 弃用
	//MessageOperation(req *pb_msg.MessageOperationReq) (resp *pb_msg.MessageOperationResp)
}

type msgClient struct {
	opt *xgrpc.ClientDialOption
}

func NewMsgClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) MsgClient {
	return &msgClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *msgClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *msgClient) SendChatMessage(req *pb_msg.SendChatMessageReq) (resp *pb_msg.SendChatMessageResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_msg.NewMessageClient(conn)
	var err error
	resp, err = client.SendChatMessage(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *msgClient) MessageOperation(req *pb_msg.MessageOperationReq) (resp *pb_msg.MessageOperationResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_msg.NewMessageClient(conn)
	var err error
	resp, err = client.MessageOperation(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
