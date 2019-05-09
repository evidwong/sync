package models

import (
	"database/sql"
	"time"
)

type SActionlog struct {
	FunctionCode string         `gorm:"column:FunctionCode" json:"FunctionCode"`
	AddTime      time.Time      `gorm:"column:add_time" json:"add_time"`
	Cid          sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Code         sql.NullString `gorm:"column:code" json:"code"`
	Comno        sql.NullString `gorm:"column:comno" json:"comno"`
	Content      sql.NullString `gorm:"column:content" json:"content"`
	FunctionCodes sql.NullString `gorm:"column:function_code" json:"function_code"`
	FunctionType string         `gorm:"column:function_type" json:"function_type"`
	ID           int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	IP           sql.NullString `gorm:"column:ip" json:"ip"`
	OperatorCode sql.NullString `gorm:"column:operator_code" json:"operator_code"`
	OperatorName sql.NullString `gorm:"column:operator_name" json:"operator_name"`
	StoreID      sql.NullInt64  `gorm:"column:store_id" json:"store_id"`
	StoreName    sql.NullString `gorm:"column:store_name" json:"store_name"`
}

// TableName sets the insert table name for this struct type
func (s *SActionlog) TableName() string {
	return "s_actionlog"
}
