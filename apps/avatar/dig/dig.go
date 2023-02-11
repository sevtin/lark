package dig

import (
	"go.uber.org/dig"
	"lark/apps/avatar/internal/config"
	"lark/apps/avatar/internal/server"
	"lark/apps/avatar/internal/server/avatar"
	"lark/apps/avatar/internal/service"
	"lark/domain/repo"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	container.Provide(server.NewServer)
	container.Provide(avatar.NewAvatarServer)
	container.Provide(service.NewAvatarService)
	container.Provide(repo.NewAvatarRepository)
	container.Provide(repo.NewChatMemberRepository)
	container.Provide(repo.NewChatRepository)
	container.Provide(repo.NewUserRepository)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
