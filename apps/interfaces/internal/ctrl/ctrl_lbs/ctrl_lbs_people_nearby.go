package ctrl_lbs

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/dto/dto_lbs"
	"lark/pkg/common/xgin"
	"lark/pkg/common/xlog"
	"lark/pkg/xhttp"
)

func (ctrl *LbsCtrl) PeopleNearby(ctx *gin.Context) {
	var (
		params = new(dto_lbs.PeopleNearbyReq)
		uid    int64
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.ShouldBindQuery(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	uid = xgin.GetUid(ctx)
	if uid == 0 {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST, xhttp.ERROR_HTTP_USER_ID_DOESNOT_EXIST)
		return
	}
	resp = ctrl.lbsService.PeopleNearby(params, uid)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
