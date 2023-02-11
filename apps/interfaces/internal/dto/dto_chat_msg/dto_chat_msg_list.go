package dto_chat_msg

import (
	"lark/pkg/proto/pb_enum"
)

type GetChatMessageListReq struct {
	ChatId int64  `form:"chat_id" json:"chat_id" validate:"required,gt=0"`
	SeqIds string `form:"seq_ids" json:"seq_ids" validate:"required"`
	MsgTs  int64  `form:"msg_ts" json:"msg_ts" validate:"omitempty,gte=0"`
	Order  int32  `form:"order" json:"order" validate:"omitempty,gte=0,lte=1"`
}

type ChatMessages struct {
	LastSeqId int64             `json:"last_seq_id"`
	List      []*SrvChatMessage `json:"list"`
}

type GetChatMessagesReq struct {
	ChatType int32 `form:"chat_type" json:"chat_type" validate:"required,gte=0,lte=127"`
	ChatId   int64 `form:"chat_id" json:"chat_id" validate:"required,gt=0"`
	SeqId    int64 `form:"seq_id" json:"seq_id" validate:"omitempty,gte=0"`
	Limit    int32 `form:"limit" json:"limit" validate:"required,gte=10,lte=50"`
	New      bool  `form:"new" json:"new"`
	MsgTs    int64 `form:"msg_ts" json:"msg_ts" validate:"omitempty,gte=0"`
}

type SrvChatMessage struct {
	SrvMsgId       int64                 `json:"srv_msg_id"`      // 服务端消息号
	CliMsgId       int64                 `json:"cli_msg_id"`      // 客户端消息号
	SenderId       int64                 `json:"sender_id"`       // 发送者uid
	SenderPlatform pb_enum.PLATFORM_TYPE `json:"sender_platform"` // 发送者平台
	ChatId         int64                 `json:"chat_id"`         // 会话ID
	ChatType       pb_enum.CHAT_TYPE     `json:"chat_type"`       // 会话类型
	SeqId          int64                 `json:"seq_id"`          // 消息唯一ID
	MsgFrom        pb_enum.MSG_FROM      `json:"msg_from"`        // 消息来源
	MsgType        pb_enum.MSG_TYPE      `json:"msg_type"`        // 消息类型
	Body           string                `json:"body"`            // 消息本体
	Status         int32                 `json:"status"`          // 消息状态
	SentTs         int64                 `json:"sent_ts"`         // 客户端本地发送时间
	SrvTs          int64                 `json:"srv_ts"`          // 服务端接收消息的时间
}
