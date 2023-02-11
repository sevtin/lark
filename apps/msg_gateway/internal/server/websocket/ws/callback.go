package ws

import "lark/pkg/proto/pb_msg"

type Message struct {
	Uid      int64 // 用户ID
	Platform int32 // 平台ID
	Hub      *Hub
	Packet   *pb_msg.Packet
}

type MessageCallback func(msg *Message)
