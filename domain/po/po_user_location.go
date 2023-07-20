package po

import "lark/pkg/entity"

type UserLocation struct {
	entity.GormEntityTs
	Uid       int64   `gorm:"column:uid;primary_key" json:"uid"`                           // 用户ID
	Longitude float64 `gorm:"column:longitude;default:0.000000;NOT NULL" json:"longitude"` // 经度
	Latitude  float64 `gorm:"column:latitude;default:0.000000;NOT NULL" json:"latitude"`   // 纬度
	OnlineTs  int64   `gorm:"column:online_ts;default:0;NOT NULL" json:"online_ts"`        // 最后一次上线时间
}
