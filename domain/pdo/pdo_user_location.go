package pdo

import "lark/pkg/utils"

var (
	field_tag_user_location string
)

type UserLocation struct {
	Uid      int64  `json:"uid" field:"uid"`           // uid
	Nickname string `json:"nickname" field:"nickname"` // 昵称
	Gender   int32  `json:"gender" field:"gender"`     // 性别
	BirthTs  int64  `json:"birth_ts" field:"birth_ts"` // 生日
	Avatar   string `json:"avatar" field:"avatar"`     // 小图 72*72
}

func (p *UserLocation) GetFields() string {
	if field_tag_user_location == "" {
		field_tag_user_location = utils.GetFields(*p)
	}
	return field_tag_user_location
}
