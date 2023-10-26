package xredis

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	redsyncredis "github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
)

var (
	Cli *RedisClient
)

type RedisClient struct {
	Client   *redis.ClusterClient
	RedsSync *redsync.Redsync
	Prefix   string
}

func NewRedisClient(cfg *conf.Redis) *redis.ClusterClient {
	var (
		client   *redis.ClusterClient
		pool     redsyncredis.Pool
		redsSync *redsync.Redsync
		err      error
	)
	// 集群redis
	client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: cfg.Address,
	})
	// 判断是否能够链接到redis
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		xlog.Warn(err.Error())
		return nil
	}
	// redis 锁
	pool = goredis.NewPool(client)
	redsSync = redsync.New(pool)

	Cli = &RedisClient{client, redsSync, cfg.Prefix}
	return client
}
