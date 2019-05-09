package models

import "database/sql"

type RemindMessage struct {
	Addtime      sql.NullInt64  `gorm:"column:addtime" json:"addtime"`
	Carid        sql.NullInt64  `gorm:"column:carid" json:"carid"`
	Cid          sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Content      sql.NullString `gorm:"column:content" json:"content"`
	HandPhone    sql.NullString `gorm:"column:handPhone" json:"handPhone"`
	ID           int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isread       sql.NullInt64  `gorm:"column:isread" json:"isread"`
	LastTime     sql.NullInt64  `gorm:"column:last_time" json:"last_time"`
	Orderid      sql.NullString `gorm:"column:orderid" json:"orderid"`
	Status       sql.NullString `gorm:"column:status" json:"status"`
	Title        sql.NullString `gorm:"column:title" json:"title"`
	Type         sql.NullString `gorm:"column:type" json:"type"`
	UID          sql.NullInt64  `gorm:"column:uid" json:"uid"`
	WebappNotify sql.NullInt64  `gorm:"column:webapp_notify" json:"webapp_notify"`
	WechatNotify sql.NullInt64  `gorm:"column:wechat_notify" json:"wechat_notify"`
}

// TableName sets the insert table name for this struct type
func (r *RemindMessage) TableName() string {
	return "remind_message"
}
