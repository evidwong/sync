package models

import "database/sql"

type CInsbillp struct {
	Attribute1     sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2     sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3     sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4     sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5     sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6     sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7     sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8     sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	CInsBillCode   string          `gorm:"column:CInsBillCode;primary_key" json:"CInsBillCode;primary_key"`
	CPAmount       sql.NullFloat64 `gorm:"column:CPAmount" json:"CPAmount"`
	CostPrice      sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	EnglishName    sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode          sql.NullString  `gorm:"column:FCode" json:"FCode"`
	GroupNo        sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	IndexID        int             `gorm:"column:IndexID;primary_key" json:"IndexID;primary_key"`
	ItemName       sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	OrderID        sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	PAmendsRemarks sql.NullString  `gorm:"column:PAmendsRemarks" json:"PAmendsRemarks"`
	PPayAmount     sql.NullFloat64 `gorm:"column:PPayAmount" json:"PPayAmount"`
	PPrice         sql.NullFloat64 `gorm:"column:PPrice" json:"PPrice"`
	PQty           sql.NullFloat64 `gorm:"column:PQty" json:"PQty"`
	Unit           sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode    sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	WPrice         sql.NullFloat64 `gorm:"column:WPrice" json:"WPrice"`
	WQty           sql.NullFloat64 `gorm:"column:WQty" json:"WQty"`
	WarrantAmount  sql.NullFloat64 `gorm:"column:WarrantAmount" json:"WarrantAmount"`
	Cid            sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode       string          `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (c *CInsbillp) TableName() string {
	return "c_insbillp"
}
