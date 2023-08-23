package dig

import (
	"lark/apps/interfaces/internal/service/svc_auth"
)

func provideAuth() {
	Provide(svc_auth.NewAuthService)
}
