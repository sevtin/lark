package svc_order

import (
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_order"
	order_client "lark/apps/order/client"
	"lark/pkg/xhttp"
)

type OrderService interface {
	CreateRedEnvelopeOrder(params *dto_order.CreateRedEnvelopeOrderReq, uid int64) (resp *xhttp.Resp)
}

type orderService struct {
	orderClient order_client.OrderClient
}

func NewOrderService(conf *config.Config) OrderService {
	orderClient := order_client.NewOrderClient(conf.Etcd, conf.OrderServer, conf.Jaeger, conf.Name)
	return &orderService{orderClient: orderClient}
}
