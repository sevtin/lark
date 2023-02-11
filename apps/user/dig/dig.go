package dig

import (
	"go.uber.org/dig"
	"lark/apps/user/internal/config"
	"lark/apps/user/internal/server"
	"lark/apps/user/internal/server/user"
	"lark/apps/user/internal/service"
	"lark/domain/cache"
	"lark/domain/repo"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(user.NewUserServer)
	container.Provide(service.NewUserService)
	container.Provide(repo.NewUserRepository)
	container.Provide(repo.NewAvatarRepository)
	container.Provide(repo.NewChatMemberRepository)
	container.Provide(cache.NewChatMemberCache)
	container.Provide(cache.NewUserCache)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
