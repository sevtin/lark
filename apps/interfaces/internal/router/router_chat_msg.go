package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_chat_msg"
	"lark/apps/interfaces/internal/service/svc_chat_msg"
)

func registerChatMessageRouter(group *gin.RouterGroup) {
	var svc svc_chat_msg.ChatMessageService
	dig.Invoke(func(s svc_chat_msg.ChatMessageService) {
		svc = s
	})
	ctrl := ctrl_chat_msg.NewChatMessageCtrl(svc)
	router := group.Group("chat_msg")
	router.GET("list", ctrl.GetChatMessageList)
	// 弃用
	router.GET("messages", ctrl.GetChatMessages)
	router.GET("search", ctrl.Search)
	router.POST("operation", ctrl.MessageOperation)
	router.POST("send_msg", ctrl.SendChatMessage)
}
