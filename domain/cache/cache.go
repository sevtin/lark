package cache

import (
	"lark/pkg/common/xlog"
	"lark/pkg/common/xredis"
	"lark/pkg/utils"
	"time"
)

const (
	ERROR_CODE_CACHE_REDIS_GET_FAILED       int32 = 801
	ERROR_CODE_CACHE_REDIS_SET_FAILED       int32 = 802
	ERROR_CODE_CACHE_REDIS_DELETE_FAILED    int32 = 803
	ERROR_CODE_CACHE_PROTOCOL_MARSHAL_ERR   int32 = 804
	ERROR_CODE_CACHE_PROTOCOL_UNMARSHAL_ERR int32 = 805
	ERROR_CODE_CACHE_SET_EXPIRE_FAILED      int32 = 806
	ERROR_CODE_CACHE_GET_SEQ_ID_FAILED      int32 = 807
)

const (
	ERROR_CACHE_REDIS_GET_FAILED       = "读取redis缓存失败"
	ERROR_CACHE_REDIS_SET_FAILED       = "缓存数据失败"
	ERROR_CACHE_REDIS_DELETE_FAILED    = "删除缓存数据失败"
	ERROR_CACHE_PROTOCOL_MARSHAL_ERR   = "协议序列化错误"
	ERROR_CACHE_PROTOCOL_UNMARSHAL_ERR = "协议反序列化错误"
	ERROR_CACHE_SET_EXPIRE_FAILED      = "设置过期时间失败"
	ERROR_CACHE_GET_SEQ_ID_FAILED      = "生成消息 Sequence ID 失败"
)

func Get(key string, out interface{}) (err error) {
	var (
		jsonStr string
	)
	jsonStr, err = xredis.Get(key)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_GET_FAILED, ERROR_CACHE_REDIS_GET_FAILED, key, err.Error())
		return
	}
	if jsonStr == "" {
		return
	}
	err = utils.Unmarshal(jsonStr, &out)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_PROTOCOL_UNMARSHAL_ERR, ERROR_CACHE_PROTOCOL_UNMARSHAL_ERR, err.Error())
	}
	return
}

// tp 不能为指针
func Gets(keys []string, tp interface{}) (list []interface{}, err error) {
	list = make([]interface{}, 0)
	var (
		vals []string
		val  string
	)
	vals, err = xredis.CMGet(keys)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_GET_FAILED, ERROR_CACHE_REDIS_GET_FAILED, keys, err.Error())
		return
	}
	for _, val = range vals {
		if val == "" {
			return
		}
		err = utils.Unmarshal(val, &tp)
		if err != nil {
			xlog.Warn(ERROR_CODE_CACHE_PROTOCOL_UNMARSHAL_ERR, ERROR_CACHE_PROTOCOL_UNMARSHAL_ERR, val, err.Error())
		}
		list = append(list, tp)
	}
	return
}

func Set(key string, in interface{}, expire time.Duration) (err error) {
	var (
		val interface{}
	)
	switch in.(type) {
	case string, int64, int, int32:
		val = in
	default:
		val, err = utils.Marshal(in)
		if err != nil {
			xlog.Warn(ERROR_CODE_CACHE_PROTOCOL_MARSHAL_ERR, ERROR_CACHE_PROTOCOL_MARSHAL_ERR, val, err.Error())
			return
		}
	}
	if val == nil {
		return
	}
	err = xredis.Set(key, val, expire)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_SET_FAILED, ERROR_CACHE_REDIS_SET_FAILED, val, err.Error())
	}
	return
}

func Delete(key string) (err error) {
	err = xredis.Del(key)
	if err != nil {
		xlog.Warn(ERROR_CODE_CACHE_REDIS_DELETE_FAILED, ERROR_CACHE_REDIS_DELETE_FAILED, key, err.Error())
	}
	return
}
