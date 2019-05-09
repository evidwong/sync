package models

import (
	"database/sql"
	"time"
)

type APayoutm struct {
	AccAppDate           time.Time       `gorm:"column:AccAppDate" json:"AccAppDate"`
	AccAppPersonCode     sql.NullString  `gorm:"column:AccAppPersonCode" json:"AccAppPersonCode"`
	AccAppPersonName     sql.NullString  `gorm:"column:AccAppPersonName" json:"AccAppPersonName"`
	AccType              sql.NullString  `gorm:"column:AccType" json:"AccType"`
	ApprovalDate         time.Time       `gorm:"column:ApprovalDate" json:"ApprovalDate"`
	ApproverCode         sql.NullString  `gorm:"column:ApproverCode" json:"ApproverCode"`
	ApproverName         sql.NullString  `gorm:"column:ApproverName" json:"ApproverName"`
	ArrearDiscountAmount sql.NullFloat64 `gorm:"column:ArrearDiscountAmount" json:"ArrearDiscountAmount"`
	ArrearPayAmount      sql.NullFloat64 `gorm:"column:ArrearPayAmount" json:"ArrearPayAmount"`
	ArrearPrePayAmount   sql.NullFloat64 `gorm:"column:ArrearPrePayAmount" json:"ArrearPrePayAmount"`
	BankAccountCode      sql.NullString  `gorm:"column:BankAccountCode" json:"BankAccountCode"`
	BankAccountName      sql.NullString  `gorm:"column:BankAccountName" json:"BankAccountName"`
	COMNo                sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	CustomerCode         sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName         sql.NullString  `gorm:"column:CustomerName" json:"CustomerName"`
	Department           sql.NullString  `gorm:"column:Department" json:"Department"`
	DepartmentID         sql.NullInt64   `gorm:"column:Department_id" json:"Department_id"`
	DocDate              time.Time       `gorm:"column:DocDate" json:"DocDate"`
	DocState             sql.NullString  `gorm:"column:DocState" json:"DocState"`
	FunctionCode         sql.NullString  `gorm:"column:FunctionCode" json:"FunctionCode"`
	HandlerCode          sql.NullString  `gorm:"column:HandlerCode" json:"HandlerCode"`
	HandlerName          sql.NullString  `gorm:"column:HandlerName" json:"HandlerName"`
	HandlerNameID        sql.NullInt64   `gorm:"column:HandlerName_id" json:"HandlerName_id"`
	IsAccApp             sql.NullInt64   `gorm:"column:IsAccApp" json:"IsAccApp"`
	IsArrear             sql.NullInt64   `gorm:"column:IsArrear" json:"IsArrear"`
	ManualBillNo         sql.NullString  `gorm:"column:ManualBillNo" json:"ManualBillNo"`
	ModifyDate           time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP             sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode     sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName     sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation    sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	MustPayDate          time.Time       `gorm:"column:MustPayDate" json:"MustPayDate"`
	NoteNo               sql.NullString  `gorm:"column:NoteNo" json:"NoteNo"`
	OperatorCode         sql.NullString  `gorm:"column:OperatorCode" json:"OperatorCode"`
	OperatorName         sql.NullString  `gorm:"column:OperatorName" json:"OperatorName"`
	PayoutCode           sql.NullString  `gorm:"column:PayoutCode" json:"PayoutCode"`
	Remarks              sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	TotalPayoutAmount    sql.NullFloat64 `gorm:"column:TotalPayoutAmount" json:"TotalPayoutAmount"`
	UnitCode             sql.NullString  `gorm:"column:UnitCode" json:"UnitCode"`
	UnitName             sql.NullString  `gorm:"column:UnitName" json:"UnitName"`
	Cid                  int             `gorm:"column:cid" json:"cid"`
	ID                   int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Image                sql.NullString  `gorm:"column:image" json:"image"`
	IsDel                int             `gorm:"column:is_del" json:"is_del"`
	OldData              sql.NullString  `gorm:"column:old_data" json:"old_data"`
	Status               int             `gorm:"column:status" json:"status"`
}

// TableName sets the insert table name for this struct type
func (a *APayoutm) TableName() string {
	return "a_payoutm"
}
