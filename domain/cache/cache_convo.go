package cache

import (
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/utils"
)

// 弃用
type ConvoCache interface {
	MGetSeqIdTsList(chatIdList []string) (seqIdTsList []string, err error)
}

type convoCache struct {
}

func NewConvoCache() ConvoCache {
	return &convoCache{}
}

func (c *convoCache) MGetSeqIdTsList(chatIdList []string) (seqIdTsList []string, err error) {
	var (
		index  int
		chatId string
		cid    int64
		keys   = make([]string, len(chatIdList))
	)
	for index, chatId = range chatIdList {
		cid, _ = utils.ToInt64(chatId)
		keys[index] = constant.RK_MSG_SEQ_TS + utils.GetHashTagKey(cid)
	}
	seqIdTsList, err = xredis.CMGet(keys)
	return
}
