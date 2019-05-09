package models

import (
	"database/sql"
	"time"
)

type GAttributetype struct {
	AddTime       time.Time       `gorm:"column:add_time" json:"add_time"`
	Attributecode string          `gorm:"column:attributecode" json:"attributecode"`
	Attributetype string          `gorm:"column:attributetype" json:"attributetype"`
	Bonuspercent  sql.NullFloat64 `gorm:"column:bonuspercent" json:"bonuspercent"`
	Cid           sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Comno         sql.NullString  `gorm:"column:comno" json:"comno"`
	Comnocode     sql.NullString  `gorm:"column:comnocode" json:"comnocode"`
	ID            int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Lastno        string          `gorm:"column:lastno" json:"lastno"`
	Level         int             `gorm:"column:level" json:"level"`
	Pid           string          `gorm:"column:pid" json:"pid"`
	Pname         sql.NullString  `gorm:"column:pname" json:"pname"`
	Status        sql.NullInt64   `gorm:"column:status" json:"status"`
	StoreID       sql.NullString  `gorm:"column:store_id" json:"store_id"`
	StoreName     sql.NullString  `gorm:"column:store_name" json:"store_name"`
	UpdateTime    time.Time       `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (g *GAttributetype) TableName() string {
	return "g_attributetype"
}
