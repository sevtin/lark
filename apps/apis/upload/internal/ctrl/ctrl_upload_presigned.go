package ctrl

import (
	"github.com/gin-gonic/gin"
	"lark/apps/apis/upload/internal/dto"
	"lark/pkg/common/xgin"
	"lark/pkg/common/xlog"
	"lark/pkg/xhttp"
)

func (ctrl *UploadCtrl) Presigned(ctx *gin.Context) {
	var (
		params = new(dto.PresignedReq)
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.ShouldBindQuery(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	resp = ctrl.svc.Presigned(ctx, params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
