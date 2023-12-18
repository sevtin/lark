package dig

import (
	"lark/apps/interfaces/internal/service/svc_red_env"
)

func init() {
	Provide(svc_red_env.NewRedEnvService)
}
