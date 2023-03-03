package dto_user

type SearchUserReq struct {
	Uid     int64  `form:"uid" json:"uid" binding:"required,gt=0"`
	LastUid int64  `form:"last_uid" json:"last_uid" binding:"omitempty,gte=0"`
	Query   string `form:"query" json:"query" binding:"required,min=1,max=128"`
	Size    int32  `form:"size" json:"size" binding:"required,gte=10,lte=50"`
}

type SearchUserReqResp struct {
	Total int64          `json:"total"`
	List  []*UserSummary `json:"list"`
}

type UserSummary struct {
	Uid       int64  `json:"uid"`        // uid
	LarkId    string `json:"lark_id"`    // 账户ID
	Status    int32  `json:"status"`     // 用户状态
	Nickname  string `json:"nickname"`   // 昵称
	AvatarKey string `json:"avatar_key"` // 头像
	Gender    int32  `json:"gender"`     // 性别
	BirthTs   int64  `json:"birth_ts"`   // 生日
	CityId    int64  `json:"city_id"`    // 城市ID
}
