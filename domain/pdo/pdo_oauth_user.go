package pdo

import (
	"lark/pkg/utils"
)

var (
	field_tag_oauth_user string
)

type OauthUser struct {
	Openid string `json:"openid" field:"openid"` // 第三方用户ID
	Uid    int64  `json:"uid" field:"uid"`       // lark uid
}

func (p *OauthUser) GetField() string {
	if field_tag_oauth_user == "" {
		field_tag_oauth_user = utils.GetFields(*p)
	}
	return field_tag_oauth_user
}
