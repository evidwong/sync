package models

import (
	"database/sql"
	"time"
)

type AReceivabled struct {
	AccType                sql.NullString  `gorm:"column:AccType" json:"AccType"`
	AccountDate            time.Time       `gorm:"column:AccountDate" json:"AccountDate"`
	AccountPersonCode      sql.NullString  `gorm:"column:AccountPersonCode" json:"AccountPersonCode"`
	AccountPersonName      sql.NullString  `gorm:"column:AccountPersonName" json:"AccountPersonName"`
	AccountRemarks         sql.NullString  `gorm:"column:AccountRemarks" json:"AccountRemarks"`
	ActionID               sql.NullInt64   `gorm:"column:ActionID" json:"ActionID"`
	ArrearAmount           sql.NullFloat64 `gorm:"column:ArrearAmount" json:"ArrearAmount"`
	ArrearPreReceiveAmount sql.NullFloat64 `gorm:"column:ArrearPreReceiveAmount" json:"ArrearPreReceiveAmount"`
	ArrearReceiveAmount    sql.NullFloat64 `gorm:"column:ArrearReceiveAmount" json:"ArrearReceiveAmount"`
	BadDebtsAmount         sql.NullFloat64 `gorm:"column:BadDebtsAmount" json:"BadDebtsAmount"`
	BadDebtsDate           time.Time       `gorm:"column:BadDebtsDate" json:"BadDebtsDate"`
	BadDebtsPersonCode     sql.NullString  `gorm:"column:BadDebtsPersonCode" json:"BadDebtsPersonCode"`
	BadDebtsPersonName     sql.NullString  `gorm:"column:BadDebtsPersonName" json:"BadDebtsPersonName"`
	BadDebtsReason         sql.NullString  `gorm:"column:BadDebtsReason" json:"BadDebtsReason"`
	BalanceAmount          sql.NullFloat64 `gorm:"column:BalanceAmount" json:"BalanceAmount"`
	BillCode               sql.NullString  `gorm:"column:BillCode" json:"BillCode"`
	BillFun                sql.NullString  `gorm:"column:BillFun" json:"BillFun"`
	BillIndexID            sql.NullInt64   `gorm:"column:BillIndexID" json:"BillIndexID"`
	BusinessType           sql.NullString  `gorm:"column:BusinessType" json:"BusinessType"`
	COMNo                  sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	CustomerCode           sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	DepartmentM            sql.NullString  `gorm:"column:DepartmentM" json:"DepartmentM"`
	DiscountAmount         sql.NullFloat64 `gorm:"column:DiscountAmount" json:"DiscountAmount"`
	DiscountPerson         sql.NullString  `gorm:"column:DiscountPerson" json:"DiscountPerson"`
	DocDate                time.Time       `gorm:"column:DocDate" json:"DocDate"`
	FunctionCode           sql.NullString  `gorm:"column:FunctionCode" json:"FunctionCode"`
	InvoiceDate            time.Time       `gorm:"column:InvoiceDate" json:"InvoiceDate"`
	InvoiceNo              sql.NullString  `gorm:"column:InvoiceNo" json:"InvoiceNo"`
	InvoiceType            sql.NullString  `gorm:"column:InvoiceType" json:"InvoiceType"`
	IsArrearage            sql.NullInt64   `gorm:"column:IsArrearage" json:"IsArrearage"`
	IsBadDebts             sql.NullInt64   `gorm:"column:IsBadDebts" json:"IsBadDebts"`
	IsOverMustPayDate      sql.NullInt64   `gorm:"column:IsOverMustPayDate" json:"IsOverMustPayDate"`
	LinkCustomerCode       sql.NullString  `gorm:"column:LinkCustomerCode" json:"LinkCustomerCode"`
	LinkCustomerName       sql.NullString  `gorm:"column:LinkCustomerName" json:"LinkCustomerName"`
	ModifyDate             time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	MustPayDate            time.Time       `gorm:"column:MustPayDate" json:"MustPayDate"`
	OperationType          sql.NullString  `gorm:"column:OperationType" json:"OperationType"`
	RegisterNo             sql.NullString  `gorm:"column:RegisterNo" json:"RegisterNo"`
	Remarks                sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	SalesmanCode           sql.NullString  `gorm:"column:SalesmanCode" json:"SalesmanCode"`
	SalesmanName           sql.NullString  `gorm:"column:SalesmanName" json:"SalesmanName"`
	SourceFunction         sql.NullString  `gorm:"column:SourceFunction" json:"SourceFunction"`
	UserRemarks            sql.NullString  `gorm:"column:UserRemarks" json:"UserRemarks"`
	VehicleCode            sql.NullString  `gorm:"column:VehicleCode" json:"VehicleCode"`
	Cid                    int             `gorm:"column:cid" json:"cid"`
	ID                     int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel                  int             `gorm:"column:isdel" json:"isdel"`
}

// TableName sets the insert table name for this struct type
func (a *AReceivabled) TableName() string {
	return "a_receivabled"
}
