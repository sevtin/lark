package cache

import (
	"context"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_red_env"
	"lark/pkg/utils"
	"strings"
)

type RedEnvelopeCache interface {
	SetRedEnvelopeInfo(envId int64, info *pb_red_env.RedEnvelopeInfo) (err error)
	SetRedEnvelope(envId int64, info *pb_red_env.RedEnvelopeInfo, envStatus int) (err error)
	GetRedEnvelopeInfo(envId int64) (info *pb_red_env.RedEnvelopeInfo, err error)
	SetRedEnvelopeStatus(envId int64, status int32) (err error)
	GetRedEnvelopeStatus(envId int64) (int, error)
	GetRedEnvelopeDist(envId int64, uid int64) (dist *pb_red_env.RedEnvelopeDistribution, err error)
	RollbackRedEnvelopeDist(dist *pb_red_env.RedEnvelopeDistribution) (err error)
	GenerateRedEnvelopeKey(envId int64) (redEnvKey string, err error)
	GetRedEnvelopeKey(envId int64) (redEnvKey string, err error)
}

type redEnvelopeCache struct {
}

func NewRedEnvelopeCache() RedEnvelopeCache {
	return &redEnvelopeCache{}
}

func (c *redEnvelopeCache) SetRedEnvelopeInfo(envId int64, info *pb_red_env.RedEnvelopeInfo) (err error) {
	var (
		key = constant.RK_SYNC_RED_ENV_INFO + utils.GetHashTagKey(envId)
	)
	err = Set(key, info, constant.CONST_DURATION_RED_ENVELOPE_EXPIRE_SECOND)
	return
}

func (c *redEnvelopeCache) SetRedEnvelope(envId int64, info *pb_red_env.RedEnvelopeInfo, envStatus int) (err error) {
	var (
		prefix = xredis.GetPrefix()
		tagKey = utils.GetHashTagKey(envId)
		key1   = prefix + constant.RK_SYNC_RED_ENV_STATUS + tagKey
		val1   = envStatus

		key2    = prefix + constant.RK_SYNC_RED_ENV_INFO + tagKey
		val2, _ = utils.Marshal(info)

		key3 = prefix + constant.RK_SYNC_RED_REMAIN_QUANTITY + tagKey
		val3 = info.RemainQuantity

		key4 = prefix + constant.RK_SYNC_RED_REMAIN_AMOUNT + tagKey
		val4 = info.RemainAmount

		pipe = xredis.Cli.Client.Pipeline()
	)
	pipe.Set(context.Background(), key1, val1, constant.CONST_DURATION_RED_ENVELOPE_EXPIRE_SECOND)
	pipe.Set(context.Background(), key2, val2, constant.CONST_DURATION_RED_ENVELOPE_EXPIRE_SECOND)
	pipe.Set(context.Background(), key3, val3, constant.CONST_DURATION_RED_ENVELOPE_EXPIRE_SECOND)
	pipe.Set(context.Background(), key4, val4, constant.CONST_DURATION_RED_ENVELOPE_EXPIRE_SECOND)
	_, err = pipe.Exec(context.Background())
	return
}

func (c *redEnvelopeCache) GetRedEnvelopeDist(envId int64, uid int64) (dist *pb_red_env.RedEnvelopeDistribution, err error) {
	var (
		value interface{}
		keys  = []string{utils.GetHashTagKey(envId)}
		args  = []interface{}{uid}
	)
	dist = new(pb_red_env.RedEnvelopeDistribution)
	dist.EnvId = envId
	dist.Uid = uid
	value, err = xredis.EvalShaResult(xredis.SHA_DISTRIBUTION_RED_ENVELOPE, keys, args)
	if err != nil {
		return
	}
	if value == nil {
		return
	}
	switch value.(type) {
	case string:
		splits := strings.Split(value.(string), ":")
		if len(splits) != 5 {
			return
		}
		dist.Status = splits[0]
		dist.Desc = splits[1]
		dist.RemainAmount, _ = utils.ToInt64(splits[2])
		dist.RemainQuantity, _ = utils.ToInt64(splits[3])
		dist.DistAmount, _ = utils.ToInt64(splits[4])
	}
	return
}

func (c *redEnvelopeCache) RollbackRedEnvelopeDist(dist *pb_red_env.RedEnvelopeDistribution) (err error) {
	var (
		keys = []string{utils.GetHashTagKey(dist.EnvId)}
		args = []interface{}{dist.Uid, dist.DistAmount}
	)
	err = xredis.EvalSha(xredis.SHA_ROLLBACK_RED_ENVELOPE, keys, args)
	return
}

func (c *redEnvelopeCache) GetRedEnvelopeInfo(envId int64) (info *pb_red_env.RedEnvelopeInfo, err error) {
	info = new(pb_red_env.RedEnvelopeInfo)
	var (
		key = constant.RK_SYNC_RED_ENV_INFO + utils.GetHashTagKey(envId)
	)
	err = Get(key, info)
	return
}

func (c *redEnvelopeCache) SetRedEnvelopeStatus(envId int64, status int32) (err error) {
	var (
		key = constant.RK_SYNC_RED_ENV_STATUS + utils.GetHashTagKey(envId)
	)
	return xredis.Set(key, status, constant.CONST_DURATION_RED_ENVELOPE_EXPIRE_SECOND)
}

func (c *redEnvelopeCache) GetRedEnvelopeStatus(envId int64) (int, error) {
	var (
		key = constant.RK_SYNC_RED_ENV_STATUS + utils.GetHashTagKey(envId)
	)
	return xredis.GetInt(key)
}

func (c *redEnvelopeCache) GenerateRedEnvelopeKey(envId int64) (redEnvKey string, err error) {
	var (
		key = constant.RK_SYNC_RED_ENV_KEY + utils.GetHashTagKey(envId)
	)
	redEnvKey = utils.NewUUID()
	err = Set(key, redEnvKey, constant.CONST_DURATION_RED_ENVELOPE_KEY_EXPIRE_SECOND)
	return
}

func (c *redEnvelopeCache) GetRedEnvelopeKey(envId int64) (redEnvKey string, err error) {
	var (
		key = constant.RK_SYNC_RED_ENV_KEY + utils.GetHashTagKey(envId)
	)
	redEnvKey = utils.NewUUID()
	redEnvKey, err = xredis.Get(key)
	return
}
