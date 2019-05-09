package models

import (
	"database/sql"
	"time"
)

type DailyreportAmountpaidD struct {
	Amount sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	COMNo  sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	Charge sql.NullString  `gorm:"column:Charge" json:"Charge"`
	Date   time.Time       `gorm:"column:Date" json:"Date"`
	Cid    sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ID     int             `gorm:"column:id;primary_key" json:"id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (d *DailyreportAmountpaidD) TableName() string {
	return "dailyreport_amountpaid_d"
}
