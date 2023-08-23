package dig

import (
	"lark/apps/interfaces/internal/service/svc_red_env"
)

func provideRedEnv() {
	Provide(svc_red_env.NewRedEnvService)
}
