package pdo

import "lark/pkg/utils"

type OauthUserInfo struct {
	Uid      int64  `json:"uid" field:"uid"`
	Openid   string `json:"openid" field:"openid"`
	Nickname string `json:"nickname" field:"nickname"`
	//AccessToken string `json:"access_token" field:"access_token"`
	Email      string `json:"email" field:"email"`
	AvatarUrl  string `json:"avatar_url" field:"avatar_url"`
	RegisterTs int64  `json:"register_ts" field:"register_ts"`
	Ex         string `json:"ex" field:"ex"`
	SignedIn   bool   `json:"signed_in"`
	InviteCode string `json:"invite_code"`
}

/*
SELECT u.uid,u.openid,u.nickname,u.access_token,u.avatar_url
FROM oauth_users u
*/

var (
	field_tag_oauth_user_info string
)

func (p *OauthUserInfo) GetFields() string {
	if field_tag_oauth_user_info == "" {
		field_tag_oauth_user_info = utils.GetFields(*p)
	}
	return field_tag_oauth_user_info
}
