package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_lbs"
	"lark/apps/interfaces/internal/service/svc_lbs"
)

func registerLbsRouter(group *gin.RouterGroup) {
	var svc svc_lbs.LbsService
	dig.Invoke(func(s svc_lbs.LbsService) {
		svc = s
	})
	ctrl := ctrl_lbs.NewLbsCtrl(svc)
	router := group.Group("lbs")
	router.POST("report_lng_lat", ctrl.ReportLngLat)
	router.GET("people_nearby", ctrl.PeopleNearby)
}
