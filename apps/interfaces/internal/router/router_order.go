package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_order"
	"lark/apps/interfaces/internal/service/svc_order"
)

func registerOrderRouter(group *gin.RouterGroup) {
	var svc svc_order.OrderService
	_ = dig.Invoke(func(s svc_order.OrderService) {
		svc = s
	})
	ctrl := ctrl_order.NewOrderCtrl(svc)
	router := group.Group("order")
	router.POST("create_red_rnv", ctrl.CreateRedEnvelopeOrder)
	router.GET("info", ctrl.Info)
}
