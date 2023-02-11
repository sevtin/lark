package avatar_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_avatar"
)

type AvatarClient interface {
	SetAvatar(req *pb_avatar.SetAvatarReq) (resp *pb_avatar.SetAvatarResp)
}

type avatarClient struct {
	opt *xgrpc.ClientDialOption
}

func NewAvatarClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) AvatarClient {
	return &avatarClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *avatarClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *avatarClient) SetAvatar(req *pb_avatar.SetAvatarReq) (resp *pb_avatar.SetAvatarResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_avatar.NewAvatarClient(conn)
	var err error
	resp, err = client.SetAvatar(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
