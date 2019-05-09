package models

import (
	"database/sql"
	"time"
)

type SNewno struct {
	COMNo        sql.NullString `gorm:"column:COMNo" json:"COMNo"`
	FunctionCode string         `gorm:"column:FunctionCode;primary_key" json:"FunctionCode;primary_key"`
	IsUse        sql.NullInt64  `gorm:"column:IsUse" json:"IsUse"`
	LastDate     time.Time      `gorm:"column:LastDate" json:"LastDate"`
	LastNo       int            `gorm:"column:LastNo" json:"LastNo"`
	Cid          sql.NullInt64  `gorm:"column:cid" json:"cid"`
	StoreID      sql.NullInt64  `gorm:"column:store_id" json:"store_id"`
}

// TableName sets the insert table name for this struct type
func (s *SNewno) TableName() string {
	return "s_newno"
}
