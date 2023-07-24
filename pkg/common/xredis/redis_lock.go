package xredis

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	"time"
)

func NewMutex(key string, options ...redsync.Option) *redsync.Mutex {
	key = RealKey(key)
	return Cli.RedsSync.NewMutex(key, options...)
}

func SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	key = RealKey(key)
	return Cli.Client.SetNX(context.Background(), key, value, expiration).Result()
}
