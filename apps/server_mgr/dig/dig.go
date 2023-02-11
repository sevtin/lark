package dig

import (
	"go.uber.org/dig"
	"lark/apps/server_mgr/internal/config"
	"lark/apps/server_mgr/internal/server"
	"lark/apps/server_mgr/internal/server/server_mgr"
	"lark/apps/server_mgr/internal/service"
	"lark/domain/cache"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(server_mgr.NewServerMgrServer)
	container.Provide(service.NewServerMgrService)
	container.Provide(cache.NewServerMgrCache)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
