package middleware

import (
	"github.com/gin-gonic/gin"
	"lark/pkg/common/xredis"
	"lark/pkg/constant"
)

func Statistics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.FullPath()
		xredis.ZIncrBy(constant.RK_SYNC_API_REQUESTS, 1, path)
		ctx.Next()
	}
}
