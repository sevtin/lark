package cr_user

import (
	"lark/domain/cache"
	"lark/domain/pdo"
	"lark/domain/repo"
	"lark/pkg/entity"
)

func GetOauthUserInfo(oauthUserCache cache.OauthUserCache, userRepo repo.OauthUserRepository, uid int64) (info *pdo.OauthUserInfo, err error) {
	info, _ = oauthUserCache.GetOauthUserInfo(uid)
	if info.Uid > 0 {
		return
	}
	var (
		q = entity.NewMysqlQuery()
	)
	q.SetFilter("uid=?", uid)
	info, err = userRepo.GetOAuthUserInfo(q)
	if err != nil {
		return
	}
	oauthUserCache.SetOauthUserInfo(info)
	return
}
