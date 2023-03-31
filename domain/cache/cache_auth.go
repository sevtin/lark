package cache

import (
	"context"
	"github.com/go-redis/redis/v9"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/utils"
)

type AuthCache interface {
	SetAccessTokenSessionId(uid int64, platform int32, sessionId string) (err error)
	GetRefreshTokenSessionId(uid int64, platform int32) (val string, err error)
	SetSessionId(uid int64, platform int32, accessSessionId string, refreshSessionId string) (err error)
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
	if err == redis.Nil {
		err = nil
	}
	return
}

func (c *authCache) SetSessionId(uid int64, platform int32, accessSessionId string, refreshSessionId string) (err error) {
	var (
		htk         = utils.GetHashTagKey(uid)
		platformStr = utils.Int32ToStr(platform)
		pipe        = xredis.Cli.Client.Pipeline()
	)
	pipe.Set(context.Background(), xredis.RealKey(constant.RK_SYNC_USER_ACCESS_TOKEN_SESSION_ID+htk+":"+platformStr), accessSessionId, constant.CONST_DURATION_JWT_ACCESS_TOKEN_EXPIRE_IN_SECOND)
	pipe.Set(context.Background(), xredis.RealKey(constant.RK_SYNC_USER_REFRESH_TOKEN_SESSION_ID+htk+":"+platformStr), refreshSessionId, constant.CONST_DURATION_JWT_REFRESH_TOKEN_EXPIRE_IN_SECOND)
	_, err = pipe.Exec(context.Background())
	return
}
