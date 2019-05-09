package models

import (
	"database/sql"
	"time"
)

type DailyreportMainD struct {
	Amount     sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	COMNo      sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	CostAmount sql.NullFloat64 `gorm:"column:CostAmount" json:"CostAmount"`
	Date       time.Time       `gorm:"column:Date" json:"Date"`
	Qty        sql.NullFloat64 `gorm:"column:Qty" json:"Qty"`
	RecType    sql.NullString  `gorm:"column:RecType" json:"RecType"`
	Cid        sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ID         int             `gorm:"column:id;primary_key" json:"id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (d *DailyreportMainD) TableName() string {
	return "dailyreport_main_d"
}
