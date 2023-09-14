package ctrl_auth

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/pkg/common/xgin"
	"lark/pkg/common/xlog"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/xhttp"
)

func (ctrl *AuthCtrl) GoogleOAuth2Callback(ctx *gin.Context) {
	var (
		params = new(dto_auth.GoogleOauthCallbackReq)
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.ShouldBindQuery(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	params.Platform = pb_enum.PLATFORM_TYPE_WEB
	resp = ctrl.authService.GoogleOAuth2Callback(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
