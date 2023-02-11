package dig

import (
	"lark/apps/interfaces/internal/service/svc_user"
)

func provideUser() {
	container.Provide(svc_user.NewUserService)
}
