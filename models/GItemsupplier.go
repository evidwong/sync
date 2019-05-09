package models

import "database/sql"

type GItemsupplier struct {
	COMNo        sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	IndexID      int             `gorm:"column:IndexID;primary_key" json:"IndexID;primary_key"`
	IsDefault    sql.NullInt64   `gorm:"column:IsDefault" json:"IsDefault"`
	ItemCode     string          `gorm:"column:ItemCode;primary_key" json:"ItemCode;primary_key"`
	MinOrderQty  sql.NullFloat64 `gorm:"column:MinOrderQty" json:"MinOrderQty"`
	OrderID      sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	OrderPrice   sql.NullFloat64 `gorm:"column:OrderPrice" json:"OrderPrice"`
	Remarks      sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	SupplierCode sql.NullString  `gorm:"column:SupplierCode" json:"SupplierCode"`
	SupplierName sql.NullString  `gorm:"column:SupplierName" json:"SupplierName"`
}

// TableName sets the insert table name for this struct type
func (g *GItemsupplier) TableName() string {
	return "g_itemsupplier"
}
