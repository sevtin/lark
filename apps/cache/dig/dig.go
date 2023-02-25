package dig

import (
	"go.uber.org/dig"
	"lark/apps/cache/internal/config"
	"lark/apps/cache/internal/server"
	srv_cache "lark/apps/cache/internal/server/cache"
	"lark/apps/cache/internal/service"
	"lark/domain/cache"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(cache.NewChatMemberCache)
	container.Provide(service.NewCacheService)
	container.Provide(srv_cache.NewCacheServer)
	container.Provide(server.NewServer)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
