package chat_invite_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_invite"
)

type ChatInviteClient interface {
	InitiateChatInvite(req *pb_invite.InitiateChatInviteReq) (resp *pb_invite.InitiateChatInviteResp)
	ChatInviteHandle(req *pb_invite.ChatInviteHandleReq) (resp *pb_invite.ChatInviteHandleResp)
	ChatInviteList(req *pb_invite.ChatInviteListReq) (resp *pb_invite.ChatInviteListResp)
}

type chatInviteClient struct {
	opt *xgrpc.ClientDialOption
}

func NewChatInviteClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) ChatInviteClient {
	return &chatInviteClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *chatInviteClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *chatInviteClient) InitiateChatInvite(req *pb_invite.InitiateChatInviteReq) (resp *pb_invite.InitiateChatInviteResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_invite.NewInviteClient(conn)
	var err error
	resp, err = client.InitiateChatInvite(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatInviteClient) ChatInviteHandle(req *pb_invite.ChatInviteHandleReq) (resp *pb_invite.ChatInviteHandleResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_invite.NewInviteClient(conn)
	var err error
	resp, err = client.ChatInviteHandle(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatInviteClient) ChatInviteList(req *pb_invite.ChatInviteListReq) (resp *pb_invite.ChatInviteListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_invite.NewInviteClient(conn)
	var err error
	resp, err = client.ChatInviteList(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
