package models

import "database/sql"

type TbBrand struct {
	FirstChr string         `gorm:"column:first_chr" json:"first_chr"`
	ID       int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	ImgName  sql.NullString `gorm:"column:imgName" json:"imgName"`
	IsHot    int            `gorm:"column:is_hot" json:"is_hot"`
	Name     sql.NullString `gorm:"column:name" json:"name"`
}

// TableName sets the insert table name for this struct type
func (t *TbBrand) TableName() string {
	return "tb_brand"
}
