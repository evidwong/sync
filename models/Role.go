package models

import "database/sql"

type Role struct {
	Cid     int            `gorm:"column:cid" json:"cid"`
	ID      int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name    string         `gorm:"column:name" json:"name"`
	Pid     sql.NullInt64  `gorm:"column:pid" json:"pid"`
	Remark  sql.NullString `gorm:"column:remark" json:"remark"`
	Status  sql.NullInt64  `gorm:"column:status" json:"status"`
	StoreID sql.NullString `gorm:"column:store_id" json:"store_id"`
}

// TableName sets the insert table name for this struct type
func (r *Role) TableName() string {
	return "role"
}
