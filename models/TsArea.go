package models

import (
	"database/sql"
	"time"
)

type TsArea struct {
	Abbreviation sql.NullString `gorm:"column:abbreviation" json:"abbreviation"`
	Areaname     string         `gorm:"column:areaname" json:"areaname"`
	Createtime   time.Time      `gorm:"column:createtime" json:"createtime"`
	Disabled     sql.NullInt64  `gorm:"column:disabled" json:"disabled"`
	FullAddress  sql.NullString `gorm:"column:full_address" json:"full_address"`
	ID           int64          `gorm:"column:id;primary_key" json:"id;primary_key"`
	Level        int            `gorm:"column:level" json:"level"`
	ParentID     int64          `gorm:"column:parent_id" json:"parent_id"`
	Sequence     int            `gorm:"column:sequence" json:"sequence"`
}

// TableName sets the insert table name for this struct type
func (t *TsArea) TableName() string {
	return "ts_area"
}
