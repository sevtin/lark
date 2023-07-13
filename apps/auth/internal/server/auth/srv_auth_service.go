package auth

import (
	"context"
	"lark/pkg/proto/pb_auth"
)

func (s *authServer) SignUp(ctx context.Context, req *pb_auth.SignUpReq) (resp *pb_auth.SignUpResp, err error) {
	return s.authService.SignUp(ctx, req)
}

func (s *authServer) SignIn(ctx context.Context, req *pb_auth.SignInReq) (resp *pb_auth.SignInResp, err error) {
	return s.authService.SignIn(ctx, req)
}

func (s *authServer) RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp, err error) {
	return s.authService.RefreshToken(ctx, req)
}

func (s *authServer) SignOut(ctx context.Context, req *pb_auth.SignOutReq) (resp *pb_auth.SignOutResp, err error) {
	return s.authService.SignOut(ctx, req)
}

func (s *authServer) GithubOAuth2Callback(ctx context.Context, req *pb_auth.GithubOAuth2CallbackReq) (resp *pb_auth.GithubOAuth2CallbackResp, err error) {
	return s.authService.GithubOAuth2Callback(ctx, req)
}
