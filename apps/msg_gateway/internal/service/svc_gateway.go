package service

type GatewayService interface {
}

type gatewayService struct {
}

func NewGatewayService() GatewayService {
	return &gatewayService{}
}
