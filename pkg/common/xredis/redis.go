package xredis

import (
	"lark/pkg/conf"
)

var (
	cli RedisIface
)

func NewRedisClient(cfg *conf.Redis) RedisIface {
	if cfg.Single == true {
		cli = NewRedisSingle(cfg)
		return cli
	}
	cli = NewRedisCluster(cfg)
	return cli
}
