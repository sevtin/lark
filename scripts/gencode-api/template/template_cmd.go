package template

var CmdTemplate = ParseTemplate(`
package main

import (
	"lark/apps/apis/{{.PackageName}}/internal/config"
	"lark/apps/apis/{{.PackageName}}/internal/server"
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
	commands.Run(server.NewServer())
}
`)
