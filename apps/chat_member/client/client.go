package chat_member_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_chat_member"
)

type ChatMemberClient interface {
	GetChatMemberInfo(req *pb_chat_member.GetChatMemberInfoReq) (resp *pb_chat_member.GetChatMemberInfoResp)
	ChatMemberOnOffLine(req *pb_chat_member.ChatMemberOnOffLineReq) (resp *pb_chat_member.ChatMemberOnOffLineResp)
	GetDistMemberList(req *pb_chat_member.GetDistMemberListReq) (resp *pb_chat_member.GetDistMemberListResp)
	GetChatMemberList(req *pb_chat_member.GetChatMemberListReq) (resp *pb_chat_member.GetChatMemberListResp)
	GetContactList(req *pb_chat_member.GetContactListReq) (resp *pb_chat_member.GetContactListResp)
	GetGroupChatList(req *pb_chat_member.GetGroupChatListReq) (resp *pb_chat_member.GetGroupChatListResp)
}

type chatMemberClient struct {
	opt *xgrpc.ClientDialOption
}

func NewChatMemberClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) ChatMemberClient {
	return &chatMemberClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *chatMemberClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *chatMemberClient) GetChatMemberInfo(req *pb_chat_member.GetChatMemberInfoReq) (resp *pb_chat_member.GetChatMemberInfoResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_member.NewChatMemberClient(conn)
	var err error
	resp, err = client.GetChatMemberInfo(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatMemberClient) ChatMemberOnOffLine(req *pb_chat_member.ChatMemberOnOffLineReq) (resp *pb_chat_member.ChatMemberOnOffLineResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_member.NewChatMemberClient(conn)
	var err error
	resp, err = client.ChatMemberOnOffLine(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatMemberClient) GetDistMemberList(req *pb_chat_member.GetDistMemberListReq) (resp *pb_chat_member.GetDistMemberListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_member.NewChatMemberClient(conn)
	var err error
	resp, err = client.GetDistMemberList(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatMemberClient) GetChatMemberList(req *pb_chat_member.GetChatMemberListReq) (resp *pb_chat_member.GetChatMemberListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_member.NewChatMemberClient(conn)
	var err error
	resp, err = client.GetChatMemberList(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatMemberClient) GetContactList(req *pb_chat_member.GetContactListReq) (resp *pb_chat_member.GetContactListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_member.NewChatMemberClient(conn)
	var err error
	resp, err = client.GetContactList(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatMemberClient) GetGroupChatList(req *pb_chat_member.GetGroupChatListReq) (resp *pb_chat_member.GetGroupChatListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat_member.NewChatMemberClient(conn)
	var err error
	resp, err = client.GetGroupChatList(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
