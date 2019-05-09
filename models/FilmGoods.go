package models

import "database/sql"

type FilmGoods struct {
	AddTime    int           `gorm:"column:add_time" json:"add_time"`
	BrandType  int           `gorm:"column:brand_type" json:"brand_type"`
	Cid        int           `gorm:"column:cid" json:"cid"`
	ID         int           `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name       string        `gorm:"column:name" json:"name"`
	Pid        int           `gorm:"column:pid" json:"pid"`
	Recommond  int           `gorm:"column:recommond" json:"recommond"`
	Sort       int           `gorm:"column:sort" json:"sort"`
	Status     sql.NullInt64 `gorm:"column:status" json:"status"`
	UpdateTime int           `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (f *FilmGoods) TableName() string {
	return "film_goods"
}
