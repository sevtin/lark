package dig

import (
	"go.uber.org/dig"
	"lark/apps/auth/internal/config"
	"lark/apps/auth/internal/server"
	"lark/apps/auth/internal/server/auth"
	"lark/apps/auth/internal/service"
	"lark/domain/cache"
	"lark/domain/repo"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(auth.NewAuthServer)
	container.Provide(service.NewAuthService)
	container.Provide(repo.NewAuthRepository)
	container.Provide(repo.NewAvatarRepository)
	container.Provide(repo.NewUserRepository)
	container.Provide(repo.NewChatMemberRepository)
	container.Provide(cache.NewAuthCache)
	container.Provide(cache.NewUserCache)
	container.Provide(cache.NewServerMgrCache)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
