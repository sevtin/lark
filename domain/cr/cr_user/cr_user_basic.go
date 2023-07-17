package cr_user

import (
	"lark/domain/cache"
	"lark/domain/pdo"
	"lark/domain/repo"
	"lark/pkg/common/xants"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_user"
)

func GetBasicUserInfo(userCache cache.UserCache, userRepo repo.UserRepository, uid int64) (info *pb_user.BasicUserInfo, err error) {
	var (
		user = new(pdo.BasicUserInfo)
		q    = entity.NewMysqlQuery()
	)
	info, err = userCache.GetBasicUserInfo(uid)
	if err != nil {
		return
	}
	if info.Uid > 0 {
		return
	}
	q.Fields = user.GetFields()
	q.SetFilter("uid=?", uid)
	err = userRepo.QueryUser(q, user)
	if err != nil {
		return
	}
	if user.Uid == 0 {
		err = ERROR_CR_USER_QUERY_FAILED
		return
	}
	info = &pb_user.BasicUserInfo{
		Uid:      user.Uid,
		LarkId:   user.LarkId,
		Nickname: user.Nickname,
		Gender:   user.Gender,
		BirthTs:  user.BirthTs,
		CityId:   user.CityId,
		Avatar:   user.Avatar,
	}
	xants.Submit(func() {
		userCache.SetBasicUserInfo(info)
	})
	return
}
