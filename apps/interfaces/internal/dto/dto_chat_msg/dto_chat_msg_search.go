package dto_chat_msg

import (
	"lark/pkg/proto/pb_enum"
)

type SearchMessageReq struct {
	ChatId    int64  `form:"chat_id" json:"chat_id" binding:"required,gt=0"`
	LastMsgId int64  `form:"last_msg_id" json:"last_msg_id" binding:"omitempty,gte=0"`
	Query     string `form:"query" json:"query" binding:"required,min=1,max=128"`
	Size      int32  `form:"size" json:"size" binding:"required,gte=10,lte=50"`
}

type SearchMessageResp struct {
	Total int64             `json:"total"`
	List  []*MessageSummary `json:"list"`
}

type MessageSummary struct {
	SrvMsgId       int64                 `json:"srv_msg_id"`      // 服务端消息号
	CliMsgId       int64                 `json:"cli_msg_id"`      // 客户端消息号
	SenderId       int64                 `json:"sender_id"`       // 发送者uid
	SenderPlatform pb_enum.PLATFORM_TYPE `json:"sender_platform"` // 发送者平台
	ChatId         int64                 `json:"chat_id"`         // 会话ID
	ChatType       pb_enum.CHAT_TYPE     `json:"chat_type"`       // 会话类型
	SeqId          int64                 `json:"seq_id"`          // 消息唯一ID
	MsgFrom        pb_enum.MSG_FROM      `json:"msg_from"`        // 消息来源
	MsgType        pb_enum.MSG_TYPE      `json:"msg_type"`        // 消息类型
	Rt             string                `json:"rt"`              // 高亮消息本体
	Status         int32                 `json:"status"`          // 消息状态
	SentTs         int64                 `json:"sent_ts"`         // 客户端本地发送时间
	SrvTs          int64                 `json:"srv_ts"`          // 服务端接收消息的时间
}
