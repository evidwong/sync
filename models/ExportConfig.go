package models

import "database/sql"

type ExportConfig struct {
	AddTime   sql.NullInt64  `gorm:"column:add_time" json:"add_time"`
	Cid       sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Columns   sql.NullString `gorm:"column:columns" json:"columns"`
	ID        int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	TableName sql.NullString `gorm:"column:table_name" json:"table_name"`
	Type      sql.NullString `gorm:"column:type" json:"type"`
	UID       sql.NullInt64  `gorm:"column:uid" json:"uid"`
}

