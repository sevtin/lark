package svc_payment

import (
	"lark/apps/interfaces/internal/config"
	payment_client "lark/apps/payment/client"
)

type PaymentService interface {
}

type paymentService struct {
	paymentClient payment_client.PaymentClient
}

func NewPaymentService(conf *config.Config) PaymentService {
	paymentClient := payment_client.NewPaymentClient(conf.Etcd, nil, conf.Jaeger, conf.Name)
	return &paymentService{paymentClient: paymentClient}
}
