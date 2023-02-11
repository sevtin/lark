package xhttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Msg  string `json:"msg"`
	Code int32  `json:"code"`
}

type Resp struct {
	Result
	Data interface{}
}

func Success(ctx *gin.Context, data ...interface{}) {
	if len(data) == 0 || data[0] == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data[0],
	})
}

func Error(ctx *gin.Context, code int32, err string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  err,
	})
}

func (r *Resp) SetResult(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
