package dig

import (
	"go.uber.org/dig"
	"lark/apps/cache/internal/config"
	"lark/apps/cache/internal/server"
	srv_cache "lark/apps/cache/internal/server/cache"
	"lark/apps/cache/internal/service"
	"lark/domain/cache"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(cache.NewChatMemberCache)
	Provide(service.NewCacheService)
	Provide(srv_cache.NewCacheServer)
	Provide(server.NewServer)
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
