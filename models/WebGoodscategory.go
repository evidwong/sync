package models

import (
	"database/sql"
	"time"
)

type WebGoodscategory struct {
	Cid                   sql.NullInt64 `gorm:"column:cid" json:"cid"`
	CreateTime            time.Time     `gorm:"column:create_time" json:"create_time"`
	GoodscategoryDescribe string        `gorm:"column:goodscategory_describe" json:"goodscategory_describe"`
	GoodscategoryName     string        `gorm:"column:goodscategory_name" json:"goodscategory_name"`
	ID                    int           `gorm:"column:id;primary_key" json:"id;primary_key"`
	Image                 string        `gorm:"column:image" json:"image"`
	Pid                   sql.NullInt64 `gorm:"column:pid" json:"pid"`
	Thumb                 string        `gorm:"column:thumb" json:"thumb"`
	UpdateTime            time.Time     `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (w *WebGoodscategory) TableName() string {
	return "web_goodscategory"
}
