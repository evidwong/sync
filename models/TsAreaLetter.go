package models

import "database/sql"

type TsAreaLetter struct {
	AreaName sql.NullString `gorm:"column:areaName" json:"areaName"`
	Areaid   int            `gorm:"column:areaid" json:"areaid"`
	ID       int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Letter   string         `gorm:"column:letter" json:"letter"`
}

// TableName sets the insert table name for this struct type
func (t *TsAreaLetter) TableName() string {
	return "ts_area_letter"
}
