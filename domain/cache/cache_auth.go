package cache

import (
	"github.com/go-redis/redis/v9"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/utils"
)

type AuthCache interface {
	SetAccessTokenSessionId(uid int64, platform int32, sessionId string) (err error)
	GetRefreshTokenSessionId(uid int64, platform int32) (val string, err error)
	SetSessionId(prefix string, uid int64, platform int32, accessSessionId string, refreshSessionId string) (err error)
}

type authCache struct {
}

func NewAuthCache() AuthCache {
	return &authCache{}
}

func (c *authCache) SetAccessTokenSessionId(uid int64, platform int32, sessionId string) (err error) {
	var (
		key = constant.RK_SYNC_USER_ACCESS_TOKEN_SESSION_ID + utils.Int64ToStr(uid) + ":" + utils.Int32ToStr(platform)
	)
	return Set(key, sessionId, 0)
}

func (c *authCache) GetRefreshTokenSessionId(uid int64, platform int32) (val string, err error) {
	var (
		key = constant.RK_SYNC_USER_REFRESH_TOKEN_SESSION_ID + utils.Int64ToStr(uid) + ":" + utils.Int32ToStr(platform)
	)
	val, err = xredis.Get(key)
	if err == redis.Nil {
		err = nil
	}
	return
}

func (c *authCache) SetSessionId(prefix string, uid int64, platform int32, accessSessionId string, refreshSessionId string) (err error) {
	var (
		uidStr      = utils.Int64ToStr(uid)
		platformStr = utils.Int32ToStr(platform)
		keys        = make([]string, 2)
		vals        = make([]interface{}, 4)
	)
	keys[0] = prefix + constant.RK_SYNC_USER_ACCESS_TOKEN_SESSION_ID + uidStr + ":" + platformStr
	keys[1] = prefix + constant.RK_SYNC_USER_REFRESH_TOKEN_SESSION_ID + uidStr + ":" + platformStr

	vals[0] = accessSessionId
	vals[1] = constant.CONST_DURATION_SHA_JWT_ACCESS_TOKEN_EXPIRE_IN_SECOND
	vals[2] = refreshSessionId
	vals[3] = constant.CONST_DURATION_SHA_JWT_REFRESH_TOKEN_EXPIRE_IN_SECOND
	return xredis.EvalSha(xredis.SHA_MULTIPLE_SET_EXPIRE, keys, vals)
}
