package models

import "database/sql"

type WechatUsertag struct {
	Cid        int           `gorm:"column:cid" json:"cid"`
	ID         int           `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel      int           `gorm:"column:isdel" json:"isdel"`
	Labeltag   string        `gorm:"column:labeltag" json:"labeltag"`
	Status     int           `gorm:"column:status" json:"status"`
	Tagid      sql.NullInt64 `gorm:"column:tagid" json:"tagid"`
	Updatetime int           `gorm:"column:updatetime" json:"updatetime"`
	UserCount  sql.NullInt64 `gorm:"column:user_count" json:"user_count"`
}

// TableName sets the insert table name for this struct type
func (w *WechatUsertag) TableName() string {
	return "wechat_usertag"
}
