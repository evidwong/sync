package models

import (
	"database/sql"
	"time"
)
//AAccountin ...
type AAccountin struct {
	AccType             sql.NullString  `gorm:"column:AccType" json:"AccType"`
	AccountDate         time.Time       `gorm:"column:AccountDate" json:"AccountDate"`
	AccountPersonCode   sql.NullString  `gorm:"column:AccountPersonCode" json:"AccountPersonCode"`
	AccountPersonName   sql.NullString  `gorm:"column:AccountPersonName" json:"AccountPersonName"`
	ActionID            sql.NullInt64   `gorm:"column:ActionID" json:"ActionID"`
	Amount              sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	Bank                sql.NullString  `gorm:"column:Bank" json:"Bank"`
	BankAccount         sql.NullString  `gorm:"column:BankAccount" json:"BankAccount"`
	BillCode            sql.NullString  `gorm:"column:BillCode" json:"BillCode"`
	BillIndexID         sql.NullInt64   `gorm:"column:BillIndexID" json:"BillIndexID"`
	BillNo2             sql.NullString  `gorm:"column:BillNo2" json:"BillNo2"`
	BusinessType        sql.NullString  `gorm:"column:BusinessType" json:"BusinessType"`
	COMNo               sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	CustomerCode        sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName        sql.NullString  `gorm:"column:CustomerName" json:"CustomerName"`
	DepartmentM         sql.NullString  `gorm:"column:DepartmentM" json:"DepartmentM"`
	DiscountAmount      sql.NullFloat64 `gorm:"column:DiscountAmount" json:"DiscountAmount"`
	DiscountPerson      sql.NullString  `gorm:"column:DiscountPerson" json:"DiscountPerson"`
	DocDate             time.Time       `gorm:"column:DocDate" json:"DocDate"`
	EarningItem2        sql.NullString  `gorm:"column:EarningItem2" json:"EarningItem2"`
	FunctionCode        sql.NullString  `gorm:"column:FunctionCode" json:"FunctionCode"`
	InvoiceDate         time.Time       `gorm:"column:InvoiceDate" json:"InvoiceDate"`
	InvoiceNo           sql.NullString  `gorm:"column:InvoiceNo" json:"InvoiceNo"`
	InvoiceType         sql.NullString  `gorm:"column:InvoiceType" json:"InvoiceType"`
	OperationType       sql.NullString  `gorm:"column:OperationType" json:"OperationType"`
	OutArrearSettleCode sql.NullString  `gorm:"column:OutArrearSettleCode" json:"OutArrearSettleCode"`
	PreReceiveAmount    sql.NullFloat64 `gorm:"column:PreReceiveAmount" json:"PreReceiveAmount"`
	Purpose2            sql.NullString  `gorm:"column:Purpose2" json:"Purpose2"`
	RegisterNo          sql.NullString  `gorm:"column:RegisterNo" json:"RegisterNo"`
	Remarks             sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	Remarks2            sql.NullString  `gorm:"column:Remarks2" json:"Remarks2"`
	Remarks3            sql.NullString  `gorm:"column:Remarks3" json:"Remarks3"`
	Remarks4            sql.NullString  `gorm:"column:Remarks4" json:"Remarks4"`
	SourceFunction      sql.NullString  `gorm:"column:SourceFunction" json:"SourceFunction"`
	Status              int             `gorm:"column:Status" json:"Status"`
	ValueCardNo         sql.NullString  `gorm:"column:ValueCardNo" json:"ValueCardNo"`
	VehicleCode         sql.NullString  `gorm:"column:VehicleCode" json:"VehicleCode"`
	Cid                 sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ID                  int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel               int             `gorm:"column:isdel" json:"isdel"`
}

// TableName sets the insert table name for this struct type
func (a *AAccountin) TableName() string {
	return "a_accountin"
}
