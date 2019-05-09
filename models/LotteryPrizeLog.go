package models

import "time"

type LotteryPrizeLog struct {
	ActivityID   int       `gorm:"column:activity_id" json:"activity_id"`
	Brand        string    `gorm:"column:brand" json:"brand"`
	CardIds      string    `gorm:"column:card_ids" json:"card_ids"`
	Cid          int       `gorm:"column:cid" json:"cid"`
	City         string    `gorm:"column:city" json:"city"`
	Comno        string    `gorm:"column:comno" json:"comno"`
	County       string    `gorm:"column:county" json:"county"`
	CreateAt     time.Time `gorm:"column:create_at" json:"create_at"`
	Customercode string    `gorm:"column:customercode" json:"customercode"`
	ID           int       `gorm:"column:id;primary_key" json:"id;primary_key"`
	Nickname     string    `gorm:"column:nickname" json:"nickname"`
	Openid       string    `gorm:"column:openid" json:"openid"`
	Orderid      string    `gorm:"column:orderid" json:"orderid"`
	PayAt        time.Time `gorm:"column:pay_at" json:"pay_at"`
	Phone        string    `gorm:"column:phone" json:"phone"`
	Price        float64   `gorm:"column:price" json:"price"`
	Progress     int       `gorm:"column:progress" json:"progress"`
	Province     string    `gorm:"column:province" json:"province"`
	Registerno   string    `gorm:"column:registerno" json:"registerno"`
	Status       int       `gorm:"column:status" json:"status"`
	Type         int       `gorm:"column:type" json:"type"`
	UpdateAt     time.Time `gorm:"column:update_at" json:"update_at"`
	Vehiclecode  string    `gorm:"column:vehiclecode" json:"vehiclecode"`
}

// TableName sets the insert table name for this struct type
func (l *LotteryPrizeLog) TableName() string {
	return "lottery_prize_log"
}
