package models

import "database/sql"

type WebDisOrder struct {
	Amount       sql.NullFloat64 `gorm:"column:amount" json:"amount"`
	Cid          sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ClientName   sql.NullString  `gorm:"column:client_name" json:"client_name"`
	Comno        sql.NullString  `gorm:"column:comno" json:"comno"`
	CouponID     sql.NullString  `gorm:"column:coupon_id" json:"coupon_id"`
	CustomerName sql.NullString  `gorm:"column:customer_name" json:"customer_name"`
	HandPhone    sql.NullString  `gorm:"column:hand_phone" json:"hand_phone"`
	ID           int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	OptUID       sql.NullInt64   `gorm:"column:opt_uid" json:"opt_uid"`
	OptUsername  sql.NullString  `gorm:"column:opt_username" json:"opt_username"`
	OrderTime    sql.NullInt64   `gorm:"column:order_time" json:"order_time"`
	OrderType    sql.NullString  `gorm:"column:order_type" json:"order_type"`
	Orderid      sql.NullString  `gorm:"column:orderid" json:"orderid"`
	Status       sql.NullInt64   `gorm:"column:status" json:"status"`
	StoreID      sql.NullInt64   `gorm:"column:store_id" json:"store_id"`
	UpdateTime   sql.NullInt64   `gorm:"column:update_time" json:"update_time"`
	VerifyType   sql.NullString  `gorm:"column:verify_type" json:"verify_type"`
}

// TableName sets the insert table name for this struct type
func (w *WebDisOrder) TableName() string {
	return "web_dis_order"
}
