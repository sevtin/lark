package main

import (
	"lark/apps/upload/internal/config"
	"lark/apps/upload/internal/server"
	"lark/pkg/commands"
	"lark/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	commands.Run(server.NewServer())
}
