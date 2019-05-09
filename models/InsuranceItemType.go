package models

import "database/sql"

type InsuranceItemType struct {
	AddTime    sql.NullInt64 `gorm:"column:add_time" json:"add_time"`
	Aid        sql.NullInt64 `gorm:"column:aid" json:"aid"`
	Auid       sql.NullInt64 `gorm:"column:auid" json:"auid"`
	Channel    sql.NullInt64 `gorm:"column:channel" json:"channel"`
	Cid        int           `gorm:"column:cid" json:"cid"`
	ID         int           `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel      sql.NullInt64 `gorm:"column:isdel" json:"isdel"`
	Name       string        `gorm:"column:name" json:"name"`
	Status     sql.NullInt64 `gorm:"column:status" json:"status"`
	UpdateTime sql.NullInt64 `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (i *InsuranceItemType) TableName() string {
	return "insurance_item_type"
}
