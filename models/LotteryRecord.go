package models

import (
	"database/sql"
	"time"
)

type LotteryRecord struct {
	ActivityID    sql.NullInt64  `gorm:"column:activity_id" json:"activity_id"`
	ActivityItem  sql.NullString `gorm:"column:activity_item" json:"activity_item"`
	ActivityTitle sql.NullString `gorm:"column:activity_title" json:"activity_title"`
	AddTime       time.Time      `gorm:"column:add_time" json:"add_time"`
	BatchCode     sql.NullString `gorm:"column:batch_code" json:"batch_code"`
	Cid           sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Comments      string         `gorm:"column:comments" json:"comments"`
	EnrollTime    time.Time      `gorm:"column:enroll_time" json:"enroll_time"`
	ID            int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Nickname      sql.NullString `gorm:"column:nickname" json:"nickname"`
	Number        int            `gorm:"column:number" json:"number"`
	Openid        sql.NullString `gorm:"column:openid" json:"openid"`
	Phone         sql.NullString `gorm:"column:phone" json:"phone"`
	PrizeID       sql.NullInt64  `gorm:"column:prize_id" json:"prize_id"`
	Remark        sql.NullString `gorm:"column:remark" json:"remark"`
	SignTime      time.Time      `gorm:"column:sign_time" json:"sign_time"`
	Status        int            `gorm:"column:status" json:"status"`
	Type          int            `gorm:"column:type" json:"type"`
	UID           sql.NullInt64  `gorm:"column:uid" json:"uid"`
	WxName        sql.NullString `gorm:"column:wx_name" json:"wx_name"`
}

// TableName sets the insert table name for this struct type
func (l *LotteryRecord) TableName() string {
	return "lottery_record"
}
