package main

import (
	"lark/apps/msg_hot/dig"
	"lark/apps/msg_hot/internal/config"
	"lark/pkg/commands"
	"lark/pkg/common/xmongo"
	"lark/pkg/common/xredis"
)

// 弃用
func init() {
	conf := config.GetConfig()
	xmongo.NewMongoClient(conf.Mongo)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}
