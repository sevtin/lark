package dig

import (
	"lark/apps/interfaces/internal/service/svc_auth"
)

func provideAuth() {
	container.Provide(svc_auth.NewAuthService)
}
