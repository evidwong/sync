package models

import (
	"database/sql"
	"time"
)

type JPartsbusiness struct {
	AddTime      time.Time       `gorm:"column:add_time" json:"add_time"`
	Bonuspercent sql.NullFloat64 `gorm:"column:bonuspercent" json:"bonuspercent"`
	Cid          sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Comno        sql.NullString  `gorm:"column:comno" json:"comno"`
	Comnocode    sql.NullString  `gorm:"column:comnocode" json:"comnocode"`
	Depot        sql.NullString  `gorm:"column:depot" json:"depot"`
	Depotid      sql.NullInt64   `gorm:"column:depotid" json:"depotid"`
	GroupName    sql.NullString  `gorm:"column:group_name" json:"group_name"`
	ID           int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Indexid      sql.NullInt64   `gorm:"column:indexid" json:"indexid"`
	Name         sql.NullString  `gorm:"column:name" json:"name"`
	Pid          sql.NullInt64   `gorm:"column:pid" json:"pid"`
	StoreName    sql.NullString  `gorm:"column:store_name" json:"store_name"`
	Storeid      sql.NullString  `gorm:"column:storeid" json:"storeid"`
	UpdateTime   time.Time       `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (j *JPartsbusiness) TableName() string {
	return "j_partsbusiness"
}
