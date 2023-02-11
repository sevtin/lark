package cache

import (
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/utils"
)

type ConvoCache interface {
	ZAdd(keys []string, vals []interface{}) (err error)
	ZRevRange(uid int64, start int64, stop int64) (list []string)
	MGetSeqIdList(prefix string, chatIdList []string) (seqIdList []interface{}, err error)
	ZMScore(uid int64, members ...string) (scores []float64)
}

type convoCache struct {
}

func NewConvoCache() ConvoCache {
	return &convoCache{}
}

func (c *convoCache) ZAdd(keys []string, vals []interface{}) (err error) {
	return xredis.EvalSha(xredis.SHA_ZADD_CONVERSATION, keys, vals)
}

func (c *convoCache) ZRevRange(uid int64, start int64, stop int64) (list []string) {
	list = xredis.ZRevRange(constant.RK_SYNC_CONVO_LIST+utils.Int64ToStr(uid), start, stop)
	return
}

func (c *convoCache) ZMScore(uid int64, members ...string) (scores []float64) {
	return xredis.ZMScore(constant.RK_SYNC_CONVO_LIST+utils.Int64ToStr(uid), members...)
}

func (c *convoCache) MGetSeqIdList(prefix string, chatIdList []string) (seqIdList []interface{}, err error) {
	var (
		index  int
		chatId string
		keys   = make([]string, len(chatIdList))
	)
	for index, chatId = range chatIdList {
		keys[index] = prefix + xredis.MSG_SEQ_ID + chatId
	}
	seqIdList, err = xredis.MGet(keys...)
	return
}
