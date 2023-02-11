package main

import (
	"lark/apps/user/dig"
	"lark/apps/user/internal/config"
	"lark/pkg/commands"
	"lark/pkg/common/xes"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xmysql.NewMysqlClient(conf.Mysql)
	xredis.NewRedisClient(conf.Redis)
	xes.NewElasticsearchClient(conf.Elasticsearch)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}
