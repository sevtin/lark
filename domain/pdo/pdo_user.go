package pdo

import "lark/pkg/utils"

var (
	field_tag_user_info string
)

type UserInfo struct {
	Uid       int64  `json:"uid" field:"uid"`             // uid
	LarkId    string `json:"lark_id" field:"lark_id"`     // 账户ID
	Status    int32  `json:"status" field:"status"`       // 用户状态
	Nickname  string `json:"nickname" field:"nickname"`   // 昵称
	Firstname string `json:"firstname" field:"firstname"` // firstname
	Lastname  string `json:"lastname" field:"lastname"`   // lastname
	Gender    int32  `json:"gender" field:"gender"`       // 性别
	BirthTs   int64  `json:"birth_ts" field:"birth_ts"`   // 生日
	Mobile    string `json:"mobile" field:"mobile"`       // 手机号
	CityId    int64  `json:"city_id" field:"city_id"`     // 城市ID
}

func (p *UserInfo) GetField() string {
	if field_tag_user_info == "" {
		field_tag_user_info = utils.GetFields(*p)
	}
	return field_tag_user_info
}
