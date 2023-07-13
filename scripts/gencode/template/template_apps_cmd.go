package template

var AppsCmdTemplate = ParseTemplate(`
package main

import (
	"lark/apps/{{.PackageName}}/dig"
	"lark/apps/{{.PackageName}}/internal/config"
	"lark/pkg/commands"
	"lark/pkg/common/xmysql"
	"lark/pkg/common/xredis"
)

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
`)
