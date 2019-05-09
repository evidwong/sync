package models

import (
	"database/sql"
	"time"
)

type APayablem struct {
	BalanceAmount  sql.NullFloat64 `gorm:"column:BalanceAmount" json:"BalanceAmount"`
	COMNo          sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	DiscountAmount sql.NullFloat64 `gorm:"column:DiscountAmount" json:"DiscountAmount"`
	ModifyDate     time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	OccurAmount    sql.NullFloat64 `gorm:"column:OccurAmount" json:"OccurAmount"`
	PayAmount      sql.NullFloat64 `gorm:"column:PayAmount" json:"PayAmount"`
	PrePayAmount   sql.NullFloat64 `gorm:"column:PrePayAmount" json:"PrePayAmount"`
	SupplierCode   sql.NullString  `gorm:"column:SupplierCode" json:"SupplierCode"`
	SupplierName   sql.NullString  `gorm:"column:SupplierName" json:"SupplierName"`
	Cid            int             `gorm:"column:cid" json:"cid"`
	ID             int             `gorm:"column:id;primary_key" json:"id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (a *APayablem) TableName() string {
	return "a_payablem"
}
