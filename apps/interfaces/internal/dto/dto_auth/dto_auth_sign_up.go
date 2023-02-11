package dto_auth

import (
	"lark/pkg/proto/pb_enum"
)

/*
http://t.zoukankan.com/MyUniverse-p-15227003.html
https://www.h5w3.com/235615.html
http://www.zzvips.com/article/222569.html
*/
type SignUpReq struct {
	AccountType int32                 `json:"account_type" validate:"required,oneof=1 2"`       // 登录类型 1:手机号 2:lark账户
	RegPlatform pb_enum.PLATFORM_TYPE `json:"reg_platform" validate:"required,oneof=1 2 3 4 5"` // 注册平台 1:iOS 2:安卓
	Nickname    string                `json:"nickname" validate:"required,min=1,max=20"`        // 昵称
	Password    string                `json:"password" validate:"required,len=32"`              // 密码
	Firstname   string                `json:"firstname" validate:"required,min=1,max=20"`       // firstname
	Lastname    string                `json:"lastname" validate:"required,min=1,max=20"`        // lastname
	Gender      int32                 `json:"gender" validate:"omitempty,oneof=0 1 2"`          // 性别
	BirthTs     int64                 `json:"birth_ts" validate:"omitempty,gt=0"`               // 生日
	Email       string                `json:"email" validate:"omitempty,email"`                 // Email
	Mobile      string                `json:"mobile" validate:"required,min=8,max=20"`          // 手机号
	AvatarKey   string                `json:"avatar_key" validate:"omitempty"`                  // 头像(暂时弃用)
	CityId      int64                 `json:"city_id" validate:"omitempty,gte=0"`               // 城市ID
	Code        string                `json:"code" validate:"omitempty,min=4,max=6"`            // 城市ID
	Udid        string                `json:"udid" validate:"required,len=40"`                  // UDID
}
