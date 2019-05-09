package models

import (
	"database/sql"
	"time"
)

type CIntegralgift struct {
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
	GroupNo           sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	Integral          sql.NullFloat64 `gorm:"column:Integral" json:"Integral"`
	ItemLib           string          `gorm:"column:ItemLib;primary_key" json:"ItemLib;primary_key"`
	ItemName          sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	ModifyDate        time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP          sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode  sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName  sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	Unit              sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode       sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	Cid               sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode          string          `gorm:"column:itemcode;primary_key" json:"itemcode;primary_key"`
}

// TableName sets the insert table name for this struct type
func (c *CIntegralgift) TableName() string {
	return "c_integralgift"
}
