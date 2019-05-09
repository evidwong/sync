package models

import "database/sql"

type PBatchpriceadjustd struct {
	ActAdjustFieldName    sql.NullString  `gorm:"column:ActAdjustFieldName" json:"ActAdjustFieldName"`
	AdjustPercent         sql.NullFloat64 `gorm:"column:AdjustPercent" json:"AdjustPercent"`
	Attribute1            sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2            sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3            sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4            sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5            sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6            sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7            sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8            sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	COMNo                 sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	EnglishName           sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode                 sql.NullString  `gorm:"column:FCode" json:"FCode"`
	GroupNo               sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	IndexID               int             `gorm:"column:IndexID;primary_key" json:"IndexID;primary_key"`
	ItemName              sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	NewPrice              sql.NullFloat64 `gorm:"column:NewPrice" json:"NewPrice"`
	OldPrice              sql.NullFloat64 `gorm:"column:OldPrice" json:"OldPrice"`
	OrderID               sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	PBatchPriceAdjustCode string          `gorm:"column:PBatchPriceAdjustCode;primary_key" json:"PBatchPriceAdjustCode;primary_key"`
	StockID               sql.NullInt64   `gorm:"column:StockID" json:"StockID"`
	Unit                  sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode           sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	Cid                   sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode              string          `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (p *PBatchpriceadjustd) TableName() string {
	return "p_batchpriceadjustd"
}
