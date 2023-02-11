package dto_user

type UserInfoReq struct {
	IsSelf bool  `form:"is_self" json:"is_self"`
	Uid    int64 `form:"uid" json:"uid" validate:"omitempty,gt=0"`
}

type UserInfoResp struct {
	UserInfo *UserInfo `json:"user_info"`
}
