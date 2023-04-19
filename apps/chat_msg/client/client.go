package chat_msg_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_chat_msg"
)

type ChatMessageClient interface {
	GetChatMessageList(req *pb_chat_msg.GetChatMessageListReq) (resp *pb_chat_msg.GetChatMessageListResp)
	// 弃用
	// GetChatMessages(req *pb_chat_msg.GetChatMessagesReq) (resp *pb_chat_msg.GetChatMessagesResp)
	SearchMessage(req *pb_chat_msg.SearchMessageReq) (resp *pb_chat_msg.SearchMessageResp)
	MessageOperation(req *pb_chat_msg.MessageOperationReq) (resp *pb_chat_msg.MessageOperationResp)
}

type chatMessageClient struct {
	opt *xgrpc.ClientDialOption
}

func NewChatMessageClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) ChatMessageClient {
	return &chatMessageClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *chatMessageClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *chatMessageClient) GetChatMessageList(req *pb_chat_msg.GetChatMessageListReq) (resp *pb_chat_msg.GetChatMessageListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_msg.NewChatMessageClient(conn)
	var err error
	resp, err = client.GetChatMessageList(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

//func (c *chatMessageClient) GetChatMessages(req *pb_chat_msg.GetChatMessagesReq) (resp *pb_chat_msg.GetChatMessagesResp) {
//	conn := c.GetClientConn()
//	if conn == nil {
//		return
//	}
//	client := pb_chat_msg.NewChatMessageClient(conn)
//	var err error
//	resp, err = client.GetChatMessages(context.Background(), req)
//	if err != nil {
//		xlog.Warn(err.Error())
//	}
//	return
//}

func (c *chatMessageClient) SearchMessage(req *pb_chat_msg.SearchMessageReq) (resp *pb_chat_msg.SearchMessageResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_msg.NewChatMessageClient(conn)
	var err error
	resp, err = client.SearchMessage(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatMessageClient) MessageOperation(req *pb_chat_msg.MessageOperationReq) (resp *pb_chat_msg.MessageOperationResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_msg.NewChatMessageClient(conn)
	var err error
	resp, err = client.MessageOperation(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
