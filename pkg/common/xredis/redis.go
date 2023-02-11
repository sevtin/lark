package xredis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/go-redsync/redsync/v4"
	redsyncredis "github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"lark/pkg/conf"
)

var (
	cli *RedisClient
)

type RedisClient struct {
	client   *redis.Client
	RedsSync *redsync.Redsync
	Prefix   string
}

func NewRedisClient(cfg *conf.Redis) *redis.Client {
	var (
		client   *redis.Client
		pool     redsyncredis.Pool
		redsSync *redsync.Redsync
		err      error
	)
	// 单机redis
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Address[0],
		Password: cfg.Password,
		DB:       cfg.Db,
	})
	// 判断是否能够链接到redis
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	// redis 锁
	pool = goredis.NewPool(client)
	redsSync = redsync.New(pool)

	cli = &RedisClient{client, redsSync, cfg.Prefix}
	return client
}
