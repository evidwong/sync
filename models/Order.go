package models

import "database/sql"

type Order struct {
	ActivityID      int             `gorm:"column:activity_id" json:"activity_id"`
	Address         sql.NullString  `gorm:"column:address" json:"address"`
	AddressID       sql.NullInt64   `gorm:"column:address_id" json:"address_id"`
	CardFee         sql.NullFloat64 `gorm:"column:card_fee" json:"card_fee"`
	Carid           sql.NullInt64   `gorm:"column:carid" json:"carid"`
	Cid             sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ComeFrom        sql.NullString  `gorm:"column:comeFrom" json:"comeFrom"`
	Delivery        sql.NullInt64   `gorm:"column:delivery" json:"delivery"`
	EvaluateContent sql.NullString  `gorm:"column:evaluate_content" json:"evaluate_content"`
	ID              int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Invoice         sql.NullString  `gorm:"column:invoice" json:"invoice"`
	InvoiceContent  sql.NullString  `gorm:"column:invoice_content" json:"invoice_content"`
	InvoiceNo       sql.NullString  `gorm:"column:invoice_no" json:"invoice_no"`
	IsEvaluate      int             `gorm:"column:is_evaluate" json:"is_evaluate"`
	Isdel           sql.NullInt64   `gorm:"column:isdel" json:"isdel"`
	MemberCard      sql.NullString  `gorm:"column:member_card" json:"member_card"`
	Name            sql.NullString  `gorm:"column:name" json:"name"`
	Nickname        sql.NullString  `gorm:"column:nickname" json:"nickname"`
	Num             sql.NullInt64   `gorm:"column:num" json:"num"`
	OilCard         sql.NullString  `gorm:"column:oil_card" json:"oil_card"`
	OilChannel      sql.NullString  `gorm:"column:oil_channel" json:"oil_channel"`
	OilRechargeInfo sql.NullString  `gorm:"column:oil_recharge_info" json:"oil_recharge_info"`
	OrderPics       sql.NullString  `gorm:"column:order_pics" json:"order_pics"`
	OrderRemark     sql.NullString  `gorm:"column:order_remark" json:"order_remark"`
	OrderTime       sql.NullInt64   `gorm:"column:order_time" json:"order_time"`
	Orderid         sql.NullString  `gorm:"column:orderid" json:"orderid"`
	PayFee          sql.NullFloat64 `gorm:"column:pay_fee" json:"pay_fee"`
	PayTime         sql.NullInt64   `gorm:"column:pay_time" json:"pay_time"`
	Paystate        sql.NullInt64   `gorm:"column:paystate" json:"paystate"`
	Phone           sql.NullString  `gorm:"column:phone" json:"phone"`
	PostExpress     sql.NullString  `gorm:"column:post_express" json:"post_express"`
	PostNo          sql.NullString  `gorm:"column:post_no" json:"post_no"`
	PostTime        sql.NullInt64   `gorm:"column:post_time" json:"post_time"`
	Price           sql.NullFloat64 `gorm:"column:price" json:"price"`
	RechargeStatus  int             `gorm:"column:recharge_status" json:"recharge_status"`
	Referid         sql.NullInt64   `gorm:"column:referid" json:"referid"`
	Remark          sql.NullString  `gorm:"column:remark" json:"remark"`
	ReserveTime     sql.NullInt64   `gorm:"column:reserve_time" json:"reserve_time"`
	Score           int             `gorm:"column:score" json:"score"`
	ServiceStoreid  sql.NullInt64   `gorm:"column:service_storeid" json:"service_storeid"`
	Status          sql.NullInt64   `gorm:"column:status" json:"status"`
	StatusTxt       sql.NullString  `gorm:"column:statusTxt" json:"statusTxt"`
	StoreID         sql.NullInt64   `gorm:"column:store_id" json:"store_id"`
	UID             sql.NullInt64   `gorm:"column:uid" json:"uid"`
	UpdateTime      sql.NullInt64   `gorm:"column:update_time" json:"update_time"`
	Vehicle         sql.NullString  `gorm:"column:vehicle" json:"vehicle"`
	VerifyTime      sql.NullInt64   `gorm:"column:verify_time" json:"verify_time"`
}

// TableName sets the insert table name for this struct type
func (o *Order) TableName() string {
	return "order"
}
