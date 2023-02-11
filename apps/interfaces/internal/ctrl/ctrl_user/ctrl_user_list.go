package ctrl_user

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/dto/dto_user"
	"lark/pkg/common/xgin"
	"lark/pkg/common/xlog"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
)

func (ctrl *UserCtrl) UserList(ctx *gin.Context) {
	var (
		params = new(dto_user.UserListReq)
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.ShouldBindQuery(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	if err = utils.Struct(params); err != nil {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, err.Error())
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	resp = ctrl.userService.UserList(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
