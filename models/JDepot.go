package models

import (
	"database/sql"
	"time"
)

type JDepot struct {
	AddTime    time.Time      `gorm:"column:add_time" json:"add_time"`
	Cid        sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Comno      sql.NullString `gorm:"column:comno" json:"comno"`
	Comnocode  sql.NullString `gorm:"column:comnocode" json:"comnocode"`
	Discount   sql.NullInt64  `gorm:"column:discount" json:"discount"`
	Handphone  sql.NullString `gorm:"column:handphone" json:"handphone"`
	ID         int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Indexid    sql.NullInt64  `gorm:"column:indexid" json:"indexid"`
	Name       sql.NullString `gorm:"column:name" json:"name"`
	StoreID    sql.NullString `gorm:"column:store_id" json:"store_id"`
	StoreName  sql.NullString `gorm:"column:store_name" json:"store_name"`
	Surpass    sql.NullString `gorm:"column:surpass" json:"surpass"`
	UpdateTime time.Time      `gorm:"column:update_time" json:"update_time"`
	Waste      sql.NullInt64  `gorm:"column:waste" json:"waste"`
}

// TableName sets the insert table name for this struct type
func (j *JDepot) TableName() string {
	return "j_depot"
}
