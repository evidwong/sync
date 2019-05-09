package models

import (
	"database/sql"
	"time"
)

type LotteryActivity struct {
	ActiveObject       sql.NullString  `gorm:"column:active_object" json:"active_object"`
	ActivityBalanceNum int             `gorm:"column:activity_balance_num" json:"activity_balance_num"`
	ActivityNum        sql.NullInt64   `gorm:"column:activity_num" json:"activity_num"`
	ActivityTime       time.Time       `gorm:"column:activity_time" json:"activity_time"`
	ActivityType       sql.NullString  `gorm:"column:activity_type" json:"activity_type"`
	ActivityUseNum     int             `gorm:"column:activity_use_num" json:"activity_use_num"`
	AddTime            time.Time       `gorm:"column:add_time" json:"add_time"`
	Address            sql.NullString  `gorm:"column:address" json:"address"`
	BrandIds           string          `gorm:"column:brand_ids" json:"brand_ids"`
	CardIds            string          `gorm:"column:card_ids" json:"card_ids"`
	Cid                sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Code               sql.NullString  `gorm:"column:code" json:"code"`
	Comno              sql.NullString  `gorm:"column:comno" json:"comno"`
	ComnoName          sql.NullString  `gorm:"column:comno_name" json:"comno_name"`
	Description        sql.NullString  `gorm:"column:description" json:"description"`
	EachLotteryNum     sql.NullInt64   `gorm:"column:each_lottery_num" json:"each_lottery_num"`
	EndTime            time.Time       `gorm:"column:end_time" json:"end_time"`
	GetType            sql.NullString  `gorm:"column:get_type" json:"get_type"`
	GoodID             int             `gorm:"column:good_id" json:"good_id"`
	GoodsID            int             `gorm:"column:goods_id" json:"goods_id"`
	ID                 int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Image              sql.NullString  `gorm:"column:image" json:"image"`
	Intro              sql.NullString  `gorm:"column:intro" json:"intro"`
	IsArea             int             `gorm:"column:is_area" json:"is_area"`
	IsBrand            int             `gorm:"column:is_brand" json:"is_brand"`
	IsFees             int             `gorm:"column:is_fees" json:"is_fees"`
	IsLotteryDay       sql.NullInt64   `gorm:"column:is_lottery_day" json:"is_lottery_day"`
	IsPay              int             `gorm:"column:is_pay" json:"is_pay"`
	IsShare            int             `gorm:"column:is_share" json:"is_share"`
	IsWinLottery       sql.NullInt64   `gorm:"column:is_win_lottery" json:"is_win_lottery"`
	Isdel              sql.NullInt64   `gorm:"column:isdel" json:"isdel"`
	Isdraw             sql.NullInt64   `gorm:"column:isdraw" json:"isdraw"`
	Issign             sql.NullInt64   `gorm:"column:issign" json:"issign"`
	Istotal            int             `gorm:"column:istotal" json:"istotal"`
	Latitude           sql.NullFloat64 `gorm:"column:latitude" json:"latitude"`
	Longitude          sql.NullFloat64 `gorm:"column:longitude" json:"longitude"`
	LotteryNum         sql.NullInt64   `gorm:"column:lottery_num" json:"lottery_num"`
	PayPrice           float64         `gorm:"column:pay_price" json:"pay_price"`
	ReceiveEtime       time.Time       `gorm:"column:receive_etime" json:"receive_etime"`
	ReceiveStime       time.Time       `gorm:"column:receive_stime" json:"receive_stime"`
	StartTime          time.Time       `gorm:"column:start_time" json:"start_time"`
	StoreID            int             `gorm:"column:store_id" json:"store_id"`
	Storeid            sql.NullInt64   `gorm:"column:storeid" json:"storeid"`
	Thumb              sql.NullString  `gorm:"column:thumb" json:"thumb"`
	Title              sql.NullString  `gorm:"column:title" json:"title"`
	Tweet              sql.NullInt64   `gorm:"column:tweet" json:"tweet"`
	UserType           sql.NullInt64   `gorm:"column:user_type" json:"user_type"`
}

// TableName sets the insert table name for this struct type
func (l *LotteryActivity) TableName() string {
	return "lottery_activity"
}
