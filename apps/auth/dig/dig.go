package dig

import (
	"go.uber.org/dig"
	"lark/apps/auth/internal/config"
	"lark/apps/auth/internal/server"
	"lark/apps/auth/internal/server/auth"
	"lark/apps/auth/internal/service"
	"lark/domain/cache"
	"lark/domain/repo"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(auth.NewAuthServer)
	Provide(service.NewAuthService)
	Provide(repo.NewOauthUserRepository)
	Provide(repo.NewAvatarRepository)
	Provide(repo.NewUserRepository)
	Provide(repo.NewChatMemberRepository)
	Provide(cache.NewAuthCache)
	Provide(cache.NewUserCache)
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
