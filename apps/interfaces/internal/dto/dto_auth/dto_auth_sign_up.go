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
	RegPlatform pb_enum.PLATFORM_TYPE `json:"reg_platform" binding:"required,oneof=1 2 3 4 5"` // 注册平台 1:iOS 2:安卓
	Nickname    string                `json:"nickname" binding:"required,min=1,max=20"`        // 昵称
	Password    string                `json:"password" binding:"required,len=32"`              // 密码
	Firstname   string                `json:"firstname" binding:"required,min=1,max=20"`       // firstname
	Lastname    string                `json:"lastname" binding:"required,min=1,max=20"`        // lastname
	Gender      int32                 `json:"gender" binding:"omitempty,oneof=0 1 2"`          // 性别
	BirthTs     int64                 `json:"birth_ts" binding:"omitempty,gt=0"`               // 生日
	Email       string                `json:"email" binding:"omitempty,email"`                 // Email
	Mobile      string                `json:"mobile" binding:"required,min=8,max=20"`          // 手机号
	Avatar      string                `json:"avatar" binding:"omitempty"`                      // 头像(暂时弃用)
	CityId      int64                 `json:"city_id" binding:"omitempty,gte=0"`               // 城市ID
	Code        string                `json:"code" binding:"omitempty,min=4,max=6"`            // 验证码(暂时弃用)
	Udid        string                `json:"udid" binding:"required,len=40"`                  // UDID
}
