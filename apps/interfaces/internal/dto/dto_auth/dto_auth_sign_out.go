package dto_auth

type SignOutReq struct {
	Uid      int64 `json:"uid"`      // uid
	Platform int32 `json:"platform"` // 平台
}
