package ctrl_chat_msg

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/dto/dto_chat_msg"
	"lark/pkg/common/xgin"
	"lark/pkg/common/xlog"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
)

func (ctrl *ChatMessageCtrl) GetChatMessageList(ctx *gin.Context) {
	var (
		params = new(dto_chat_msg.GetChatMessageListReq)
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.ShouldBindQuery(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	resp = ctrl.chatMessageService.GetChatMessageList(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}

// 弃用
func (ctrl *ChatMessageCtrl) GetChatMessages(ctx *gin.Context) {
	var (
		params = new(dto_chat_msg.GetChatMessagesReq)
		resp   *xhttp.Resp
		err    error
	)
	if err = ctx.ShouldBindQuery(params); err != nil {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_DESERIALIZE_FAILED, xhttp.ERROR_HTTP_REQ_DESERIALIZE_FAILED, err.Error())
		return
	}
	if err = utils.Struct(params); err != nil {
		xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, err.Error())
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	resp = ctrl.chatMessageService.GetChatMessages(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
