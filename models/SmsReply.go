package models

import "database/sql"

type SmsReply struct {
	AddTime     sql.NullInt64  `gorm:"column:add_time" json:"add_time"`
	Cid         sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Content     sql.NullString `gorm:"column:content" json:"content"`
	Flag        sql.NullInt64  `gorm:"column:flag" json:"flag"`
	FlagContent sql.NullString `gorm:"column:flag_content" json:"flag_content"`
	ID          int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel       sql.NullInt64  `gorm:"column:isdel" json:"isdel"`
	Phones      sql.NullString `gorm:"column:phones" json:"phones"`
	Status      sql.NullInt64  `gorm:"column:status" json:"status"`
	Tempid      sql.NullInt64  `gorm:"column:tempid" json:"tempid"`
	Title       sql.NullString `gorm:"column:title" json:"title"`
	UID         sql.NullInt64  `gorm:"column:uid" json:"uid"`
	UpdateTime  sql.NullInt64  `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (s *SmsReply) TableName() string {
	return "sms_reply"
}
