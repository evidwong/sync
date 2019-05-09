package models

import (
	"database/sql"
	"time"
)

type PInoutprofit struct {
	ActionID     int             `gorm:"column:ActionID;primary_key" json:"ActionID;primary_key"`
	Attribute3   sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4   sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	BillCode     string          `gorm:"column:BillCode" json:"BillCode"`
	BillDate     time.Time       `gorm:"column:BillDate" json:"BillDate"`
	ComNo        string          `gorm:"column:ComNo" json:"ComNo"`
	CostAmount   sql.NullFloat64 `gorm:"column:CostAmount" json:"CostAmount"`
	CostPrice    sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	CustomerCode sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName sql.NullString  `gorm:"column:CustomerName" json:"CustomerName"`
	FunctionCode string          `gorm:"column:FunctionCode" json:"FunctionCode"`
	InCostPrice  sql.NullFloat64 `gorm:"column:InCostPrice" json:"InCostPrice"`
	ItemName     sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	OutCostPrice sql.NullFloat64 `gorm:"column:OutCostPrice" json:"OutCostPrice"`
	Qty          sql.NullFloat64 `gorm:"column:Qty" json:"Qty"`
	Unit         sql.NullString  `gorm:"column:Unit" json:"Unit"`
	Cid          sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode     string          `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (p *PInoutprofit) TableName() string {
	return "p_inoutprofit"
}
