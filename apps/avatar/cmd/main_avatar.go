package main

import (
	"lark/apps/avatar/dig"
	"lark/apps/avatar/internal/config"
	"lark/pkg/commands"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xredis"
)

// 弃用
func init() {
	conf := config.GetConfig()
	xmysql.NewMysqlClient(conf.Mysql)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}
