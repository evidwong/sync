package models

import (
	"database/sql"
	"time"
)

type JAccount struct {
	AddTime    time.Time      `gorm:"column:add_time" json:"add_time"`
	Category   sql.NullString `gorm:"column:category" json:"category"`
	Cid        sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Comno      sql.NullString `gorm:"column:comno" json:"comno"`
	Comnocode  sql.NullString `gorm:"column:comnocode" json:"comnocode"`
	GroupName  sql.NullString `gorm:"column:group_name" json:"group_name"`
	ID         int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Indexid    sql.NullInt64  `gorm:"column:indexid" json:"indexid"`
	Name       sql.NullString `gorm:"column:name" json:"name"`
	Pid        sql.NullInt64  `gorm:"column:pid" json:"pid"`
	Price      sql.NullString `gorm:"column:price" json:"price"`
	StoreID    sql.NullString `gorm:"column:store_id" json:"store_id"`
	StoreName  sql.NullString `gorm:"column:store_name" json:"store_name"`
	UpdateTime time.Time      `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (j *JAccount) TableName() string {
	return "j_account"
}
