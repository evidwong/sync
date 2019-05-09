package models

import "database/sql"

type PMovestockd struct {
	Attribute1      sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2      sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3      sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4      sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5      sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6      sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7      sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8      sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	CostPrice       sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	DepartmentD     sql.NullString  `gorm:"column:DepartmentD" json:"DepartmentD"`
	DestStockID     sql.NullInt64   `gorm:"column:DestStockID" json:"DestStockID"`
	EnglishName     sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode           sql.NullString  `gorm:"column:FCode" json:"FCode"`
	FromBin         sql.NullString  `gorm:"column:FromBin" json:"FromBin"`
	FromLocation    sql.NullString  `gorm:"column:FromLocation" json:"FromLocation"`
	GroupNo         sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	IndexID         int             `gorm:"column:IndexID;primary_key" json:"IndexID;primary_key"`
	ItemCode        string          `gorm:"column:ItemCode" json:"ItemCode"`
	ItemName        sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	NoTaxCostPrice  sql.NullFloat64 `gorm:"column:NoTaxCostPrice" json:"NoTaxCostPrice"`
	OrderID         sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	PMoveStockCode  string          `gorm:"column:PMoveStockCode;primary_key" json:"PMoveStockCode;primary_key"`
	PackQty         sql.NullFloat64 `gorm:"column:PackQty" json:"PackQty"`
	PackUnit        sql.NullString  `gorm:"column:PackUnit" json:"PackUnit"`
	PackagePieceQty sql.NullFloat64 `gorm:"column:PackagePieceQty" json:"PackagePieceQty"`
	PackageQty      sql.NullInt64   `gorm:"column:PackageQty" json:"PackageQty"`
	Qty             sql.NullFloat64 `gorm:"column:Qty" json:"Qty"`
	StockID         sql.NullInt64   `gorm:"column:StockID" json:"StockID"`
	ToBin           sql.NullString  `gorm:"column:ToBin" json:"ToBin"`
	ToLocation      sql.NullString  `gorm:"column:ToLocation" json:"ToLocation"`
	Unit            sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode     sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	Cid             sql.NullInt64   `gorm:"column:cid" json:"cid"`
}

// TableName sets the insert table name for this struct type
func (p *PMovestockd) TableName() string {
	return "p_movestockd"
}
