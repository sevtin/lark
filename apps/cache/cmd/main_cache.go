package main

import (
	"lark/apps/cache/dig"
	"lark/apps/cache/internal/config"
	"lark/pkg/commands"
	"lark/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}
