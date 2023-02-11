package dto_chat_invite

type ChatInviteHandleReq struct {
	InviteId int64 `json:"invite_id" validate:"required,gt=0"`
	//HandlerUid   int64  `json:"handler_uid" validate:"required,gt=0"`       // 处理人 UID
	HandleResult int32  `json:"handle_result" validate:"required,oneof=1 2"`  // 结果
	HandleMsg    string `json:"handle_msg" validate:"required,min=0,max=128"` // 处理消息
}
