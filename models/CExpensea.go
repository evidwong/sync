package models

import "database/sql"

type CExpensea struct {
	AddIn          sql.NullString  `gorm:"column:AddIn" json:"AddIn"`
	AddInPrice     sql.NullFloat64 `gorm:"column:AddInPrice" json:"AddInPrice"`
	AddInType      sql.NullString  `gorm:"column:AddInType" json:"AddInType"`
	Amount         sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	CExpenseCode   string          `gorm:"column:CExpenseCode;primary_key" json:"CExpenseCode;primary_key"`
	CIWPCode       sql.NullString  `gorm:"column:CIWPCode" json:"CIWPCode"`
	CReceptionCode sql.NullString  `gorm:"column:CReceptionCode" json:"CReceptionCode"`
	Class          sql.NullString  `gorm:"column:Class" json:"Class"`
	CostPrice      sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	DepartmentD    sql.NullString  `gorm:"column:DepartmentD" json:"DepartmentD"`
	DiscountRate   sql.NullFloat64 `gorm:"column:DiscountRate" json:"DiscountRate"`
	IndexID        int             `gorm:"column:IndexID;primary_key" json:"IndexID;primary_key"`
	IsCalcAddIn    int             `gorm:"column:IsCalcAddIn" json:"IsCalcAddIn"`
	OrderID        sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	Price          sql.NullFloat64 `gorm:"column:Price" json:"Price"`
	Qty            sql.NullFloat64 `gorm:"column:Qty" json:"Qty"`
	Unit           sql.NullString  `gorm:"column:Unit" json:"Unit"`
}

// TableName sets the insert table name for this struct type
func (c *CExpensea) TableName() string {
	return "c_expensea"
}
