package svc_red_env_receive

import (
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_red_env_receive"
	red_env_receive_client "lark/apps/red_env_receive/client"
	"lark/pkg/xhttp"
)

type RedEnvReceiveService interface {
	GrabRedEnvelope(params *dto_red_env_receive.GrabRedEnvelopeReq, uid int64) (resp *xhttp.Resp)
	OpenRedEnvelope(params *dto_red_env_receive.OpenRedEnvelopeReq, uid int64) (resp *xhttp.Resp)
}

type redEnvReceiveService struct {
	redEnvReceiveClient red_env_receive_client.RedEnvReceiveClient
}

func NewRedEnvReceiveService(conf *config.Config) RedEnvReceiveService {
	redEnvReceiveClient := red_env_receive_client.NewRedEnvReceiveClient(conf.Etcd, conf.RedEnvReceiveServer, conf.Jaeger, conf.Name)
	return &redEnvReceiveService{redEnvReceiveClient: redEnvReceiveClient}
}
