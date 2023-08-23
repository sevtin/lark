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
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(lbs.NewLbsServer)
	Provide(service.NewLbsService)
	Provide(mrepo.NewLbsRepository)
	Provide(repo.NewUserRepository)
	Provide(repo.NewUserLocationRepository)
	Provide(cache.NewUserCache)
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
