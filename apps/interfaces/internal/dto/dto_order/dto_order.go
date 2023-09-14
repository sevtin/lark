package dto_order

type OrderEditReq struct {
}

type OrderInfoReq struct {
}

type CreateRedEnvelopeOrderReq struct {
	Uid      int64 `json:"uid"`                         // 用户ID
	EnvId    int64 `json:"env_id" binding:"required"`   // 红包ID
	Amount   int64 `json:"amount" binding:"required"`   // 红包金额(分)
	Platform int32 `json:"platform"`                    // 平台
	PayType  int32 `json:"pay_type" binding:"required"` // PayType
}
