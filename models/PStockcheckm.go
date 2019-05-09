package models

import (
	"database/sql"
	"time"
)

type PStockcheckm struct {
	Approval          sql.NullInt64  `gorm:"column:Approval" json:"Approval"`
	ApprovalDate      time.Time      `gorm:"column:ApprovalDate" json:"ApprovalDate"`
	ApproverCode      sql.NullString `gorm:"column:ApproverCode" json:"ApproverCode"`
	ApproverName      sql.NullString `gorm:"column:ApproverName" json:"ApproverName"`
	COMNo             string         `gorm:"column:COMNo" json:"COMNo"`
	CostPrice         float64        `gorm:"column:CostPrice" json:"CostPrice"`
	DisuseDate        time.Time      `gorm:"column:DisuseDate" json:"DisuseDate"`
	DisuseReason      sql.NullString `gorm:"column:DisuseReason" json:"DisuseReason"`
	DocDate           time.Time      `gorm:"column:DocDate" json:"DocDate"`
	IsDel             int            `gorm:"column:IsDel" json:"IsDel"`
	ModifyDate        time.Time      `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP          sql.NullString `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode  sql.NullString `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName  sql.NullString `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation sql.NullString `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	OperatorCode      sql.NullString `gorm:"column:OperatorCode" json:"OperatorCode"`
	OperatorName      sql.NullString `gorm:"column:OperatorName" json:"OperatorName"`
	PProfitLossCode   sql.NullString `gorm:"column:PProfitLossCode" json:"PProfitLossCode"`
	PStockCheckCode   string         `gorm:"column:PStockCheckCode" json:"PStockCheckCode"`
	Qty               float32        `gorm:"column:Qty" json:"Qty"`
	Remarks           sql.NullString `gorm:"column:Remarks" json:"Remarks"`
	Cid               int            `gorm:"column:cid" json:"cid"`
}

// TableName sets the insert table name for this struct type
func (p *PStockcheckm) TableName() string {
	return "p_stockcheckm"
}
