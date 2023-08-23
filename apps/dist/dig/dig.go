package dig

import (
	"go.uber.org/dig"
	"lark/apps/dist/internal/config"
	"lark/apps/dist/internal/server"
	"lark/apps/dist/internal/server/dist"
	"lark/apps/dist/internal/service"
	"lark/domain/cache"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(dist.NewDistServer)
	Provide(service.NewDistService)
	Provide(cache.NewServerMgrCache)
	Provide(cache.NewConvoCache)
	Provide(cache.NewChatMemberCache)
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
