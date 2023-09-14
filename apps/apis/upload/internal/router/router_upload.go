package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/apis/upload/dig"
	"lark/apps/apis/upload/internal/ctrl"
	"lark/apps/apis/upload/internal/service"
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
