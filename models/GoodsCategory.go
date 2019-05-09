package models

import "database/sql"

type GoodsCategory struct {
	Addtime     int           `gorm:"column:addtime" json:"addtime"`
	Alias       string        `gorm:"column:alias" json:"alias"`
	Cid         int           `gorm:"column:cid" json:"cid"`
	Description string        `gorm:"column:description" json:"description"`
	ID          int           `gorm:"column:id;primary_key" json:"id;primary_key"`
	Image       string        `gorm:"column:image" json:"image"`
	Keyword     string        `gorm:"column:keyword" json:"keyword"`
	Name        string        `gorm:"column:name" json:"name"`
	Pid         int           `gorm:"column:pid" json:"pid"`
	ShowIndex   int           `gorm:"column:showIndex" json:"showIndex"`
	Sort        int           `gorm:"column:sort" json:"sort"`
	Status      sql.NullInt64 `gorm:"column:status" json:"status"`
	Type        int           `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (g *GoodsCategory) TableName() string {
	return "goods_category"
}
