package models

import "database/sql"

type Reserve struct {
	Cid           int            `gorm:"column:cid" json:"cid"`
	FormatContent sql.NullString `gorm:"column:format_content" json:"format_content"`
	ID            int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Items         sql.NullString `gorm:"column:items" json:"items"`
	Message       string         `gorm:"column:message" json:"message"`
	Name          string         `gorm:"column:name" json:"name"`
	Phone         string         `gorm:"column:phone" json:"phone"`
	Rejob         string         `gorm:"column:rejob" json:"rejob"`
	Status        int            `gorm:"column:status" json:"status"`
	Storeid       int            `gorm:"column:storeid" json:"storeid"`
	Times         int            `gorm:"column:times" json:"times"`
	Type          int            `gorm:"column:type" json:"type"`
	UID           int            `gorm:"column:uid" json:"uid"`
	Updatetime    int            `gorm:"column:updatetime" json:"updatetime"`
	Vehicle       string         `gorm:"column:vehicle" json:"vehicle"`
	WechatNotify  int            `gorm:"column:wechat_notify" json:"wechat_notify"`
}

// TableName sets the insert table name for this struct type
func (r *Reserve) TableName() string {
	return "reserve"
}
