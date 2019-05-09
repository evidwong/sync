package models

import "database/sql"

type Stockrecord struct {
	Cid          int            `gorm:"column:cid" json:"cid"`
	Comno        sql.NullString `gorm:"column:comno" json:"comno"`
	CompanyName  sql.NullString `gorm:"column:company_name" json:"company_name"`
	CouponCardID string         `gorm:"column:coupon_card_id" json:"coupon_card_id"`
	GoodName     string         `gorm:"column:good_name" json:"good_name"`
	GoodsID      string         `gorm:"column:goods_id" json:"goods_id"`
	ID           int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	IsSms        sql.NullInt64  `gorm:"column:is_sms" json:"is_sms"`
	Isdel        int            `gorm:"column:isdel" json:"isdel"`
	Num          sql.NullString `gorm:"column:num" json:"num"`
	Status       int            `gorm:"column:status" json:"status"`
	StockID      int            `gorm:"column:stock_id" json:"stock_id"`
	StockName    string         `gorm:"column:stock_name" json:"stock_name"`
	StockTime    int            `gorm:"column:stock_time" json:"stock_time"`
	StoreName    sql.NullString `gorm:"column:store_name" json:"store_name"`
	Type         int            `gorm:"column:type" json:"type"`
	Uids         string         `gorm:"column:uids" json:"uids"`
	UserCode     sql.NullString `gorm:"column:user_code" json:"user_code"`
}

// TableName sets the insert table name for this struct type
func (s *Stockrecord) TableName() string {
	return "stockrecord"
}
