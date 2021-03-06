package models

import "database/sql"

type CSalespromotionp struct {
	Attribute1         sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2         sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3         sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4         sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5         sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6         sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7         sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8         sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	EnglishName        sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode              sql.NullString  `gorm:"column:FCode" json:"FCode"`
	GroupNo            sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	IndexID            int             `gorm:"column:IndexID;primary_key" json:"IndexID;primary_key"`
	IsCanIntegral      sql.NullInt64   `gorm:"column:IsCanIntegral" json:"IsCanIntegral"`
	ItemLib            sql.NullString  `gorm:"column:ItemLib" json:"ItemLib"`
	ItemName           sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	OldPrice           sql.NullFloat64 `gorm:"column:OldPrice" json:"OldPrice"`
	OrderID            sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	PromotionPrice     sql.NullFloat64 `gorm:"column:PromotionPrice" json:"PromotionPrice"`
	Remarks            sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	SalesPromotionCode string          `gorm:"column:SalesPromotionCode;primary_key" json:"SalesPromotionCode;primary_key"`
	Unit               sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode        sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	Cid                sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode           sql.NullString  `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (c *CSalespromotionp) TableName() string {
	return "c_salespromotionp"
}
