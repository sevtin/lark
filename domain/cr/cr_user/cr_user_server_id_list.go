package cr_user

import (
	"lark/domain/cache"
	"lark/domain/repo"
)

func GetUserServerList(userCache cache.UserCache, userRepo repo.UserRepository, uids []int64) (userSrvMaps map[int64]int64, err error) {
	var (
		notUids []int64
		//server  *pb_user.UserServerId
		//q       = entity.NewMysqlQuery()
		//list    []*pb_user.UserServerId
	)
	userSrvMaps, notUids, err = userCache.GetUserServerList(uids)
	if err != nil {
		return
	}
	if len(notUids) == 0 {
		return
	}
	// 超过120天未登录的用户,不主动进行推送
	/*
		q.SetFilter("uid IN(?)", uids)
		list, err = userRepo.UserServerList(q)
		if err != nil {
			return
		}
		for _, server = range list {
			userSrvMaps[server.Uid] = server.ServerId
		}
		xants.Submit(func() {
			userCache.SetUserServerList(list)
		})
	*/
	return
}
