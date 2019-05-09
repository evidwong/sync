package models

import (
	"database/sql"
	"time"
)

type WebBanner struct {
	BannerImage sql.NullString `gorm:"column:banner_image" json:"banner_image"`
	BannerLink  sql.NullString `gorm:"column:banner_link" json:"banner_link"`
	BannerName  string         `gorm:"column:banner_name" json:"banner_name"`
	BannerSort  int            `gorm:"column:banner_sort" json:"banner_sort"`
	BannerThumb string         `gorm:"column:banner_thumb" json:"banner_thumb"`
	Cid         sql.NullInt64  `gorm:"column:cid" json:"cid"`
	CreateTime  time.Time      `gorm:"column:create_time" json:"create_time"`
	CreateUser  string         `gorm:"column:create_user" json:"create_user"`
	ID          int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isbanner    sql.NullInt64  `gorm:"column:isbanner" json:"isbanner"`
	UpdateTime  time.Time      `gorm:"column:update_time" json:"update_time"`
	UpdateUser  string         `gorm:"column:update_user" json:"update_user"`
}

// TableName sets the insert table name for this struct type
func (w *WebBanner) TableName() string {
	return "web_banner"
}
