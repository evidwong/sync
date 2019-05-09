package models

import "database/sql"

type TableDesc struct {
	ChineseName    sql.NullString `gorm:"column:chinese_name" json:"chinese_name"`
	ColumnName     sql.NullString `gorm:"column:column_name" json:"column_name"`
	ColumnType     sql.NullString `gorm:"column:column_type" json:"column_type"`
	Comnent        sql.NullString `gorm:"column:comnent" json:"comnent"`
	DefaultValue   sql.NullString `gorm:"column:default_value" json:"default_value"`
	ID             int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isshow         sql.NullInt64  `gorm:"column:isshow" json:"isshow"`
	TableName      sql.NullString `gorm:"column:table_name" json:"table_name"`
	UpdateIP       sql.NullString `gorm:"column:update_ip" json:"update_ip"`
	UpdateStation  sql.NullString `gorm:"column:update_station" json:"update_station"`
	UpdateTime     sql.NullString `gorm:"column:update_time" json:"update_time"`
	UpdateUID      sql.NullString `gorm:"column:update_uid" json:"update_uid"`
	UpdateUsername sql.NullString `gorm:"column:update_username" json:"update_username"`
}
