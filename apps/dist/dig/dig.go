package dig

import (
	"go.uber.org/dig"
	"lark/apps/dist/internal/config"
	"lark/apps/dist/internal/server"
	"lark/apps/dist/internal/server/dist"
	"lark/apps/dist/internal/service"
	"lark/domain/cache"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(dist.NewDistServer)
	container.Provide(service.NewDistService)
	container.Provide(cache.NewServerMgrCache)
	container.Provide(cache.NewConvoCache)
	container.Provide(cache.NewChatMemberCache)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
