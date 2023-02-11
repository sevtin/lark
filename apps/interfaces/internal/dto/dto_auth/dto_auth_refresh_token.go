package dto_auth

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"  validate:"required"` // 刷新token
}

//type RefreshTokenResp struct {
//	AccessToken string `json:"access_token"` // 用户token
//	Expire      int64  `json:"expire"`       // token过期时间戳（秒）
//}
