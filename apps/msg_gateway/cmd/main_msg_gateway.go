package main

import (
	"lark/apps/msg_gateway/dig"
	"lark/apps/msg_gateway/internal/config"
	"lark/pkg/commands"
	"lark/pkg/common/xredis"
	"runtime"
)

func init() {
	conf := config.GetConfig()
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}
