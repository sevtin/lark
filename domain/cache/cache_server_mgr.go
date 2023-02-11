package cache

import (
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
)

type ServerMgrCache interface {
	ZAddMsgGateway(score float64, member string) (err error)
	ZRemMsgGateway(member string) (err error)
	ZRevRangeMsgGateway(start int64, stop int64) []string
	ZRangeMsgGateway(start int64, stop int64) []string
	SetVersion(server string, version int) (err error)
	GetVersion(server string) (int, error)
	GetIndex(member string) (int64, error)
}

type serverMgrCache struct {
}

func NewServerMgrCache() ServerMgrCache {
	return &serverMgrCache{}
}

func (c *serverMgrCache) ZAddMsgGateway(score float64, member string) (err error) {
	return xredis.ZAdd(constant.RK_SYNC_SERVER_MSG_GATEWAY, score, member)
}

func (c *serverMgrCache) ZRemMsgGateway(member string) (err error) {
	return xredis.ZRem(constant.RK_SYNC_SERVER_MSG_GATEWAY, member)
}

func (c *serverMgrCache) ZRevRangeMsgGateway(start int64, stop int64) []string {
	return xredis.ZRevRange(constant.RK_SYNC_SERVER_MSG_GATEWAY, start, stop)
}

func (c *serverMgrCache) ZRangeMsgGateway(start int64, stop int64) []string {
	return xredis.ZRange(constant.RK_SYNC_SERVER_MSG_GATEWAY, start, stop)
}

func (c *serverMgrCache) SetVersion(server string, version int) (err error) {
	key := constant.RK_SYNC_SERVER_MGR + server
	return Set(key, version, constant.CONST_DURATION_SERVER_MGR_SECOND)
}

func (c *serverMgrCache) GetVersion(server string) (int, error) {
	key := constant.RK_SYNC_SERVER_MGR + server
	return xredis.GetInt(key)
}

func (c *serverMgrCache) GetIndex(member string) (int64, error) {
	return xredis.ZRank(constant.RK_SYNC_SERVER_MSG_GATEWAY, member)
}
