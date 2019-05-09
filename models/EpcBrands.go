package models

import "database/sql"

type EpcBrands struct {
	Acronym  sql.NullString `gorm:"column:acronym" json:"acronym"`
	AreaName sql.NullString `gorm:"column:area_name" json:"area_name"`
	Brand    string         `gorm:"column:brand" json:"brand"`
	ID       int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Img      sql.NullString `gorm:"column:img" json:"img"`
	Name     sql.NullString `gorm:"column:name" json:"name"`
	URI      sql.NullString `gorm:"column:uri" json:"uri"`
}

// TableName sets the insert table name for this struct type
func (e *EpcBrands) TableName() string {
	return "epc_brands"
}
