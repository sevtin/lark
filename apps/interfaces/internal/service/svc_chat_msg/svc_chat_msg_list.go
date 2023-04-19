package svc_chat_msg

import (
	"github.com/jinzhu/copier"
	"lark/apps/interfaces/internal/dto/dto_chat_msg"
	"lark/pkg/common/xlog"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_chat_msg"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
	"strings"
)

func (s *chatMessageService) GetChatMessageList(req *dto_chat_msg.GetChatMessageListReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs  = new(pb_chat_msg.GetChatMessageListReq)
		reply    *pb_chat_msg.GetChatMessageListResp
		list     []string
		seqId    string
		index    int
		messages = new(dto_chat_msg.ChatMessages)
	)
	copier.Copy(reqArgs, req)

	list = strings.Split(req.SeqIds, ",")
	if len(list) > constant.CONST_MSG_PAGINATION_MAXIMUM_NUMBER_OF_ROWS {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_PAGINATION_LIMIT_EXCEEDED, xhttp.ERROR_HTTP_PAGINATION_LIMIT_EXCEEDED)
		return
	}
	reqArgs.SeqIds = make([]int64, len(list))
	for index, seqId = range list {
		reqArgs.SeqIds[index] = utils.StrToInt64(seqId)
	}

	reply = s.chatMessageClient.GetChatMessageList(reqArgs)
	if reply == nil {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if reply.Code > 0 {
		resp.SetResult(reply.Code, reply.Msg)
		xlog.Warn(reply.Code, reply.Msg)
		return
	}
	copier.Copy(&messages, reply.Msgs)
	resp.Data = messages
	return
}

// 弃用
//func (s *chatMessageService) GetChatMessages(req *dto_chat_msg.GetChatMessagesReq) (resp *xhttp.Resp) {
//	resp = new(xhttp.Resp)
//	var (
//		getChatMessagesReq  = new(pb_chat_msg.GetChatMessagesReq)
//		getChatMessagesResp *pb_chat_msg.GetChatMessagesResp
//		list                = make([]*dto_chat_msg.SrvChatMessage, 0)
//	)
//	copier.Copy(getChatMessagesReq, req)
//	getChatMessagesResp = s.chatMessageClient.GetChatMessages(getChatMessagesReq)
//	if getChatMessagesResp == nil {
//		resp.SetResult(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
//		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
//		return
//	}
//	if getChatMessagesResp.Code > 0 {
//		resp.SetResult(getChatMessagesResp.Code, getChatMessagesResp.Msg)
//		xlog.Warn(getChatMessagesResp.Code, getChatMessagesResp.Msg)
//		return
//	}
//	copier.Copy(&list, getChatMessagesResp.List)
//	resp.Data = list
//	return
//}
