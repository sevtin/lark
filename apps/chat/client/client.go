package chat_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_chat"
)

type ChatClient interface {
	CreateGroupChat(req *pb_chat.CreateGroupChatReq) (resp *pb_chat.CreateGroupChatResp)
	EditGroupChat(req *pb_chat.EditGroupChatReq) (resp *pb_chat.EditGroupChatResp)
	GroupChatDetails(req *pb_chat.GroupChatDetailsReq) (resp *pb_chat.GroupChatDetailsResp)
	RemoveGroupChatMember(req *pb_chat.RemoveGroupChatMemberReq) (resp *pb_chat.RemoveGroupChatMemberResp)
	QuitGroupChat(req *pb_chat.QuitGroupChatReq) (resp *pb_chat.QuitGroupChatResp)
	DeleteContact(req *pb_chat.DeleteContactReq) (resp *pb_chat.DeleteContactResp)
	UploadAvatar(req *pb_chat.UploadAvatarReq) (resp *pb_chat.UploadAvatarResp)
	GetChatInfo(req *pb_chat.GetChatInfoReq) (resp *pb_chat.GetChatInfoResp)
}

type chatClient struct {
	opt *xgrpc.ClientDialOption
}

func NewChatClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) ChatClient {
	return &chatClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *chatClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *chatClient) CreateGroupChat(req *pb_chat.CreateGroupChatReq) (resp *pb_chat.CreateGroupChatResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat.NewChatClient(conn)
	var err error
	resp, err = client.CreateGroupChat(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatClient) EditGroupChat(req *pb_chat.EditGroupChatReq) (resp *pb_chat.EditGroupChatResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat.NewChatClient(conn)
	var err error
	resp, err = client.EditGroupChat(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatClient) GroupChatDetails(req *pb_chat.GroupChatDetailsReq) (resp *pb_chat.GroupChatDetailsResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat.NewChatClient(conn)
	var err error
	resp, err = client.GroupChatDetails(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatClient) RemoveGroupChatMember(req *pb_chat.RemoveGroupChatMemberReq) (resp *pb_chat.RemoveGroupChatMemberResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat.NewChatClient(conn)
	var err error
	resp, err = client.RemoveGroupChatMember(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatClient) QuitGroupChat(req *pb_chat.QuitGroupChatReq) (resp *pb_chat.QuitGroupChatResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat.NewChatClient(conn)
	var err error
	resp, err = client.QuitGroupChat(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatClient) DeleteContact(req *pb_chat.DeleteContactReq) (resp *pb_chat.DeleteContactResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat.NewChatClient(conn)
	var err error
	resp, err = client.DeleteContact(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatClient) UploadAvatar(req *pb_chat.UploadAvatarReq) (resp *pb_chat.UploadAvatarResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat.NewChatClient(conn)
	var err error
	resp, err = client.UploadAvatar(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *chatClient) GetChatInfo(req *pb_chat.GetChatInfoReq) (resp *pb_chat.GetChatInfoResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_chat.NewChatClient(conn)
	var err error
	resp, err = client.GetChatInfo(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
