package cache

import (
	"context"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/utils"
	"strconv"
)

type AuthCache interface {
	SetAccessTokenSessionId(uid int64, platform int32, sessionId string) (err error)
	GetRefreshTokenSessionId(uid int64, platform int32) (val string, err error)
	SetSessionId(uid int64, platform int32, accessSessionId string, refreshSessionId string) (err error)
	GetOauthUserToken(uid int64) (token string, err error)
	SetOauthUserToken(uid int64, token string) (err error)
	DelOauthUserToken(uid int64, platform pb_enum.PLATFORM_TYPE) (err error)
	SetInvitationCode(uidHash string, uid string) (err error)
	GetInvitationCode(uidHash string) (uid string, err error)
}

type authCache struct {
}

func NewAuthCache() AuthCache {
	return &authCache{}
}

func (c *authCache) SetAccessTokenSessionId(uid int64, platform int32, sessionId string) (err error) {
	var (
		key = constant.RK_SYNC_USER_ACCESS_TOKEN_SESSION_ID + utils.GetHashTagKey(uid) + ":" + utils.Int32ToStr(platform)
	)
	return Set(key, sessionId, 0)
}

func (c *authCache) GetRefreshTokenSessionId(uid int64, platform int32) (val string, err error) {
	var (
		key = constant.RK_SYNC_USER_REFRESH_TOKEN_SESSION_ID + utils.GetHashTagKey(uid) + ":" + utils.Int32ToStr(platform)
	)
	val, err = xredis.Get(key)
	return
}

func (c *authCache) SetSessionId(uid int64, platform int32, accessSessionId string, refreshSessionId string) (err error) {
	var (
		htk         = utils.GetHashTagKey(uid)
		platformStr = utils.Int32ToStr(platform)
		pipe        = xredis.Pipeline()
	)
	pipe.Set(context.Background(), xredis.RealKey(constant.RK_SYNC_USER_ACCESS_TOKEN_SESSION_ID+htk+":"+platformStr), accessSessionId, constant.CONST_DURATION_JWT_ACCESS_TOKEN_EXPIRE_IN_SECOND)
	pipe.Set(context.Background(), xredis.RealKey(constant.RK_SYNC_USER_REFRESH_TOKEN_SESSION_ID+htk+":"+platformStr), refreshSessionId, constant.CONST_DURATION_JWT_REFRESH_TOKEN_EXPIRE_IN_SECOND)
	_, err = pipe.Exec(context.Background())
	return
}

func (c *authCache) GetOauthUserToken(uid int64) (token string, err error) {
	var (
		key = constant.RK_SYNC_OAUTH_USER_TOKEN + utils.GetHashTagKey(uid)
	)
	token, err = xredis.Get(key)
	return
}

func (c *authCache) SetOauthUserToken(uid int64, token string) (err error) {
	var (
		key = constant.RK_SYNC_OAUTH_USER_TOKEN + utils.GetHashTagKey(uid)
	)
	Set(key, token, constant.CONST_DURATION_OAUTH_USER_TOKEN_SECOND)
	return
}

func (c *authCache) DelOauthUserToken(uid int64, platform pb_enum.PLATFORM_TYPE) (err error) {
	var (
		htk  = utils.GetHashTagKey(uid)
		pf   = strconv.Itoa(int(platform))
		key1 = constant.RK_SYNC_OAUTH_USER_TOKEN + htk
		key2 = constant.RK_SYNC_USER_ACCESS_TOKEN + htk + ":" + pf
		key3 = constant.RK_SYNC_USER_ACCESS_TOKEN_SESSION_ID + htk + ":" + pf
		key4 = constant.RK_SYNC_USER_REFRESH_TOKEN_SESSION_ID + htk + ":" + pf
		pipe = xredis.Pipeline()
	)
	pipe.Del(context.Background(), key1)
	pipe.Del(context.Background(), key2)
	pipe.Del(context.Background(), key3)
	pipe.Del(context.Background(), key4)
	_, err = pipe.Exec(context.Background())
	return
}

func (c *authCache) SetInvitationCode(uidHash string, uid string) (err error) {
	var (
		key = constant.RK_SYNC_INVITATION_CODE + uidHash
	)
	Set(key, uid, constant.CONST_DURATION_INVITATION_CODE_SECOND)
	return
}

func (c *authCache) GetInvitationCode(uidHash string) (uid string, err error) {
	var (
		key = constant.RK_SYNC_INVITATION_CODE + uidHash
	)
	uid, err = xredis.Get(key)
	return
}
