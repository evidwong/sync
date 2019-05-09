package models

import (
	"database/sql"
	"time"
)

type JRectype struct {
	AddTime      time.Time      `gorm:"column:add_time" json:"add_time"`
	CallbackTime sql.NullInt64  `gorm:"column:callback_time" json:"callback_time"`
	Cid          sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Comno        sql.NullString `gorm:"column:comno" json:"comno"`
	Comnocode    sql.NullString `gorm:"column:comnocode" json:"comnocode"`
	ExpectTime   sql.NullInt64  `gorm:"column:expect_time" json:"expect_time"`
	GroupName    sql.NullString `gorm:"column:group_name" json:"group_name"`
	ID           int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Indexid      sql.NullInt64  `gorm:"column:indexid" json:"indexid"`
	Name         sql.NullString `gorm:"column:name" json:"name"`
	Pid          sql.NullInt64  `gorm:"column:pid" json:"pid"`
	Rectype      sql.NullString `gorm:"column:rectype" json:"rectype"`
	StoreID      sql.NullString `gorm:"column:store_id" json:"store_id"`
	StoreName    sql.NullString `gorm:"column:store_name" json:"store_name"`
	UpdateTime   time.Time      `gorm:"column:update_time" json:"update_time"`
	Worktype     sql.NullString `gorm:"column:worktype" json:"worktype"`
}

// TableName sets the insert table name for this struct type
func (j *JRectype) TableName() string {
	return "j_rectype"
}
