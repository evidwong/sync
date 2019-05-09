package models

import (
	"database/sql"
	"time"
)

type LotteryType struct {
	ActivityType sql.NullString `gorm:"column:activity_type" json:"activity_type"`
	AddTime      time.Time      `gorm:"column:add_time" json:"add_time"`
	Cid          sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Comno        sql.NullString `gorm:"column:comno" json:"comno"`
	ID           int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Indexid      sql.NullInt64  `gorm:"column:indexid" json:"indexid"`
	Storeid      sql.NullInt64  `gorm:"column:storeid" json:"storeid"`
	UpdateTime   time.Time      `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (l *LotteryType) TableName() string {
	return "lottery_type"
}
