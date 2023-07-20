package mpo

const (
	MONGO_COLLECTION_USER_LOCATIONS = "user_locations"
)

type UserLocation struct {
	Uid      int64     `bson:"uid" json:"uid"`             // UID
	Gender   int32     `bson:"gender" json:"gender"`       // 性别
	BirthTs  int64     `bson:"birth_ts" json:"birth_ts"`   // 生日
	Nickname string    `bson:"nickname" json:"nickname"`   // 昵称
	Avatar   string    `bson:"avatar" json:"avatar"`       // 小图 72*72
	OnlineTs int64     `bson:"online_ts" json:"online_ts"` // 上线时间
	Location *Location `bson:"location" json:"location"`   // 经纬度
	Distance float64   `bson:"-" json:"distance"`          // 距离
}

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"` // 经纬度
}
