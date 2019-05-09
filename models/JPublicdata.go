package models

import (
	"database/sql"
	"time"
)

type JPublicdata struct {
	AddTime        time.Time       `gorm:"column:add_time" json:"add_time"`
	Cid            sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Comno          sql.NullString  `gorm:"column:comno" json:"comno"`
	Comnocode      sql.NullString  `gorm:"column:comnocode" json:"comnocode"`
	GroupName      sql.NullString  `gorm:"column:group_name" json:"group_name"`
	ID             int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Indexid        sql.NullInt64   `gorm:"column:indexid" json:"indexid"`
	Isarrear       sql.NullInt64   `gorm:"column:isarrear" json:"isarrear"`
	Isotheracctype sql.NullInt64   `gorm:"column:isotheracctype" json:"isotheracctype"`
	Linkman        sql.NullString  `gorm:"column:linkman" json:"linkman"`
	Name           sql.NullString  `gorm:"column:name" json:"name"`
	Phone          sql.NullString  `gorm:"column:phone" json:"phone"`
	Pid            sql.NullInt64   `gorm:"column:pid" json:"pid"`
	StoreID        sql.NullString  `gorm:"column:store_id" json:"store_id"`
	StoreName      sql.NullString  `gorm:"column:store_name" json:"store_name"`
	Taxrate        sql.NullFloat64 `gorm:"column:taxrate" json:"taxrate"`
	UpdateTime     time.Time       `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (j *JPublicdata) TableName() string {
	return "j_publicdata"
}
