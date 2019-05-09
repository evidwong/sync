package models

import "database/sql"

type PInquirepriced struct {
	Attribute1        sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2        sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3        sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4        sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5        sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6        sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7        sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8        sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	EnglishName       sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode             sql.NullString  `gorm:"column:FCode" json:"FCode"`
	FOBCurrency       sql.NullString  `gorm:"column:FOBCurrency" json:"FOBCurrency"`
	FOBPrice          sql.NullFloat64 `gorm:"column:FOBPrice" json:"FOBPrice"`
	FOBRate           sql.NullFloat64 `gorm:"column:FOBRate" json:"FOBRate"`
	GroupNo           sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	IndexID           int             `gorm:"column:IndexID;primary_key" json:"IndexID;primary_key"`
	InquireRemarks    sql.NullString  `gorm:"column:InquireRemarks" json:"InquireRemarks"`
	ItemName          sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	OrderID           sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	PInquirePriceCode string          `gorm:"column:PInquirePriceCode;primary_key" json:"PInquirePriceCode;primary_key"`
	Price             sql.NullFloat64 `gorm:"column:Price" json:"Price"`
	Qty               sql.NullFloat64 `gorm:"column:Qty" json:"Qty"`
	Unit              sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode       sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	Cid               sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode          string          `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (p *PInquirepriced) TableName() string {
	return "p_inquirepriced"
}
