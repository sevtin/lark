package pdo

import "lark/pkg/utils"

var (
	field_tag_basic_user_info string
)

type BasicUserInfo struct {
	Uid      int64  `json:"uid" field:"uid"`           // uid
	LarkId   string `json:"lark_id" field:"lark_id"`   // 账户ID
	Nickname string `json:"nickname" field:"nickname"` // 昵称
	Gender   int32  `json:"gender" field:"gender"`     // 性别
	BirthTs  int64  `json:"birth_ts" field:"birth_ts"` // 生日
	CityId   int64  `json:"city_id" field:"city_id"`   // 城市id
	Avatar   string `json:"avatar" field:"avatar"`     // 头像
}

func (p *BasicUserInfo) GetFields() string {
	if field_tag_basic_user_info == "" {
		field_tag_basic_user_info = utils.GetFields(*p)
	}
	return field_tag_basic_user_info
}
