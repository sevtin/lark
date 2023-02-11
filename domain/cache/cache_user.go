package cache

import (
	"github.com/go-redis/redis/v9"
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
	DelUserInfo(prefix string, uid int64) (err error)

	GetBasicUserInfo(uid int64) (info *pb_user.BasicUserInfo, err error)
	SetBasicUserInfo(info *pb_user.BasicUserInfo) (err error)
	GetBasicUserInfoList(uids []int64) (list []interface{}, err error)
	SetBasicUserInfoList(prefix string, list []*pb_user.BasicUserInfo) (err error)
	GetUserServerList(prefix string, uids []int64) (srvMaps map[int64]int64, notUids []int64, err error)
	SetUserServerList(prefix string, list []*pb_user.UserServerId) (err error)

	SignOut(prefix string, uid int64, platform pb_enum.PLATFORM_TYPE) (err error)
	SetServerId(uid int64, serverId int64) (err error)
	GetServerId(uid int64) (serverId string, err error)
	SetUserAndServer(prefix string, info *pb_user.UserInfo, serverId int64) (err error)
}

type userCache struct {
}

func NewUserCache() UserCache {
	return &userCache{}
}

func (c *userCache) GetUserInfo(uid int64) (info *pb_user.UserInfo, err error) {
	var (
		key = constant.RK_SYNC_USER_INFO + utils.Int64ToStr(uid)
	)
	info = &pb_user.UserInfo{Avatar: &pb_user.AvatarInfo{}}
	err = Get(key, info)
	return
}

func (c *userCache) SetUserInfo(info *pb_user.UserInfo) (err error) {
	var (
		key = constant.RK_SYNC_USER_INFO + utils.Int64ToStr(info.Uid)
	)
	err = Set(key, info, constant.CONST_DURATION_USER_INFO_SECOND)
	return
}

func (c *userCache) DelUserInfo(prefix string, uid int64) (err error) {
	var (
		key1 = prefix + constant.RK_SYNC_USER_INFO + utils.Int64ToStr(uid)
		key2 = prefix + constant.RK_SYNC_BASIC_USER_INFO + utils.Int64ToStr(uid)
	)
	err = xredis.Dels(key1, key2)
	return
}

func (c *userCache) GetBasicUserInfo(uid int64) (info *pb_user.BasicUserInfo, err error) {
	var (
		key = constant.RK_SYNC_BASIC_USER_INFO + utils.Int64ToStr(uid)
	)
	info = &pb_user.BasicUserInfo{}
	err = Get(key, info)
	return
}

func (c *userCache) SetBasicUserInfo(info *pb_user.BasicUserInfo) (err error) {
	var (
		key = constant.RK_SYNC_BASIC_USER_INFO + utils.Int64ToStr(info.Uid)
	)
	err = Set(key, info, constant.CONST_DURATION_BASIC_USER_INFO_SECOND)
	return
}

func (c *userCache) GetBasicUserInfoList(uids []int64) (list []interface{}, err error) {
	var (
		keys  = make([]string, len(uids))
		index int
		uid   int64
	)
	for index, uid = range uids {
		keys[index] = utils.Int64ToStr(uid)
	}
	list, err = Gets(keys, pb_user.BasicUserInfo{})
	return
}

func (c *userCache) SetBasicUserInfoList(prefix string, list []*pb_user.BasicUserInfo) (err error) {
	var (
		keys    = make([]string, len(list))
		vals    = make([]interface{}, len(list)+1)
		index   int
		srv     *pb_user.BasicUserInfo
		jsonStr string
	)
	vals[0] = constant.CONST_DURATION_SHA_BASIC_USER_INFO_SECOND
	for index, srv = range list {
		jsonStr, err = utils.Marshal(srv)
		if err != nil {
			xlog.Warn(ERROR_CODE_CACHE_PROTOCOL_MARSHAL_ERR, ERROR_CACHE_PROTOCOL_MARSHAL_ERR, err.Error())
			return
		}
		keys[index] = prefix + constant.RK_SYNC_BASIC_USER_INFO + utils.Int64ToStr(srv.Uid)
		vals[index+1] = jsonStr
	}
	return xredis.EvalSha(xredis.SHA_MSET_EXPIRE, keys, vals)
}

