package models

import "database/sql"

type WechatKeyword struct {
	Cid     int            `gorm:"column:cid" json:"cid"`
	ID      int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel   int            `gorm:"column:isdel" json:"isdel"`
	Keyword string         `gorm:"column:keyword" json:"keyword"`
	Model   string         `gorm:"column:model" json:"model"`
	Replyid sql.NullInt64  `gorm:"column:replyid" json:"replyid"`
	Status  int            `gorm:"column:status" json:"status"`
	Time    int            `gorm:"column:time" json:"time"`
	Title   sql.NullString `gorm:"column:title" json:"title"`
	Type    string         `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (w *WechatKeyword) TableName() string {
	return "wechat_keyword"
}
