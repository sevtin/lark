package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

type UserCache interface {
	GetUserInfo(uid int64) (info *pb_user.UserInfo, err error)
	SetUserInfo(info *pb_user.UserInfo) (err error)
	DelUserInfo(uid int64) (err error)
	GetBasicUserInfo(uid int64) (info *pb_user.BasicUserInfo, err error)
	SetBasicUserInfo(info *pb_user.BasicUserInfo) (err error)
	SetBasicUserInfoList(list []*pb_user.BasicUserInfo) (err error)
	SetUserServerList(list []*pb_user.UserServerId) (err error)
	GetUserServerList(uids []int64) (srvMaps map[int64]int64, notUids []int64, err error)
	SignOut(uid int64, platform pb_enum.PLATFORM_TYPE) (err error)
	GetServerId(uid int64) (serverId string, err error)
	SetServerId(uid int64, serverId int64) (err error)
	SetUserServer(uid int64, serverId int64) (err error)
}

type userCache struct {
}

func NewUserCache() UserCache {
	return &userCache{}
}

func (c *userCache) GetUserInfo(uid int64) (info *pb_user.UserInfo, err error) {
	var (
		key = constant.RK_SYNC_USER_INFO + utils.GetHashTagKey(uid)
	)
	info = &pb_user.UserInfo{Avatar: &pb_user.AvatarInfo{}}
	err = Get(key, info)
	return
}

func (c *userCache) SetUserInfo(info *pb_user.UserInfo) (err error) {
	var (
		key = constant.RK_SYNC_USER_INFO + utils.GetHashTagKey(info.Uid)
	)
	err = Set(key, info, constant.CONST_DURATION_USER_INFO_SECOND)
	return
}

func (c *userCache) DelUserInfo(uid int64) (err error) {
	var (
		htk = utils.GetHashTagKey(uid)
	)
	err = xredis.CUnlink([]string{constant.RK_SYNC_USER_INFO + htk, constant.RK_SYNC_BASIC_USER_INFO + htk})
	return
}

func (c *userCache) GetBasicUserInfo(uid int64) (info *pb_user.BasicUserInfo, err error) {
	var (
		key = constant.RK_SYNC_BASIC_USER_INFO + utils.GetHashTagKey(uid)
	)
	info = &pb_user.BasicUserInfo{}
	err = Get(key, info)
	return
}

func (c *userCache) SetBasicUserInfo(info *pb_user.BasicUserInfo) (err error) {
	var (
		key = constant.RK_SYNC_BASIC_USER_INFO + utils.GetHashTagKey(info.Uid)
	)
	err = Set(key, info, constant.CONST_DURATION_BASIC_USER_INFO_SECOND)
	return
}

func (c *userCache) SetBasicUserInfoList(list []*pb_user.BasicUserInfo) (err error) {
	var (
		srv     *pb_user.BasicUserInfo
		jsonStr string
		pipe    = xredis.Pipeline()
	)
	for _, srv = range list {
		jsonStr, err = utils.Marshal(srv)
		if err != nil {
			xlog.Warn(ERROR_CODE_CACHE_PROTOCOL_MARSHAL_ERR, ERROR_CACHE_PROTOCOL_MARSHAL_ERR, err.Error())
			return
		}
		pipe.Set(context.Background(),
			xredis.RealKey(constant.RK_SYNC_BASIC_USER_INFO+utils.GetHashTagKey(srv.Uid)),
			jsonStr,
			constant.CONST_DURATION_BASIC_USER_INFO_SECOND)
	}
	_, err = pipe.Exec(context.Background())
	return
}

func (c *userCache) SetUserServerList(list []*pb_user.UserServerId) (err error) {
	if len(list) == 0 {
		return
	}
	var (
		srv  *pb_user.UserServerId
		pipe = xredis.Pipeline()
	)
	for _, srv = range list {
		pipe.Set(context.Background(),
			xredis.RealKey(constant.RK_SYNC_USER_SERVER+utils.GetHashTagKey(srv.Uid)),
			srv.ServerId,
			constant.CONST_DURATION_USER_INFO_SECOND)
	}
	_, err = pipe.Exec(context.Background())
	return
}

