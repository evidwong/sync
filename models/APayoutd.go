package models

import "database/sql"

type APayoutd struct {
	BillNo        sql.NullString  `gorm:"column:BillNo" json:"BillNo"`
	Charge        sql.NullString  `gorm:"column:Charge" json:"Charge"`
	FunctionCode  sql.NullString  `gorm:"column:FunctionCode" json:"FunctionCode"`
	IndexID       sql.NullInt64   `gorm:"column:IndexID" json:"IndexID"`
	OrderID       sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	PayoutAmount  sql.NullFloat64 `gorm:"column:PayoutAmount" json:"PayoutAmount"`
	PayoutCode    sql.NullString  `gorm:"column:PayoutCode" json:"PayoutCode"`
	PayoutRemarks sql.NullString  `gorm:"column:PayoutRemarks" json:"PayoutRemarks"`
	Purpose       sql.NullString  `gorm:"column:Purpose" json:"Purpose"`
	Cid           int             `gorm:"column:cid" json:"cid"`
	ID            int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Image         sql.NullString  `gorm:"column:image" json:"image"`
	Status        int             `gorm:"column:status" json:"status"`
}

// TableName sets the insert table name for this struct type
func (a *APayoutd) TableName() string {
	return "a_payoutd"
}
