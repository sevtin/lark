package dto_auth

import (
	"lark/pkg/proto/pb_enum"
)

type GithubOauthCallbackReq struct {
	Code     string                `json:"code"`
	Platform pb_enum.PLATFORM_TYPE `json:"platform"` // 平台
}
