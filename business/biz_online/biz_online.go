package biz_online

import (
	"github.com/spf13/cast"
	"lark/domain/cache"
	"lark/domain/repo"
	"lark/pkg/common/xlog"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

type Online interface {
	UserOnline(uid int64, serverId int64, platform pb_enum.PLATFORM_TYPE) (srvId int64, next bool, err error)
}

type online struct {
	userCache cache.UserCache
	userRepo  repo.UserRepository
}

func NewOnline(userCache cache.UserCache, userRepo repo.UserRepository) Online {
	return &online{
		userCache: userCache,
		userRepo:  userRepo,
	}
}

func (b *online) UserOnline(uid int64, serverId int64, platform pb_enum.PLATFORM_TYPE) (srvId int64, next bool, err error) {
	var (
		oldSidStr string
		oldSidVal int64
	)
	// 1、获取旧的serverId
	oldSidStr, err = b.userCache.GetServerId(uid)
	if err != nil {
		xlog.Warnf("user online get serverid failed. err:%v, uid:%d", err.Error(), uid)
	}
	if oldSidStr == "" {
		var (
			user *pb_user.UserServerId
			w    = entity.NewMysqlQuery()
		)
		w.SetFilter("uid=?", uid)
		user, err = b.userRepo.UserServerId(w)
		if err != nil {
			xlog.Warnf("user online set serverid failed. err:%v, uid:%d", err.Error(), uid)
			return
		}
		oldSidVal = user.ServerId
	} else {
		oldSidVal = cast.ToInt64(oldSidStr)
	}
	// 2、得到新的serverId
	srvId = utils.NewServerId(oldSidVal, serverId, platform)
	// 3、是否是同一台服务器
	if oldSidVal == srvId {
		return
	}
	// 4、更新serverId缓存
	err = b.userCache.SetServerId(uid, serverId)
	if err != nil {
		xlog.Errorf("user online set serverid failed. err:%v, uid:%d, serverId:%d", err.Error(), uid, srvId)
		err = nil
	}
	// 5、更新数据库
	ent := entity.NewMysqlUpdate()
	ent.SetFilter("uid=?", uid)
	ent.Set("server_id", srvId)
	err = b.userRepo.UpdateUser(ent)
	if err != nil {
		xlog.Errorf("user online update user failed. err:%v, uid:%d, serverId:%d", err.Error(), uid, srvId)
		return
	}
	next = true
	return
}
