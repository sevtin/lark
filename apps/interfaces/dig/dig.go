package dig

import (
	"go.uber.org/dig"
	"lark/apps/interfaces/internal/config"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
	provideAuth()
	provideUser()
	provideChat()
	provideChatMessage()
	provideChatMember()
	provideChatInvite()
	provideCache()
	provideConvo()
	provideLbs()
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
