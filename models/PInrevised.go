package models

import "database/sql"

type PInrevised struct {
	Attribute1           sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2           sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3           sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4           sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5           sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6           sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7           sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8           sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	EnglishName          sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode                sql.NullString  `gorm:"column:FCode" json:"FCode"`
	GroupNo              sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	InCostPrice          sql.NullFloat64 `gorm:"column:InCostPrice" json:"InCostPrice"`
	InNoTaxCostPrice     sql.NullFloat64 `gorm:"column:InNoTaxCostPrice" json:"InNoTaxCostPrice"`
	IndexID              int             `gorm:"column:IndexID;primary_key" json:"IndexID;primary_key"`
	ItemName             sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	OrderID              sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	PInReviseCode        string          `gorm:"column:PInReviseCode;primary_key" json:"PInReviseCode;primary_key"`
	Qty                  sql.NullFloat64 `gorm:"column:Qty" json:"Qty"`
	ReviseCostPrice      sql.NullFloat64 `gorm:"column:ReviseCostPrice" json:"ReviseCostPrice"`
	ReviseDiffCostPrice  sql.NullFloat64 `gorm:"column:ReviseDiffCostPrice" json:"ReviseDiffCostPrice"`
	ReviseNoTaxCostPrice sql.NullFloat64 `gorm:"column:ReviseNoTaxCostPrice" json:"ReviseNoTaxCostPrice"`
	StockID              sql.NullInt64   `gorm:"column:StockID" json:"StockID"`
	Unit                 sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode          sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	Cid                  sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode             sql.NullString  `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (p *PInrevised) TableName() string {
	return "p_inrevised"
}
