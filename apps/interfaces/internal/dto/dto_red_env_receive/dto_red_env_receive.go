package dto_red_env_receive

type GrabRedEnvelopeReq struct {
	EnvId int64 `json:"env_id"` // 红包ID
	Uid   int64 `json:"uid"`    // 用户ID
}

type OpenRedEnvelopeReq struct {
	EnvId int64 `json:"env_id"` // 红包ID
	Uid   int64 `json:"uid"`    // 用户ID
}
