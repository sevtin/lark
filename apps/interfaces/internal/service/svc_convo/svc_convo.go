package svc_convo

import (
	convo_client "lark/apps/convo/client"
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_convo"
	"lark/pkg/xhttp"
)

type ConvoService interface {
	ConvoList(req *dto_convo.ConvoListReq, uid int64) (resp *xhttp.Resp)
	ConvoChatSeqList(params *dto_convo.ConvoChatSeqListReq, uid int64) (resp *xhttp.Resp)
}

type convoService struct {
	cfg         *config.Config
	convoClient convo_client.ConvoClient
}

func NewConvoService(conf *config.Config) ConvoService {
	convoClient := convo_client.NewConvoClient(conf.Etcd, conf.ConvoServer, conf.Jaeger, conf.Name)
	return &convoService{cfg: conf, convoClient: convoClient}
}
