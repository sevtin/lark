package router

import (
	"github.com/gin-gonic/gin"
	"lark/pkg/middleware"
)

func Register(engine *gin.Engine) {
	group := engine.Group("api")
	group.Use(middleware.JwtAuth())
	registerUploadRouter(group)
}
