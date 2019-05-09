package models

import "database/sql"

type RSmsrecord struct {
	Addtime      sql.NullInt64  `gorm:"column:addtime" json:"addtime"`
	Channel      sql.NullString `gorm:"column:channel" json:"channel"`
	Cid          int            `gorm:"column:cid" json:"cid"`
	Cmodel       sql.NullString `gorm:"column:cmodel" json:"cmodel"`
	Content      sql.NullString `gorm:"column:content" json:"content"`
	CustomerName sql.NullString `gorm:"column:customerName" json:"customerName"`
	Driver       sql.NullString `gorm:"column:driver" json:"driver"`
	Flag         sql.NullInt64  `gorm:"column:flag" json:"flag"`
	FlagContent  sql.NullString `gorm:"column:flag_content" json:"flag_content"`
	ID           int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Phone        string         `gorm:"column:phone" json:"phone"`
	Pid          sql.NullInt64  `gorm:"column:pid" json:"pid"`
	RegisterNo   sql.NullString `gorm:"column:registerNo" json:"registerNo"`
	SendTime     sql.NullInt64  `gorm:"column:send_time" json:"send_time"`
	SmsNum       sql.NullInt64  `gorm:"column:sms_num" json:"sms_num"`
	Status       sql.NullInt64  `gorm:"column:status" json:"status"`
	StoreID      sql.NullInt64  `gorm:"column:store_id" json:"store_id"`
	StoreName    sql.NullString `gorm:"column:store_name" json:"store_name"`
	Type         string         `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (r *RSmsrecord) TableName() string {
	return "r_smsrecord"
}
