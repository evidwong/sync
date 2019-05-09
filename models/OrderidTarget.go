package models

import "database/sql"

type OrderidTarget struct {
	Cid      sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Comno    sql.NullString `gorm:"column:comno" json:"comno"`
	Dateint  sql.NullInt64  `gorm:"column:dateint" json:"dateint"`
	ID       int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Pre      sql.NullString `gorm:"column:pre" json:"pre"`
	Targetid sql.NullInt64  `gorm:"column:targetid" json:"targetid"`
}

// TableName sets the insert table name for this struct type
func (o *OrderidTarget) TableName() string {
	return "orderid_target"
}
