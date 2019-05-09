package models

import (
	"database/sql"
	"time"
)

type PInoutlog struct {
	Action                sql.NullString  `gorm:"column:Action" json:"Action"`
	ActionID              int             `gorm:"column:ActionID;primary_key" json:"ActionID;primary_key"`
	Attribute1            sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2            sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3            sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4            sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5            sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6            sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7            sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8            sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	BillCode              sql.NullString  `gorm:"column:BillCode" json:"BillCode"`
	BillDate              time.Time       `gorm:"column:BillDate" json:"BillDate"`
	BillIndexID           sql.NullInt64   `gorm:"column:BillIndexID" json:"BillIndexID"`
	Bin                   sql.NullString  `gorm:"column:Bin" json:"Bin"`
	BorrowQty             sql.NullFloat64 `gorm:"column:BorrowQty" json:"BorrowQty"`
	CReceptionCode        sql.NullString  `gorm:"column:CReceptionCode" json:"CReceptionCode"`
	ComNo                 sql.NullString  `gorm:"column:ComNo" json:"ComNo"`
	CompanyCode           sql.NullString  `gorm:"column:CompanyCode" json:"CompanyCode"`
	CompanyName           sql.NullString  `gorm:"column:CompanyName" json:"CompanyName"`
	CostPrice             sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	EnglishName           sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode                 sql.NullString  `gorm:"column:FCode" json:"FCode"`
	FunctionCode          sql.NullString  `gorm:"column:FunctionCode" json:"FunctionCode"`
	GroupNo               sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	InCostAmount          sql.NullFloat64 `gorm:"column:InCostAmount" json:"InCostAmount"`
	InQty                 sql.NullFloat64 `gorm:"column:InQty" json:"InQty"`
	ItemName              sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	Location              sql.NullString  `gorm:"column:Location" json:"Location"`
	LogDate               time.Time       `gorm:"column:LogDate" json:"LogDate"`
	OperatorCode          sql.NullString  `gorm:"column:OperatorCode" json:"OperatorCode"`
	OperatorName          sql.NullString  `gorm:"column:OperatorName" json:"OperatorName"`
	OutCostAmount         sql.NullFloat64 `gorm:"column:OutCostAmount" json:"OutCostAmount"`
	OutQty                sql.NullFloat64 `gorm:"column:OutQty" json:"OutQty"`
	PYCode                sql.NullString  `gorm:"column:PYCode" json:"PYCode"`
	Price                 sql.NullFloat64 `gorm:"column:Price" json:"Price"`
	PriceAmount           sql.NullFloat64 `gorm:"column:PriceAmount" json:"PriceAmount"`
	ReseQty               sql.NullFloat64 `gorm:"column:ReseQty" json:"ReseQty"`
	SalesQty              sql.NullFloat64 `gorm:"column:SalesQty" json:"SalesQty"`
	StockBorrowCostAmount sql.NullFloat64 `gorm:"column:StockBorrowCostAmount" json:"StockBorrowCostAmount"`
	StockCostAmount       sql.NullFloat64 `gorm:"column:StockCostAmount" json:"StockCostAmount"`
	StockNoTaxCostAmount  sql.NullFloat64 `gorm:"column:StockNoTaxCostAmount" json:"StockNoTaxCostAmount"`
	StockQty              sql.NullFloat64 `gorm:"column:StockQty" json:"StockQty"`
	StockReseCostAmount   sql.NullFloat64 `gorm:"column:StockReseCostAmount" json:"StockReseCostAmount"`
	StockSalesCostAmount  sql.NullFloat64 `gorm:"column:StockSalesCostAmount" json:"StockSalesCostAmount"`
	TmpField1             sql.NullString  `gorm:"column:TmpField1" json:"TmpField1"`
	Unit                  sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode           sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	WBCode                sql.NullString  `gorm:"column:WBCode" json:"WBCode"`
	Attributecode         sql.NullString  `gorm:"column:attributecode" json:"attributecode"`
	CFunctionCode         sql.NullString  `gorm:"column:cFunctionCode" json:"cFunctionCode"`
	Cid                   sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Isdel                 int             `gorm:"column:isdel" json:"isdel"`
	Itemcode              sql.NullString  `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (p *PInoutlog) TableName() string {
	return "p_inoutlog"
}
