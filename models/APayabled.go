package models

import (
	"database/sql"
	"time"
)

type APayabled struct {
	AccType            sql.NullString  `gorm:"column:AccType" json:"AccType"`
	AccountDate        time.Time       `gorm:"column:AccountDate" json:"AccountDate"`
	AccountPersonCode  sql.NullString  `gorm:"column:AccountPersonCode" json:"AccountPersonCode"`
	AccountPersonName  sql.NullString  `gorm:"column:AccountPersonName" json:"AccountPersonName"`
	AccountRemarks     sql.NullString  `gorm:"column:AccountRemarks" json:"AccountRemarks"`
	ActionID           int             `gorm:"column:ActionID" json:"ActionID"`
	ArrearAmount       sql.NullFloat64 `gorm:"column:ArrearAmount" json:"ArrearAmount"`
	ArrearPayAmount    sql.NullFloat64 `gorm:"column:ArrearPayAmount" json:"ArrearPayAmount"`
	ArrearPrePayAmount sql.NullFloat64 `gorm:"column:ArrearPrePayAmount" json:"ArrearPrePayAmount"`
	BalanceAmount      sql.NullFloat64 `gorm:"column:BalanceAmount" json:"BalanceAmount"`
	BillCode           sql.NullString  `gorm:"column:BillCode" json:"BillCode"`
	BillFun            sql.NullString  `gorm:"column:BillFun" json:"BillFun"`
	BillIndexID        sql.NullInt64   `gorm:"column:BillIndexID" json:"BillIndexID"`
	BusinessType       sql.NullString  `gorm:"column:BusinessType" json:"BusinessType"`
	BuyerCode          sql.NullString  `gorm:"column:BuyerCode" json:"BuyerCode"`
	BuyerName          sql.NullString  `gorm:"column:BuyerName" json:"BuyerName"`
	COMNo              string          `gorm:"column:COMNo" json:"COMNo"`
	DepartmentM        sql.NullString  `gorm:"column:DepartmentM" json:"DepartmentM"`
	DiscountAmount     sql.NullFloat64 `gorm:"column:DiscountAmount" json:"DiscountAmount"`
	DocDate            time.Time       `gorm:"column:DocDate" json:"DocDate"`
	FunctionCode       sql.NullString  `gorm:"column:FunctionCode" json:"FunctionCode"`
	InvoiceDate        time.Time       `gorm:"column:InvoiceDate" json:"InvoiceDate"`
	InvoiceNo          sql.NullString  `gorm:"column:InvoiceNo" json:"InvoiceNo"`
	InvoiceType        sql.NullString  `gorm:"column:InvoiceType" json:"InvoiceType"`
	IsArrearage        int             `gorm:"column:IsArrearage" json:"IsArrearage"`
	IsOverMustPayDate  int             `gorm:"column:IsOverMustPayDate" json:"IsOverMustPayDate"`
	ModifyDate         time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	MustPayDate        time.Time       `gorm:"column:MustPayDate" json:"MustPayDate"`
	OperationType      sql.NullString  `gorm:"column:OperationType" json:"OperationType"`
	RegisterNo         sql.NullString  `gorm:"column:RegisterNo" json:"RegisterNo"`
	Remarks            sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	SourceFunction     sql.NullString  `gorm:"column:SourceFunction" json:"SourceFunction"`
	SupplierCode       string          `gorm:"column:SupplierCode" json:"SupplierCode"`
	UserRemarks        sql.NullString  `gorm:"column:UserRemarks" json:"UserRemarks"`
	VehicleCode        sql.NullString  `gorm:"column:VehicleCode" json:"VehicleCode"`
	Cid                int             `gorm:"column:cid" json:"cid"`
	ID                 int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel              int             `gorm:"column:isdel" json:"isdel"`
}

// TableName sets the insert table name for this struct type
func (a *APayabled) TableName() string {
	return "a_payabled"
}
