package models

import (
	"database/sql"
	"time"
)

type Dailyreport struct {
	COMNo     sql.NullString `gorm:"column:COMNo" json:"COMNo"`
	DocDate   time.Time      `gorm:"column:DocDate" json:"DocDate"`
	EditDate  time.Time      `gorm:"column:EditDate" json:"EditDate"`
	Isdel     sql.NullInt64  `gorm:"column:Isdel" json:"Isdel"`
	RecType   sql.NullString `gorm:"column:RecType" json:"RecType"`
	TableName sql.NullString `gorm:"column:TableName" json:"TableName"`
	Cid       sql.NullInt64  `gorm:"column:cid" json:"cid"`
	ID        int            `gorm:"column:id;primary_key" json:"id;primary_key"`
}
