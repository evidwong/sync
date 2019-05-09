package models

import "database/sql"

type Banner struct {
	AddTime sql.NullInt64 `gorm:"column:add_time" json:"add_time"`
	Cid     int           `gorm:"column:cid" json:"cid"`
	ID      int           `gorm:"column:id;primary_key" json:"id;primary_key"`
	Imgurl  string        `gorm:"column:imgurl" json:"imgurl"`
	Link    string        `gorm:"column:link" json:"link"`
	Name    string        `gorm:"column:name" json:"name"`
	Sort    int           `gorm:"column:sort" json:"sort"`
	Status  sql.NullInt64 `gorm:"column:status" json:"status"`
	Thumb   string        `gorm:"column:thumb" json:"thumb"`
	Type    int           `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (b *Banner) TableName() string {
	return "banner"
}
