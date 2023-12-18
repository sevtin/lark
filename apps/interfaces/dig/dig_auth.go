package dig

import (
	"lark/apps/interfaces/internal/service/svc_auth"
)

func init() {
	Provide(svc_auth.NewAuthService)
}
