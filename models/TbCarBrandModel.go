package models

import "database/sql"

type TbCarBrandModel struct {
	BrandName    sql.NullString `gorm:"column:brandName" json:"brandName"`
	CarLine      sql.NullString `gorm:"column:carLine" json:"carLine"`
	ID           int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	ImgName      sql.NullString `gorm:"column:imgName" json:"imgName"`
	Manufacturer sql.NullString `gorm:"column:manufacturer" json:"manufacturer"`
	PyChr        string         `gorm:"column:py_chr" json:"py_chr"`
}

// TableName sets the insert table name for this struct type
func (t *TbCarBrandModel) TableName() string {
	return "tb_car_brand_model"
}
