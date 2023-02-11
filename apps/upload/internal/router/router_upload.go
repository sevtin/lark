package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/upload/dig"
	"lark/apps/upload/internal/ctrl"
	"lark/apps/upload/internal/service"
)

func registerUploadRouter(group *gin.RouterGroup) {
	var svc service.UploadService
	dig.Invoke(func(s service.UploadService) {
		svc = s
	})
	ctrl := ctrl.NewUploadCtrl(svc)
	router := group.Group("upload")
	router.POST("avatar", ctrl.UploadAvatar)
	router.GET("presigned", ctrl.Presigned)
}
