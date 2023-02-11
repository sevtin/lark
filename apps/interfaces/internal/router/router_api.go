package router

import (
	"github.com/gin-gonic/gin"
	"lark/pkg/middleware"
)

func Register(engine *gin.Engine) {
	//engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	publicGroup := engine.Group("open")
	registerPublicRoutes(publicGroup)

	privateGroup := engine.Group("api")
	registerPrivateRouter(privateGroup)
}

// 无需验证
func registerPublicRoutes(group *gin.RouterGroup) {
	registerOpenAuthRouter(group)
}

// 需要验证
func registerPrivateRouter(group *gin.RouterGroup) {
	group.Use(middleware.JwtAuth())
	registerPublicRouter(group)
	registerPrivateAuthRouter(group)
	registerUserRouter(group)
	registerChatMessageRouter(group)
	registerChatMemberRouter(group)
	registerChatInviteRouter(group)
	registerChatRouter(group)
	registerConvoRouter(group)
}

// 开放式api无需验证会话id
func registerPublicRouter(group *gin.RouterGroup) {
	router := group.Group("public")
	registerPublicTestRouter(router)
}

func registerPublicTestRouter(group *gin.RouterGroup) {

}
