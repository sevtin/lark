package dto_lbs

type ReportLngLatReq struct {
	//Uid       int64   `form:"uid" json:"uid"`             // uid
	Longitude float64 `form:"longitude" json:"longitude"` // 经度
	Latitude  float64 `form:"latitude" json:"latitude"`   // 纬度
}

type PeopleNearbyReq struct {
	Uid       int64   `form:"uid" json:"uid"`             // uid
	LastUid   int32   `form:"last_uid" json:"last_uid"`   // 最后一个uid
	Longitude float64 `form:"longitude" json:"longitude"` // 经度
	Latitude  float64 `form:"latitude" json:"latitude"`   // 纬度
	Radius    int64   `form:"radius" json:"radius"`       // 半径
	Gender    int32   `form:"gender" json:"gender"`       // 性别
	MinAge    int32   `form:"min_age" json:"min_age"`     // 最小年龄
	MaxAge    int32   `form:"max_age" json:"max_age"`     // 最大年龄
	Limit     int32   `form:"limit" json:"limit"`         // 数量限制
	Skip      int32   `form:"skip" json:"skip"`
}
