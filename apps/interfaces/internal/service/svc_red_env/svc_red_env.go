package svc_red_env

import (
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_red_env"
	red_env_client "lark/apps/red_env/client"
	"lark/pkg/xhttp"
)

type RedEnvService interface {
	GiveRedEnvelope(params *dto_red_env.GiveRedEnvelopeReq, uid int64) (resp *xhttp.Resp)
}

type redEnvService struct {
	redEnvClient red_env_client.RedEnvClient
}

func NewRedEnvService(conf *config.Config) RedEnvService {
	redEnvClient := red_env_client.NewRedEnvClient(conf.Etcd, conf.RedEnvServer, conf.Jaeger, conf.Name)
	return &redEnvService{redEnvClient: redEnvClient}
}
