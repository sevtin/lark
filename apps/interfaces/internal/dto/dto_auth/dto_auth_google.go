package dto_auth

import "lark/pkg/proto/pb_enum"

type GoogleOauthCallbackReq struct {
	Code     string                `form:"code" json:"code"`
	State    string                `form:"state" json:"state"`
	Platform pb_enum.PLATFORM_TYPE `form:"is_self" json:"platform"` // 平台
}
