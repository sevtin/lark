package xgin

import (
	"github.com/gin-gonic/gin"
	"lark/pkg/constant"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
)

func GetUid(ctx *gin.Context) (uid int64) {
	var (
		value  any
		exists bool
	)
	value, exists = ctx.Get(constant.USER_UID)
	if exists == false {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST, xhttp.ERROR_HTTP_USER_ID_DOESNOT_EXIST)
		return
	}
	uid, _ = utils.ToInt64(value)
	if uid == 0 {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST, xhttp.ERROR_HTTP_USER_ID_DOESNOT_EXIST)
		return
	}
	return
}
