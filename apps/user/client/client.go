package user_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_user"
)

type UserClient interface {
	EditUserInfo(req *pb_user.EditUserInfoReq) (resp *pb_user.EditUserInfoResp)
	GetUserInfo(req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp)
	GetBasicUserInfo(req *pb_user.GetBasicUserInfoReq) (resp *pb_user.GetBasicUserInfoResp)
	GetUserList(req *pb_user.GetUserListReq) (resp *pb_user.GetUserListResp)
	SearchUser(req *pb_user.SearchUserReq) (resp *pb_user.SearchUserResp)
	UploadAvatar(req *pb_user.UploadAvatarReq) (resp *pb_user.UploadAvatarResp)
	GetBasicUserInfoList(req *pb_user.GetBasicUserInfoListReq) (resp *pb_user.GetBasicUserInfoListResp)
	GetServerIdList(req *pb_user.GetServerIdListReq) (resp *pb_user.GetServerIdListResp)
}

type userClient struct {
	opt *xgrpc.ClientDialOption
}

func NewUserClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) UserClient {
	return &userClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *userClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *userClient) EditUserInfo(req *pb_user.EditUserInfoReq) (resp *pb_user.EditUserInfoResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_user.NewUserClient(conn)
	var err error
	resp, err = client.EditUserInfo(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *userClient) GetUserInfo(req *pb_user.UserInfoReq) (resp *pb_user.UserInfoResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_user.NewUserClient(conn)
	var err error
	resp, err = client.GetUserInfo(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *userClient) GetBasicUserInfo(req *pb_user.GetBasicUserInfoReq) (resp *pb_user.GetBasicUserInfoResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_user.NewUserClient(conn)
	var err error
	resp, err = client.GetBasicUserInfo(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *userClient) GetUserList(req *pb_user.GetUserListReq) (resp *pb_user.GetUserListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_user.NewUserClient(conn)
	var err error
	resp, err = client.GetUserList(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *userClient) SearchUser(req *pb_user.SearchUserReq) (resp *pb_user.SearchUserResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_user.NewUserClient(conn)
	var err error
	resp, err = client.SearchUser(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *userClient) UploadAvatar(req *pb_user.UploadAvatarReq) (resp *pb_user.UploadAvatarResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_user.NewUserClient(conn)
	var err error
	resp, err = client.UploadAvatar(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *userClient) GetBasicUserInfoList(req *pb_user.GetBasicUserInfoListReq) (resp *pb_user.GetBasicUserInfoListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_user.NewUserClient(conn)
	var err error
	resp, err = client.GetBasicUserInfoList(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *userClient) GetServerIdList(req *pb_user.GetServerIdListReq) (resp *pb_user.GetServerIdListResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_user.NewUserClient(conn)
	var err error
	resp, err = client.GetServerIdList(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
