package models

import "database/sql"

type BaseConfig struct {
	Cid     sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Content sql.NullString `gorm:"column:content" json:"content"`
	ID      int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Pid     sql.NullInt64  `gorm:"column:pid" json:"pid"`
	Type    sql.NullString `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (b *BaseConfig) TableName() string {
	return "base_config"
}
