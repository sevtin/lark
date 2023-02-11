package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_chat_invite"
	"lark/apps/interfaces/internal/service/svc_chat_invite"
)

func registerChatInviteRouter(group *gin.RouterGroup) {
	var svc svc_chat_invite.ChatInviteService
	dig.Invoke(func(s svc_chat_invite.ChatInviteService) {
		svc = s
	})
	ctrl := ctrl_chat_invite.NewChatInviteCtrl(svc)
	router := group.Group("chat_invite")
	router.POST("initiate", ctrl.InitiateChatInvite)
	router.POST("handle", ctrl.ChatInviteHandle)
	router.GET("list", ctrl.ChatInviteList)
}
