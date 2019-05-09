package models

import (
	"database/sql"
	"time"
)

type CExpensev struct {
	Amount          sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	CExpenseCode    string          `gorm:"column:CExpenseCode" json:"CExpenseCode"`
	CReceptionCode  sql.NullString  `gorm:"column:CReceptionCode" json:"CReceptionCode"`
	FunctionCode    sql.NullString  `gorm:"column:FunctionCode" json:"FunctionCode"`
	IndexID         int             `gorm:"column:IndexID" json:"IndexID"`
	JobCode         sql.NullString  `gorm:"column:JobCode" json:"JobCode"`
	JobName         sql.NullString  `gorm:"column:JobName" json:"JobName"`
	JobType         sql.NullString  `gorm:"column:JobType" json:"JobType"`
	NextServiceDate time.Time       `gorm:"column:NextServiceDate" json:"NextServiceDate"`
	OrderID         sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	Proponent       sql.NullString  `gorm:"column:Proponent" json:"Proponent"`
	Remarks         sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	Cid             int             `gorm:"column:cid" json:"cid"`
	ID              int             `gorm:"column:id;primary_key" json:"id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (c *CExpensev) TableName() string {
	return "c_expensev"
}
