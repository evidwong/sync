package models

import "database/sql"

type TbDetailed struct {
	BrandName    sql.NullString `gorm:"column:brandName" json:"brandName"`
	CarLine      sql.NullString `gorm:"column:carLine" json:"carLine"`
	CarModel     sql.NullString `gorm:"column:carModel" json:"carModel"`
	ID           int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Manufacturer sql.NullString `gorm:"column:manufacturer" json:"manufacturer"`
}

// TableName sets the insert table name for this struct type
func (t *TbDetailed) TableName() string {
	return "tb_detailed"
}
