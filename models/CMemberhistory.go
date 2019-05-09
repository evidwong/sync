package models

import (
	"database/sql"
	"time"
)

type CMemberhistory struct {
	AccountAmount  sql.NullFloat64 `gorm:"column:AccountAmount" json:"AccountAmount"`
	Amount         sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	BillCode       sql.NullString  `gorm:"column:BillCode" json:"BillCode"`
	BonusAmount    sql.NullFloat64 `gorm:"column:BonusAmount" json:"BonusAmount"`
	BusinessType   sql.NullString  `gorm:"column:BusinessType" json:"BusinessType"`
	COMNo          sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	CustomerCode   sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName   sql.NullString  `gorm:"column:CustomerName" json:"CustomerName"`
	DescriptionStr sql.NullString  `gorm:"column:DescriptionStr" json:"DescriptionStr"`
	DocDate        time.Time       `gorm:"column:DocDate" json:"DocDate"`
	FunctionCode   sql.NullString  `gorm:"column:FunctionCode" json:"FunctionCode"`
	IntegralChange sql.NullFloat64 `gorm:"column:IntegralChange" json:"IntegralChange"`
	LogDate        time.Time       `gorm:"column:LogDate" json:"LogDate"`
	MemberCardNo   sql.NullString  `gorm:"column:MemberCardNo" json:"MemberCardNo"`
	MemberCode     sql.NullString  `gorm:"column:MemberCode" json:"MemberCode"`
	MemberType     sql.NullString  `gorm:"column:MemberType" json:"MemberType"`
	OffsetAmount   sql.NullFloat64 `gorm:"column:OffsetAmount" json:"OffsetAmount"`
	OperationType  sql.NullString  `gorm:"column:OperationType" json:"OperationType"`
	OperatorCode   sql.NullString  `gorm:"column:OperatorCode" json:"OperatorCode"`
	OperatorName   sql.NullString  `gorm:"column:OperatorName" json:"OperatorName"`
	RegisterNo     sql.NullString  `gorm:"column:RegisterNo" json:"RegisterNo"`
	Remarks        sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	SalesCode      sql.NullString  `gorm:"column:SalesCode" json:"SalesCode"`
	SalesName      sql.NullString  `gorm:"column:SalesName" json:"SalesName"`
	VehicleCode    sql.NullString  `gorm:"column:VehicleCode" json:"VehicleCode"`
	Cid            int             `gorm:"column:cid" json:"cid"`
	ID             int             `gorm:"column:id;primary_key" json:"id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (c *CMemberhistory) TableName() string {
	return "c_memberhistory"
}
