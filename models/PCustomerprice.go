package models

import (
	"database/sql"
	"time"
)

type PCustomerprice struct {
	AccType             sql.NullString  `gorm:"column:AccType" json:"AccType"`
	Attribute1          sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2          sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3          sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4          sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5          sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6          sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7          sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8          sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	CustomerCode        sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName        sql.NullString  `gorm:"column:CustomerName" json:"CustomerName"`
	CustomerPriceSelect sql.NullString  `gorm:"column:CustomerPriceSelect" json:"CustomerPriceSelect"`
	EnglishName         sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode               sql.NullString  `gorm:"column:FCode" json:"FCode"`
	GroupNo             sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	InvoiceType         sql.NullString  `gorm:"column:InvoiceType" json:"InvoiceType"`
	ItemName            sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	LogDate             time.Time       `gorm:"column:LogDate" json:"LogDate"`
	OperatorCode        sql.NullString  `gorm:"column:OperatorCode" json:"OperatorCode"`
	OperatorName        sql.NullString  `gorm:"column:OperatorName" json:"OperatorName"`
	PCustPriceCode      string          `gorm:"column:PCustPriceCode;primary_key" json:"PCustPriceCode;primary_key"`
	Price               sql.NullFloat64 `gorm:"column:Price" json:"Price"`
	Remarks             sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	Unit                sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode         sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	Cid                 sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode            string          `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (p *PCustomerprice) TableName() string {
	return "p_customerprice"
}
