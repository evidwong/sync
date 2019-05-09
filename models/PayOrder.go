package models

import (
	"database/sql"
	"time"
)

type PayOrder struct {
	AddTime       time.Time       `gorm:"column:add_time" json:"add_time"`
	Amount        sql.NullFloat64 `gorm:"column:amount" json:"amount"`
	Bankid        sql.NullString  `gorm:"column:bankid" json:"bankid"`
	Bankname      sql.NullString  `gorm:"column:bankname" json:"bankname"`
	Cid           sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Exist         sql.NullInt64   `gorm:"column:exist" json:"exist"`
	Funcode       sql.NullString  `gorm:"column:funcode" json:"funcode"`
	ID            int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	ItemJSON      sql.NullString  `gorm:"column:item_json" json:"item_json"`
	Jsno          sql.NullString  `gorm:"column:jsno" json:"jsno"`
	OrderFrom     sql.NullInt64   `gorm:"column:orderFrom" json:"orderFrom"`
	Orderid       sql.NullString  `gorm:"column:orderid" json:"orderid"`
	PayAmount     sql.NullFloat64 `gorm:"column:pay_amount" json:"pay_amount"`
	PayForCid     sql.NullInt64   `gorm:"column:pay_for_cid" json:"pay_for_cid"`
	PayTime       time.Time       `gorm:"column:pay_time" json:"pay_time"`
	PayType       sql.NullInt64   `gorm:"column:pay_type" json:"pay_type"`
	Paystate      sql.NullInt64   `gorm:"column:paystate" json:"paystate"`
	Phone         sql.NullString  `gorm:"column:phone" json:"phone"`
	ReceptionCode sql.NullString  `gorm:"column:receptionCode" json:"receptionCode"`
	RegisterNo    sql.NullString  `gorm:"column:registerNo" json:"registerNo"`
	Remark        sql.NullString  `gorm:"column:remark" json:"remark"`
	Storeid       sql.NullString  `gorm:"column:storeid" json:"storeid"`
	Type          sql.NullString  `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (p *PayOrder) TableName() string {
	return "pay_order"
}
