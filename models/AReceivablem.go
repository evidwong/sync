package models

import (
	"database/sql"
	"time"
)

type AReceivablem struct {
	ArrearLimitAmount  sql.NullFloat64 `gorm:"column:ArrearLimitAmount" json:"ArrearLimitAmount"`
	BadDebtsAmount     sql.NullFloat64 `gorm:"column:BadDebtsAmount" json:"BadDebtsAmount"`
	BalanceAmount      sql.NullFloat64 `gorm:"column:BalanceAmount" json:"BalanceAmount"`
	COMNo              sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	CreditCycle        sql.NullInt64   `gorm:"column:CreditCycle" json:"CreditCycle"`
	CreditDate         time.Time       `gorm:"column:CreditDate" json:"CreditDate"`
	CustomerCode       sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName       sql.NullString  `gorm:"column:CustomerName" json:"CustomerName"`
	CyclePayDay        sql.NullInt64   `gorm:"column:CyclePayDay" json:"CyclePayDay"`
	CyclePayMonth      sql.NullInt64   `gorm:"column:CyclePayMonth" json:"CyclePayMonth"`
	DelayReceiveAmount sql.NullFloat64 `gorm:"column:DelayReceiveAmount" json:"DelayReceiveAmount"`
	DelayReceiveDate   time.Time       `gorm:"column:DelayReceiveDate" json:"DelayReceiveDate"`
	DiscountAmount     sql.NullFloat64 `gorm:"column:DiscountAmount" json:"DiscountAmount"`
	IsCyclePay         sql.NullInt64   `gorm:"column:IsCyclePay" json:"IsCyclePay"`
	ModifyDate         time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	OccurAmount        sql.NullFloat64 `gorm:"column:OccurAmount" json:"OccurAmount"`
	PreReceiveAmount   sql.NullFloat64 `gorm:"column:PreReceiveAmount" json:"PreReceiveAmount"`
	ReceiveAmount      sql.NullFloat64 `gorm:"column:ReceiveAmount" json:"ReceiveAmount"`
	Cid                int             `gorm:"column:cid" json:"cid"`
	ID                 int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel              int             `gorm:"column:isdel" json:"isdel"`
}

// TableName sets the insert table name for this struct type
func (a *AReceivablem) TableName() string {
	return "a_receivablem"
}
