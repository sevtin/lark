package dig

import (
	"go.uber.org/dig"
	"lark/apps/lbs/internal/config"
	"lark/apps/lbs/internal/server"
	"lark/apps/lbs/internal/server/lbs"
	"lark/apps/lbs/internal/service"
	"lark/domain/cache"
	"lark/domain/mrepo"
	"lark/domain/repo"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(lbs.NewLbsServer)
	container.Provide(service.NewLbsService)
	container.Provide(mrepo.NewLbsRepository)
	container.Provide(repo.NewUserRepository)
	container.Provide(cache.NewUserCache)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
