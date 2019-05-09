package models

import (
	"database/sql"
	"time"
)

type PProfitlossd struct {
	ActualQty        sql.NullFloat64 `gorm:"column:ActualQty" json:"ActualQty"`
	Attribute1       sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2       sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3       sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4       sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5       sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6       sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7       sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8       sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	Bin              sql.NullString  `gorm:"column:Bin" json:"Bin"`
	CostAmount       sql.NullFloat64 `gorm:"column:CostAmount" json:"CostAmount"`
	CostPrice        sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	EnglishName      sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode            sql.NullString  `gorm:"column:FCode" json:"FCode"`
	GroupNo          sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	IndexID          int             `gorm:"column:IndexID" json:"IndexID"`
	ItemCode         string          `gorm:"column:ItemCode" json:"ItemCode"`
	ItemName         sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	Location         sql.NullString  `gorm:"column:Location" json:"Location"`
	NoTaxCostPrice   sql.NullFloat64 `gorm:"column:NoTaxCostPrice" json:"NoTaxCostPrice"`
	OrderID          int             `gorm:"column:OrderID" json:"OrderID"`
	PProfitLossCode  string          `gorm:"column:PProfitLossCode" json:"PProfitLossCode"`
	ProfitLossReason sql.NullString  `gorm:"column:ProfitLossReason" json:"ProfitLossReason"`
	Qty              sql.NullFloat64 `gorm:"column:Qty" json:"Qty"`
	StockID          sql.NullInt64   `gorm:"column:StockID" json:"StockID"`
	StockQty         sql.NullFloat64 `gorm:"column:StockQty" json:"StockQty"`
	SupplierCode     sql.NullString  `gorm:"column:SupplierCode" json:"SupplierCode"`
	SupplierName     sql.NullString  `gorm:"column:SupplierName" json:"SupplierName"`
	Unit             sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode      sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	ValidDate        time.Time       `gorm:"column:ValidDate" json:"ValidDate"`
	Cid              sql.NullInt64   `gorm:"column:cid" json:"cid"`
}

// TableName sets the insert table name for this struct type
func (p *PProfitlossd) TableName() string {
	return "p_profitlossd"
}
