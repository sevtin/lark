package cache

import (
	"lark/domain/pdo"
	"lark/pkg/constant"
	"lark/pkg/utils"
)

type OauthUserCache interface {
	SetOauthUserInfo(info *pdo.OauthUserInfo) (err error)
	GetOauthUserInfo(uid int64) (info *pdo.OauthUserInfo, err error)
	DeleteOauthUserInfo(uid int64) (info *pdo.OauthUserInfo, err error)
}

type oauthUserCache struct {
}

func NewOauthUserCache() OauthUserCache {
	return &oauthUserCache{}
}

func (c *oauthUserCache) SetOauthUserInfo(info *pdo.OauthUserInfo) (err error) {
	var (
		key = constant.RK_SYNC_USER_OAUTH_USER_INFO + utils.GetHashTagKey(info.Uid)
	)
	err = Set(key, info, constant.CONST_DURATION_USER_OAUTH_USER_INFO_SECOND)
	return
}

func (c *oauthUserCache) GetOauthUserInfo(uid int64) (info *pdo.OauthUserInfo, err error) {
	var (
		key = constant.RK_SYNC_USER_OAUTH_USER_INFO + utils.GetHashTagKey(uid)
	)
	info = &pdo.OauthUserInfo{}
	err = Get(key, info)
	return
}

func (c *oauthUserCache) DeleteOauthUserInfo(uid int64) (info *pdo.OauthUserInfo, err error) {
	var (
		key = constant.RK_SYNC_USER_OAUTH_USER_INFO + utils.GetHashTagKey(uid)
	)
	err = Delete(key)
	return
}
