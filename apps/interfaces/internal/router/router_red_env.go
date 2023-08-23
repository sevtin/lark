package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_red_env"
	"lark/apps/interfaces/internal/service/svc_red_env"
)

func registerRedEnvRouter(group *gin.RouterGroup) {
	var svc svc_red_env.RedEnvService
	dig.Invoke(func(s svc_red_env.RedEnvService) {
		svc = s
	})
	ctrl := ctrl_red_env.NewRedEnvCtrl(svc)
	router := group.Group("red_env")
	router.POST("give", ctrl.GiveRedEnvelope)
}
