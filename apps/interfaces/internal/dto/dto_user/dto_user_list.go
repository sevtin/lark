package dto_user

import (
	"lark/apps/interfaces/internal/dto/dto_avatar"
)

type UserListReq struct {
	Uids string `form:"uids" json:"uids" validate:"required"` // Uid List
}

type UserListResp struct {
	List []UserInfo `json:"list"` // user 列表
}

type UserInfo struct {
	Uid       int64                  `json:"uid"`       // uid
	LarkId    string                 `json:"lark_id"`   // 账户ID
	Status    int32                  `json:"status"`    // 用户状态
	Nickname  string                 `json:"nickname"`  // 昵称
	Firstname string                 `json:"firstname"` // firstname
	Lastname  string                 `json:"lastname"`  // lastname
	Gender    int32                  `json:"gender"`    // 性别
	BirthTs   int64                  `json:"birth_ts"`  // 生日
	CityId    int64                  `json:"city_id"`   // 城市ID
	Avatar    *dto_avatar.AvatarInfo `json:"avatar"`    // 头像
}
