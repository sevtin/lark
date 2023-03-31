package xredis

import (
	"github.com/go-redsync/redsync/v4"
)

func NewMutex(key string, options ...redsync.Option) *redsync.Mutex {
	key = RealKey(key)
	return Cli.RedsSync.NewMutex(key, options...)
}
