package models

import "database/sql"

type OrderCard struct {
	Address        string         `gorm:"column:address" json:"address"`
	AddressID      int            `gorm:"column:address_id" json:"address_id"`
	Carid          int            `gorm:"column:carid" json:"carid"`
	Cid            int            `gorm:"column:cid" json:"cid"`
	Delivery       sql.NullInt64  `gorm:"column:delivery" json:"delivery"`
	ID             int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Invoice        sql.NullString `gorm:"column:invoice" json:"invoice"`
	InvoiceContent sql.NullString `gorm:"column:invoice_content" json:"invoice_content"`
	Isdel          int            `gorm:"column:isdel" json:"isdel"`
	Name           string         `gorm:"column:name" json:"name"`
	Num            int            `gorm:"column:num" json:"num"`
	OrderTime      int            `gorm:"column:order_time" json:"order_time"`
	Orderid        string         `gorm:"column:orderid" json:"orderid"`
	PayFee         float32        `gorm:"column:pay_fee" json:"pay_fee"`
	PayTime        int            `gorm:"column:pay_time" json:"pay_time"`
	Paystate       int            `gorm:"column:paystate" json:"paystate"`
	Phone          string         `gorm:"column:phone" json:"phone"`
	PostExpress    sql.NullString `gorm:"column:post_express" json:"post_express"`
	PostNo         sql.NullString `gorm:"column:post_no" json:"post_no"`
	PostTime       int            `gorm:"column:post_time" json:"post_time"`
	Price          float32        `gorm:"column:price" json:"price"`
	Referid        int            `gorm:"column:referid" json:"referid"`
	Remark         sql.NullString `gorm:"column:remark" json:"remark"`
	ReserveTime    int            `gorm:"column:reserve_time" json:"reserve_time"`
	ServiceStoreid int            `gorm:"column:service_storeid" json:"service_storeid"`
	Status         int            `gorm:"column:status" json:"status"`
	StatusTxt      sql.NullString `gorm:"column:statusTxt" json:"statusTxt"`
	StoreID        int            `gorm:"column:store_id" json:"store_id"`
	UID            int            `gorm:"column:uid" json:"uid"`
	UpdateTime     int            `gorm:"column:update_time" json:"update_time"`
	Vehicle        string         `gorm:"column:vehicle" json:"vehicle"`
	VerifyTime     int            `gorm:"column:verify_time" json:"verify_time"`
}

// TableName sets the insert table name for this struct type
func (o *OrderCard) TableName() string {
	return "order_card"
}
