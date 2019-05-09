package models

import (
	"database/sql"
	"time"
)

type DailyreportIncrementD struct {
	Amount  sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	COMNo   sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	Date    time.Time       `gorm:"column:Date" json:"Date"`
	Rectype sql.NullString  `gorm:"column:Rectype" json:"Rectype"`
	Cid     sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ID      int             `gorm:"column:id;primary_key" json:"id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (d *DailyreportIncrementD) TableName() string {
	return "dailyreport_increment_d"
}
