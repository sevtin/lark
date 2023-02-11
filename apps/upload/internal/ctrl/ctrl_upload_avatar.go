package ctrl

import (
	"github.com/gin-gonic/gin"
	"lark/apps/upload/internal/dto"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
)

func (ctrl *UploadCtrl) UploadAvatar(ctx *gin.Context) {
	var (
		params   = new(dto.UploadAvatarReq)
		resp     *xhttp.Resp
		keyValue any
		exists   bool
		uid      int64
		err      error
	)
	if err = ctx.Bind(params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	switch pb_enum.AVATAR_OWNER(params.OwnerType) {
	case pb_enum.AVATAR_OWNER_USER_AVATAR:
		keyValue, exists = ctx.Get(constant.USER_UID)
		if exists == false {
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
			xlog.Warn(xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
			return
		}
		uid, _ = utils.ToInt64(keyValue)
		if uid == 0 {
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
			xlog.Warn(xhttp.ERROR_CODE_HTTP_GET_USER_INFO_FAILED, xhttp.ERROR_HTTP_GET_USER_INFO_FAILED)
			return
		}
		params.OwnerId = uid
	case pb_enum.AVATAR_OWNER_CHAT_AVATAR:
	default:
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR)
		xlog.Warn(ctx, xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, params.OwnerId)
		return
	}
	resp = ctrl.svc.UploadAvatar(ctx, params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
