package auth_client

import (
	"context"
	"google.golang.org/grpc"
	"lark/pkg/common/xgrpc"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
	"lark/pkg/proto/pb_auth"
)

type AuthClient interface {
	SignUp(req *pb_auth.SignUpReq) (resp *pb_auth.SignUpResp)
	SignIn(req *pb_auth.SignInReq) (resp *pb_auth.SignInResp)
	RefreshToken(req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp)
	SignOut(req *pb_auth.SignOutReq) (resp *pb_auth.SignOutResp)
	GithubOAuth2Callback(req *pb_auth.GithubOAuth2CallbackReq) (resp *pb_auth.GithubOAuth2CallbackResp)
	GoogleOAuth2Callback(req *pb_auth.GoogleOAuth2CallbackReq) (resp *pb_auth.GoogleOAuth2CallbackResp)
}

type authClient struct {
	opt *xgrpc.ClientDialOption
}

func NewAuthClient(etcd *conf.Etcd, server *conf.GrpcServer, jaeger *conf.Jaeger, clientName string) AuthClient {
	return &authClient{xgrpc.NewClientDialOption(etcd, server, jaeger, clientName)}
}

func (c *authClient) GetClientConn() (conn *grpc.ClientConn) {
	conn = xgrpc.GetClientConn(c.opt.DialOption)
	return
}

func (c *authClient) SignUp(req *pb_auth.SignUpReq) (resp *pb_auth.SignUpResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_auth.NewAuthClient(conn)
	var err error
	resp, err = client.SignUp(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *authClient) SignIn(req *pb_auth.SignInReq) (resp *pb_auth.SignInResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_auth.NewAuthClient(conn)
	var err error
	resp, err = client.SignIn(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *authClient) RefreshToken(req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_auth.NewAuthClient(conn)
	var err error
	resp, err = client.RefreshToken(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *authClient) SignOut(req *pb_auth.SignOutReq) (resp *pb_auth.SignOutResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_auth.NewAuthClient(conn)
	var err error
	resp, err = client.SignOut(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *authClient) GithubOAuth2Callback(req *pb_auth.GithubOAuth2CallbackReq) (resp *pb_auth.GithubOAuth2CallbackResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_auth.NewAuthClient(conn)
	var err error
	resp, err = client.GithubOAuth2Callback(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}

func (c *authClient) GoogleOAuth2Callback(req *pb_auth.GoogleOAuth2CallbackReq) (resp *pb_auth.GoogleOAuth2CallbackResp) {
	conn := c.GetClientConn()
	if conn == nil {
		return
	}
	client := pb_auth.NewAuthClient(conn)
	var err error
	resp, err = client.GoogleOAuth2Callback(context.Background(), req)
	if err != nil {
		xlog.Warn(err.Error())
	}
	return
}
