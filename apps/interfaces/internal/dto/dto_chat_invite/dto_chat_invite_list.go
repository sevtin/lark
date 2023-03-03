package dto_chat_invite

type ChatInviteListReq struct {
	//Uid          int64 `form:"uid" json:"uid" binding:"required,gt=0"`
	Role         int32 `form:"role" json:"role" binding:"required,gt=0"` // 角色 1:发起人 2:审批人
	MaxInviteId  int64 `form:"max_invite_id" json:"max_invite_id" binding:"omitempty,gte=0"`
	HandleResult int32 `form:"handle_result" json:"handle_result" binding:"omitempty,oneof=0 1 2"` // 结果
	Limit        int32 `form:"limit" json:"limit" binding:"required,gte=10,lte=50"`
}
