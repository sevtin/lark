package dto_convo

type ConvoChatSeqListReq struct {
	Uid     int64 `json:"uid" form:"uid"`
	LastCid int64 `json:"last_cid" form:"last_cid"`
	LastTs  int64 `json:"last_ts" form:"last_ts"`
	Limit   int32 `json:"limit" form:"limit"`
}
