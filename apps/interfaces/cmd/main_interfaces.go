package main

import (
	"lark/apps/interfaces/internal/config"
	_ "lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/server"
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
