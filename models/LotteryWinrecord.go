package models

import (
	"database/sql"
	"time"
)

type LotteryWinrecord struct {
	ActivityID    sql.NullInt64  `gorm:"column:activity_id" json:"activity_id"`
	ActivityItem  sql.NullString `gorm:"column:activity_item" json:"activity_item"`
	ActivityTitle sql.NullString `gorm:"column:activity_title" json:"activity_title"`
	ActivityType  sql.NullString `gorm:"column:activity_type" json:"activity_type"`
	AddTime       time.Time      `gorm:"column:add_time" json:"add_time"`
	Address       sql.NullString `gorm:"column:address" json:"address"`
	BjTime        time.Time      `gorm:"column:bj_time" json:"bj_time"`
	Cid           sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Comno         string         `gorm:"column:comno" json:"comno"`
	ID            int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	IsSure        sql.NullInt64  `gorm:"column:is_sure" json:"is_sure"`
	MobilePhone   sql.NullString `gorm:"column:mobile_phone" json:"mobile_phone"`
	Nickname      sql.NullString `gorm:"column:nickname" json:"nickname"`
	Openid        sql.NullString `gorm:"column:openid" json:"openid"`
	PrizeID       sql.NullInt64  `gorm:"column:prize_id" json:"prize_id"`
	PrizeItem     sql.NullString `gorm:"column:prize_item" json:"prize_item"`
	PrizeLevel    sql.NullString `gorm:"column:prize_level" json:"prize_level"`
	PrizeTitle    sql.NullString `gorm:"column:prize_title" json:"prize_title"`
	Remark        sql.NullString `gorm:"column:remark" json:"remark"`
	UID           sql.NullInt64  `gorm:"column:uid" json:"uid"`
	UserName      sql.NullString `gorm:"column:user_name" json:"user_name"`
}

// TableName sets the insert table name for this struct type
func (l *LotteryWinrecord) TableName() string {
	return "lottery_winrecord"
}
