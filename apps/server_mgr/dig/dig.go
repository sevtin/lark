package dig

import (
	"go.uber.org/dig"
	"lark/apps/server_mgr/internal/config"
	"lark/apps/server_mgr/internal/server"
	"lark/apps/server_mgr/internal/server/server_mgr"
	"lark/apps/server_mgr/internal/service"
	"lark/domain/cache"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(server_mgr.NewServerMgrServer)
	Provide(service.NewServerMgrService)
	Provide(cache.NewServerMgrCache)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}

func Provide(constructor interface{}, opts ...dig.ProvideOption) {
	err := container.Provide(constructor)
	if err != nil {
		log.Panic(err)
	}
}