func (c *userCache) SetUserServerList(prefix string, list []*pb_user.UserServerId) (err error) {
	if len(list) == 0 {
		return
	}
	var (
		keys  = make([]string, len(list))
		vals  = make([]interface{}, len(list)+1)
		index int
		srv   *pb_user.UserServerId
	)
	vals[0] = constant.CONST_DURATION_SHA_USER_INFO_SECOND
	for index, srv = range list {
		keys[index] = prefix + constant.RK_SYNC_USER_SERVER + utils.Int64ToStr(srv.Uid)
		vals[index+1] = srv.ServerId
	}
	return xredis.EvalSha(xredis.SHA_MSET_EXPIRE, keys, vals)
}

func (c *userCache) GetUserServerList(prefix string, uids []int64) (srvMaps map[int64]int64, notUids []int64, err error) {
	srvMaps = make(map[int64]int64)
	notUids = uids
	var (
		uidList  = make([]string, len(uids))
		index    int
		uid      int64
		values   []interface{}
		val      interface{}
		serverId int64
		uidMaps  = make(map[int64]int64)
	)
	if len(uids) == 0 {
		return
	}
	for index, uid = range uids {
		uidMaps[uid] = uid
		uidList[index] = prefix + constant.RK_SYNC_USER_SERVER + utils.Int64ToStr(uid)
	}
	values, err = xredis.MGet(uidList...)
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

func (c *userCache) SignOut(prefix string, uid int64, platform pb_enum.PLATFORM_TYPE) (err error) {
	var (
		uidStr      = utils.Int64ToStr(uid)
		platformStr = utils.Int32ToStr(int32(platform))
		key1        = prefix + constant.RK_SYNC_USER_ACCESS_TOKEN_SESSION_ID + uidStr + ":" + platformStr
		key2        = prefix + constant.RK_SYNC_USER_REFRESH_TOKEN_SESSION_ID + uidStr + ":" + platformStr
		key3        = prefix + constant.RK_SYNC_USER_SERVER + uidStr
	)
	err = xredis.Dels(key1, key2, key3)
	if err != nil {
		return
	}
	return
}

func (c *userCache) SetServerId(uid int64, serverId int64) (err error) {
	var (
		key = constant.RK_SYNC_USER_SERVER + utils.ToString(uid)
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
		key = constant.RK_SYNC_USER_SERVER + utils.ToString(uid)
	)
	serverId, err = xredis.Get(key)
	if err != nil {
		if err == redis.Nil {
			err = nil
		} else {
			xlog.Warn(ERROR_CODE_CACHE_REDIS_GET_FAILED, ERROR_CACHE_REDIS_GET_FAILED, err.Error())
		}
	}
	return
}

func (c *userCache) SetUserAndServer(prefix string, info *pb_user.UserInfo, serverId int64) (err error) {
	var (
		uid  = utils.Int64ToStr(info.Uid)
		val  string
		keys = make([]string, 2)
		vals = make([]interface{}, 4)
	)
	keys[0] = prefix + constant.RK_SYNC_USER_INFO + uid
	keys[1] = prefix + constant.RK_SYNC_USER_SERVER + uid
	val, err = utils.Marshal(info)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_PROTOCOL_MARSHAL_ERR, ERROR_CACHE_PROTOCOL_MARSHAL_ERR, err.Error())
		return
	}
	vals[0] = val
	vals[1] = constant.CONST_DURATION_SHA_USER_INFO_SECOND
	vals[2] = serverId
	vals[3] = constant.CONST_DURATION_SHA_USER_SERVER_ID_SECOND
	return xredis.EvalSha(xredis.SHA_MULTIPLE_SET_EXPIRE, keys, vals)
}
