package models

import "database/sql"

type WebGroup struct {
	AddTime int            `gorm:"column:add_time" json:"add_time"`
	Cid     int            `gorm:"column:cid" json:"cid"`
	ID      int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Pid     int            `gorm:"column:pid" json:"pid"`
	Rules   sql.NullString `gorm:"column:rules" json:"rules"`
	Sort    int            `gorm:"column:sort" json:"sort"`
	Status  int            `gorm:"column:status" json:"status"`
	Title   string         `gorm:"column:title" json:"title"`
}

// TableName sets the insert table name for this struct type
func (w *WebGroup) TableName() string {
	return "web_group"
}
