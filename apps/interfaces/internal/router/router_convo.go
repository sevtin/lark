package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_convo"
	"lark/apps/interfaces/internal/service/svc_convo"
)

func registerConvoRouter(group *gin.RouterGroup) {
	var svc svc_convo.ConvoService
	dig.Invoke(func(s svc_convo.ConvoService) {
		svc = s
	})
	ctrl := ctrl_convo.NewConvoCtrl(svc)
	router := group.Group("convo")
	router.POST("list", ctrl.ConvoList)
	router.GET("chat_seq_list", ctrl.ConvoChatSeqList)
}
