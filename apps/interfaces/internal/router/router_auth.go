package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_auth"
	"lark/apps/interfaces/internal/service/svc_auth"
)

func registerOpenAuthRouter(group *gin.RouterGroup) {
	var svc svc_auth.AuthService
	dig.Invoke(func(s svc_auth.AuthService) {
		svc = s
	})
	ctrl := ctrl_auth.NewAuthCtrl(svc)
	router := group.Group("auth")
	router.POST("sign_in", ctrl.SignIn)
	router.POST("sign_up", ctrl.SignUp)
	router.POST("refresh_token", ctrl.RefreshToken)
}

func registerPrivateAuthRouter(group *gin.RouterGroup) {
	var svc svc_auth.AuthService
	dig.Invoke(func(s svc_auth.AuthService) {
		svc = s
	})
	ctrl := ctrl_auth.NewAuthCtrl(svc)
	router := group.Group("auth")
	router.POST("sign_out", ctrl.SignOut)
}
