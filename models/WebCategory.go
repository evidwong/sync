package models

import (
	"database/sql"
	"time"
)

type WebCategory struct {
	Cid        sql.NullInt64  `gorm:"column:cid" json:"cid"`
	CreateTime time.Time      `gorm:"column:create_time" json:"create_time"`
	ID         int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Image      string         `gorm:"column:image" json:"image"`
	Name       sql.NullString `gorm:"column:name" json:"name"`
	Pid        int            `gorm:"column:pid" json:"pid"`
	Thumb      string         `gorm:"column:thumb" json:"thumb"`
	Type       sql.NullInt64  `gorm:"column:type" json:"type"`
	UpdateTime time.Time      `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (w *WebCategory) TableName() string {
	return "web_category"
}
