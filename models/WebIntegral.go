package models

import "database/sql"

//WebIntegral ...
type WebIntegral struct {
	AddTime      sql.NullInt64  `gorm:"column:add_time" json:"add_time"`
	Cid          sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Flag         sql.NullInt64  `gorm:"column:flag" json:"flag"`
	FlagContent  sql.NullString `gorm:"column:flag_content" json:"flag_content"`
	HandPhone    sql.NullString `gorm:"column:hand_phone" json:"hand_phone"`
	ID           int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Integral     sql.NullInt64  `gorm:"column:integral" json:"integral"`
	OptUID       sql.NullInt64  `gorm:"column:opt_uid" json:"opt_uid"`
	OptUsername  sql.NullString `gorm:"column:opt_username" json:"opt_username"`
	OrderType    sql.NullString `gorm:"column:order_type" json:"order_type"`
	Orderid      sql.NullString `gorm:"column:orderid" json:"orderid"`
	Remark       sql.NullString `gorm:"column:remark" json:"remark"`
	Status       sql.NullInt64  `gorm:"column:status" json:"status"`
	UID          sql.NullInt64  `gorm:"column:uid" json:"uid"`
	Level        int            `gorm:"column:level" json:"level"`
	CustomerName sql.NullString `gorm:"column:customer_name" json:"customer_name"`
	FromUID      int            `gorm:"column:from_uid" json:"from_uid"`
}

// TableName sets the insert table name for this struct type
func (w *WebIntegral) TableName() string {
	return "web_integral"
}
