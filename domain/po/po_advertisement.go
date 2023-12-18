package po

import "lark/pkg/entity"

type Advertisement struct {
	entity.GormEntityTs
	Id         int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`   // ID
	AdType     int    `gorm:"column:ad_type;default:0;NOT NULL" json:"ad_type"` // 广告类型
	AdTitle    string `gorm:"column:ad_title;NOT NULL" json:"ad_title"`         // 广告标题
	AdSubTitle string `gorm:"column:ad_sub_title;NOT NULL" json:"ad_sub_title"` // 广告标题
	AdImage    string `gorm:"column:ad_image;NOT NULL" json:"ad_image"`         // 广告图片
	AdContent  string `gorm:"column:ad_content;NOT NULL" json:"ad_content"`     // 广告json
}
