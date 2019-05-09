package models

import (
	"database/sql"
	"time"
)

type DailyreportInfactpayD struct {
	AccType sql.NullString  `gorm:"column:AccType" json:"AccType"`
	Amount  sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	COMNo   sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	Date    time.Time       `gorm:"column:Date" json:"Date"`
	Cid     sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ID      int             `gorm:"column:id;primary_key" json:"id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (d *DailyreportInfactpayD) TableName() string {
	return "dailyreport_infactpay_d"
}
