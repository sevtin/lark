package dig

import (
	"lark/apps/interfaces/internal/service/svc_user"
)

func provideUser() {
	Provide(svc_user.NewUserService)
}
