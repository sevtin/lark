package ctrl_red_env_receive

import (
	"lark/apps/interfaces/internal/service/svc_red_env_receive"
)

type RedEnvReceiveCtrl struct {
	redEnvReceiveService svc_red_env_receive.RedEnvReceiveService
}

func NewRedEnvReceiveCtrl(redEnvReceiveService svc_red_env_receive.RedEnvReceiveService) *RedEnvReceiveCtrl {
	return &RedEnvReceiveCtrl{redEnvReceiveService: redEnvReceiveService}
}
