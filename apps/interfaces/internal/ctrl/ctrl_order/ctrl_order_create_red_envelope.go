package ctrl_order

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/dto/dto_order"
	"lark/pkg/common/xgin"
	"lark/pkg/common/xlog"
	"lark/pkg/xhttp"
)

func (ctrl *OrderCtrl) CreateRedEnvelopeOrder(ctx *gin.Context) {
	var (
		params = new(dto_order.CreateRedEnvelopeOrderReq)
		resp   *xhttp.Resp
		uid    int64
		err    error
	)
	if err = xgin.BindJSON(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	uid = xgin.GetUid(ctx)
	if uid == 0 {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST, xhttp.ERROR_HTTP_USER_ID_DOESNOT_EXIST)
		return
	}
	params.Platform = xgin.GetPlatform(ctx)
	resp = ctrl.orderService.CreateRedEnvelopeOrder(params, uid)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
