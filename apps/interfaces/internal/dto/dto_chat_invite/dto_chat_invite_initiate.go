package dto_chat_invite

type InitiateChatInviteReq struct {
	ChatId   int64 `json:"chat_id" binding:"omitempty,gt=0"`       // chat ID
	ChatType int32 `json:"chat_type" binding:"required,oneof=1 2"` // 1:私聊/2:群聊
	//InitiatorUid  int64  `json:"initiator_uid" binding:"required,gt=0"`    // 发起人 UID
	InviteeUids   []int64 `json:"invitee_uids" binding:"required"`                 // 被邀请人 UID
	InvitationMsg string  `json:"invitation_msg" binding:"required,min=0,max=128"` // 邀请消息
}
