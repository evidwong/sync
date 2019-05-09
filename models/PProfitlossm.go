package models

import (
	"database/sql"
	"time"
)

type PProfitlossm struct {
	ApprovalDate          time.Time       `gorm:"column:ApprovalDate" json:"ApprovalDate"`
	ApproverCode          sql.NullString  `gorm:"column:ApproverCode" json:"ApproverCode"`
	ApproverName          sql.NullString  `gorm:"column:ApproverName" json:"ApproverName"`
	COMNo                 string          `gorm:"column:COMNo" json:"COMNo"`
	DocDate               time.Time       `gorm:"column:DocDate" json:"DocDate"`
	ModifyDate            time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP              sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode      sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName      sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation     sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	OperatorCode          sql.NullString  `gorm:"column:OperatorCode" json:"OperatorCode"`
	OperatorName          sql.NullString  `gorm:"column:OperatorName" json:"OperatorName"`
	PProfitLossCode       string          `gorm:"column:PProfitLossCode" json:"PProfitLossCode"`
	PStockCheckCode       sql.NullString  `gorm:"column:PStockCheckCode" json:"PStockCheckCode"`
	PinCode               sql.NullString  `gorm:"column:PinCode" json:"PinCode"`
	Remarks               sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	TotalLossCostAmount   sql.NullFloat64 `gorm:"column:TotalLossCostAmount" json:"TotalLossCostAmount"`
	TotalLossQty          sql.NullFloat64 `gorm:"column:TotalLossQty" json:"TotalLossQty"`
	TotalProfitCostAmount sql.NullFloat64 `gorm:"column:TotalProfitCostAmount" json:"TotalProfitCostAmount"`
	TotalProfitQty        sql.NullFloat64 `gorm:"column:TotalProfitQty" json:"TotalProfitQty"`
	Cid                   sql.NullInt64   `gorm:"column:cid" json:"cid"`
}

// TableName sets the insert table name for this struct type
func (p *PProfitlossm) TableName() string {
	return "p_profitlossm"
}
