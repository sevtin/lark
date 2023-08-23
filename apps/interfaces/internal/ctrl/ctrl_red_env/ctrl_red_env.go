package ctrl_red_env

import (
	"lark/apps/interfaces/internal/service/svc_red_env"
)

type RedEnvCtrl struct {
	redEnvService svc_red_env.RedEnvService
}

func NewRedEnvCtrl(redEnvService svc_red_env.RedEnvService) *RedEnvCtrl {
	return &RedEnvCtrl{redEnvService: redEnvService}
}