func (c *userCache) GetUserServerList(uids []int64) (srvMaps map[int64]int64, notUids []int64, err error) {
	srvMaps = make(map[int64]int64)
	notUids = uids
	var (
		uidList  = make([]string, len(uids))
		index    int
		uid      int64
		values   []string
		val      interface{}
		serverId int64
		uidMaps  = make(map[int64]int64)
	)
	if len(uids) == 0 {
		return
	}
	for index, uid = range uids {
		uidMaps[uid] = uid
		uidList[index] = constant.RK_SYNC_USER_SERVER + utils.GetHashTagKey(uid)
	}
	values, err = xredis.CMGet(uidList)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_GET_FAILED, ERROR_CACHE_REDIS_GET_FAILED, err.Error())
		return
	}
	for index, val = range values {
		if val == nil {
			continue
		}
		serverId, _ = utils.ToInt64(val)
		if serverId == 0 {
			continue
		}
		uid = uids[index]
		srvMaps[uid] = serverId
		delete(uidMaps, uid)
	}
	notUids = make([]int64, len(uidMaps))
	if len(uidMaps) == 0 {
		return
	}
	index = -1
	for _, uid = range uidMaps {
		index++
		notUids[index] = uid
	}
	return
}

func (c *userCache) SignOut(uid int64, platform pb_enum.PLATFORM_TYPE) (err error) {
	var (
		htk         = utils.GetHashTagKey(uid)
		platformStr = utils.Int32ToStr(int32(platform))
		key1        = constant.RK_SYNC_USER_ACCESS_TOKEN_SESSION_ID + htk + ":" + platformStr
		key2        = constant.RK_SYNC_USER_REFRESH_TOKEN_SESSION_ID + htk + ":" + platformStr
		key3        = constant.RK_SYNC_USER_SERVER + htk
	)
	err = xredis.CUnlink([]string{key1, key2, key3})
	if err != nil {
		return
	}
	return
}

func (c *userCache) SetServerId(uid int64, serverId int64) (err error) {
	var (
		key = constant.RK_SYNC_USER_SERVER + utils.GetHashTagKey(uid)
	)
	// 更新serverId缓存
	err = Set(key, serverId, constant.CONST_DURATION_USER_SERVER_ID_SECOND)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_SET_FAILED, ERROR_CACHE_REDIS_SET_FAILED, err.Error())
	}
	return
}

func (c *userCache) GetServerId(uid int64) (serverId string, err error) {
	var (
		key = constant.RK_SYNC_USER_SERVER + utils.GetHashTagKey(uid)
	)
	serverId, err = xredis.Get(key)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_GET_FAILED, ERROR_CACHE_REDIS_GET_FAILED, err.Error())
	}
	return
}

func (c *userCache) GetServerIds(uidList []int64) (serverIds []string, err error) {
	var (
		i       int
		uid     int64
		keys    = make([]string, len(uidList))
		cmdList = make([]*redis.StringCmd, len(uidList))
		cmd     *redis.StringCmd
		pipe    = xredis.Pipeline()
	)
	for i, uid = range uidList {
		cmdList[i] = pipe.Get(context.Background(), xredis.RealKey(constant.RK_SYNC_USER_SERVER+utils.GetHashTagKey(uid)))
	}
	_, err = pipe.Exec(context.Background())
	if err != nil {
		return
	}
	serverIds = make([]string, len(keys))
	for i, cmd = range cmdList {
		serverIds[i] = cmd.Val()
	}
	return
}

//func (c *userCache) SetUserAndServer(info *pb_user.UserInfo, serverId int64) (err error) {
//	var (
//		val string
//		htk = utils.GetHashTagKey(info.Uid)
//	)
//	val, err = utils.Marshal(info)
//	if err != nil {
//		xlog.Warn(ERROR_CODE_CACHE_PROTOCOL_MARSHAL_ERR, ERROR_CACHE_PROTOCOL_MARSHAL_ERR, err.Error())
//		return
//	}
//	err = xredis.CSet([]string{constant.RK_SYNC_USER_INFO + htk, constant.RK_SYNC_USER_SERVER + htk},
//		[]interface{}{val, serverId},
//		constant.CONST_DURATION_USER_INFO_SECOND)
//	return
//}

func (c *userCache) SetUserServer(uid int64, serverId int64) (err error) {
	var (
		key = constant.RK_SYNC_USER_SERVER + utils.GetHashTagKey(uid)
	)
	Set(key, serverId, constant.CONST_DURATION_USER_INFO_SECOND)
	return
}
