package dto_chat_invite

type ChatInviteHandleReq struct {
	InviteId int64 `json:"invite_id" binding:"required,gt=0"`
	//HandlerUid   int64  `json:"handler_uid" binding:"required,gt=0"`       // 处理人 UID
	HandleResult int32  `json:"handle_result" binding:"required,oneof=1 2"`  // 结果
	HandleMsg    string `json:"handle_msg" binding:"required,min=0,max=128"` // 处理消息
}
