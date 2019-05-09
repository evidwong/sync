package models

import (
	"database/sql"
	"time"
)

type LotteryPrize struct {
	ActivityID sql.NullInt64   `gorm:"column:activity_id" json:"activity_id"`
	AddTime    time.Time       `gorm:"column:add_time" json:"add_time"`
	Cid        sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ID         int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Image      sql.NullString  `gorm:"column:image" json:"image"`
	Isdel      sql.NullInt64   `gorm:"column:isdel" json:"isdel"`
	Level      sql.NullString  `gorm:"column:level" json:"level"`
	Name       string          `gorm:"column:name" json:"name"`
	Num        sql.NullInt64   `gorm:"column:num" json:"num"`
	Number     sql.NullInt64   `gorm:"column:number" json:"number"`
	Odds       sql.NullFloat64 `gorm:"column:odds" json:"odds"`
	Pnameid    sql.NullInt64   `gorm:"column:pnameid" json:"pnameid"`
	Price      sql.NullFloat64 `gorm:"column:price" json:"price"`
	Remain     sql.NullInt64   `gorm:"column:remain" json:"remain"`
	Thump      sql.NullString  `gorm:"column:thump" json:"thump"`
}

// TableName sets the insert table name for this struct type
func (l *LotteryPrize) TableName() string {
	return "lottery_prize"
}
