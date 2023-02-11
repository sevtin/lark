package dto_chat

type CreateGroupChatReq struct {
	Name    string  `json:"name" validate:"required,min=1,max=20"`   // 标题
	About   string  `json:"about" validate:"omitempty,min=1,max=20"` // About
	UidList []int64 `json:"uid_list" validate:"omitempty"`           // 邀请人员uid列表
}
