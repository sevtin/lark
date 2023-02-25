package service

import (
	"context"
	"lark/pkg/common/xjwt"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_auth"
)

func (s *authService) RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp, _ error) {
	resp = &pb_auth.RefreshTokenResp{}
	var (
		refreshToken *xjwt.JwtToken
		accessToken  *xjwt.JwtToken
		sessionId    string
		err          error
	)
	refreshToken, err = xjwt.Decode(req.RefreshToken)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_JWT_TOKEN_ERR, ERROR_AUTH_JWT_TOKEN_ERR)
		xlog.Warn(ERROR_CODE_AUTH_JWT_TOKEN_ERR, ERROR_AUTH_JWT_TOKEN_ERR, err.Error())
		return
	}
	sessionId, err = s.authCache.GetRefreshTokenSessionId(refreshToken.Uid, refreshToken.Platform)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_REDIS_GET_FAILED, ERROR_AUTH_REDIS_GET_FAILED)
		xlog.Warn(ERROR_CODE_AUTH_REDIS_GET_FAILED, ERROR_AUTH_REDIS_GET_FAILED, err.Error())
		return
	}
	if sessionId != refreshToken.SessionId {
		resp.Set(ERROR_CODE_AUTH_JWT_SESSION_ID_ERR, ERROR_AUTH_JWT_SESSION_ID_ERR)
		return
	}
	accessToken, err = xjwt.CreateToken(refreshToken.Uid, refreshToken.Platform, true, constant.CONST_DURATION_SHA_JWT_ACCESS_TOKEN_EXPIRE_IN_SECOND)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_GENERATE_TOKEN_FAILED, ERROR_AUTH_GENERATE_TOKEN_FAILED)
		xlog.Warn(ERROR_CODE_AUTH_GENERATE_TOKEN_FAILED, ERROR_AUTH_GENERATE_TOKEN_FAILED, err.Error())
		return
	}
	err = s.authCache.SetAccessTokenSessionId(refreshToken.Uid, refreshToken.Platform, accessToken.SessionId)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_REDIS_SET_FAILED, ERROR_AUTH_REDIS_SET_FAILED)
		xlog.Warn(ERROR_CODE_AUTH_REDIS_SET_FAILED, ERROR_AUTH_REDIS_SET_FAILED, err.Error())
		return
	}
	resp.AccessToken = &pb_auth.Token{
		Token:  accessToken.Token,
		Expire: accessToken.Expire,
	}
	return
}
