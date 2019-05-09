package models

import (
	"database/sql"
	"time"
)

type TelRecord struct {
	AddTime      sql.NullInt64  `gorm:"column:add_time" json:"add_time"`
	Cid          sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Cmodel       sql.NullString `gorm:"column:cmodel" json:"cmodel"`
	CustomerName sql.NullString `gorm:"column:customerName" json:"customerName"`
	Driver       sql.NullString `gorm:"column:driver" json:"driver"`
	Flag         sql.NullInt64  `gorm:"column:flag" json:"flag"`
	FlagContent  sql.NullString `gorm:"column:flag_content" json:"flag_content"`
	ID           int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel        sql.NullInt64  `gorm:"column:isdel" json:"isdel"`
	Phone        sql.NullString `gorm:"column:phone" json:"phone"`
	RecordTime   sql.NullInt64  `gorm:"column:recordTime" json:"recordTime"`
	RecordPath   sql.NullString `gorm:"column:record_path" json:"record_path"`
	RegisterNo   sql.NullString `gorm:"column:registerNo" json:"registerNo"`
	Result       sql.NullString `gorm:"column:result" json:"result"`
	StartTime    time.Time      `gorm:"column:start_time" json:"start_time"`
	Type         sql.NullString `gorm:"column:type" json:"type"`
	UID          sql.NullInt64  `gorm:"column:uid" json:"uid"`
	UpdateTime   sql.NullInt64  `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (t *TelRecord) TableName() string {
	return "tel_record"
}
