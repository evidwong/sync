package models

import "database/sql"

type BRechargeSetting struct {
	AddTime     sql.NullInt64   `gorm:"column:add_time" json:"add_time"`
	ID          int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isshow      sql.NullInt64   `gorm:"column:isshow" json:"isshow"`
	OptUID      sql.NullInt64   `gorm:"column:opt_uid" json:"opt_uid"`
	OptUsername sql.NullString  `gorm:"column:opt_username" json:"opt_username"`
	Price       sql.NullFloat64 `gorm:"column:price" json:"price"`
	Quantity    sql.NullInt64   `gorm:"column:quantity" json:"quantity"`
	Title       sql.NullString  `gorm:"column:title" json:"title"`
	Type        sql.NullString  `gorm:"column:type" json:"type"`
	UpdateTime  sql.NullInt64   `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (b *BRechargeSetting) TableName() string {
	return "b_recharge_setting"
}
