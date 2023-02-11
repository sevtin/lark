package ws

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func httpSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func httpError(ctx *gin.Context, code int32, err string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  err,
	})
}
