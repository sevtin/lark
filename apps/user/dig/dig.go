package dig

import (
	"go.uber.org/dig"
	"lark/apps/user/internal/config"
	"lark/apps/user/internal/server"
	"lark/apps/user/internal/server/user"
	"lark/apps/user/internal/service"
	"lark/domain/cache"
	"lark/domain/repo"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(user.NewUserServer)
	Provide(service.NewUserService)
	Provide(repo.NewUserRepository)
	Provide(repo.NewAvatarRepository)
	Provide(repo.NewChatMemberRepository)
	Provide(cache.NewChatMemberCache)
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
