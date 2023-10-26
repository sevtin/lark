package pdo

import "lark/pkg/utils"

var (
	field_tag_oauth_user_token string
)

type OauthUserToken struct {
	Openid      string `json:"openid" field:"openid"`
	Uid         int64  `json:"uid" field:"uid"`
	AccessToken string `json:"access_token" field:"access_token"`
}

func (p *OauthUserToken) GetFields() string {
	if field_tag_oauth_user_token == "" {
		field_tag_oauth_user_token = utils.GetFields(*p)
	}
	return field_tag_oauth_user_token
}
