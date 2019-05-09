package models

import (
	"database/sql"
	"time"
)

type GItempicture struct {
	ItemLib           string         `gorm:"column:ItemLib;primary_key" json:"ItemLib;primary_key"`
	ModifyDate        time.Time      `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP          sql.NullString `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode  sql.NullString `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName  sql.NullString `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation sql.NullString `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	Picture           []byte         `gorm:"column:Picture" json:"Picture"`
	Picture2          []byte         `gorm:"column:Picture2" json:"Picture2"`
	Type              sql.NullString `gorm:"column:Type" json:"Type"`
	Type2             sql.NullString `gorm:"column:Type2" json:"Type2"`
	Cid               sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Itemcode          string         `gorm:"column:itemcode;primary_key" json:"itemcode;primary_key"`
}

// TableName sets the insert table name for this struct type
func (g *GItempicture) TableName() string {
	return "g_itempicture"
}
