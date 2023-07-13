package main

import (
	"lark/apps/lbs/dig"
	"lark/apps/lbs/internal/config"
	"lark/pkg/commands"
	"lark/pkg/common/xmongo"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xmongo.NewMongoClient(conf.Mongo)
	xmysql.NewMysqlClient(conf.Mysql)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}
