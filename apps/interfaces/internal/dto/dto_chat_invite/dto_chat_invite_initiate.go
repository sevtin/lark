package dto_chat_invite

type InitiateChatInviteReq struct {
	ChatId   int64 `json:"chat_id" validate:"omitempty,gt=0"`       // chat ID
	ChatType int32 `json:"chat_type" validate:"required,oneof=1 2"` // 1:私聊/2:群聊
	//InitiatorUid  int64  `json:"initiator_uid" validate:"required,gt=0"`    // 发起人 UID
	InviteeUids   []int64 `json:"invitee_uids" validate:"required"`                 // 被邀请人 UID
	InvitationMsg string  `json:"invitation_msg" validate:"required,min=0,max=128"` // 邀请消息
}
