package models

import (
	"database/sql"
	"time"
)

type SystemParameter struct {
	Cid             sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Describe        sql.NullString `gorm:"column:describe" json:"describe"`
	Edittime        time.Time      `gorm:"column:edittime" json:"edittime"`
	GroupName       sql.NullString `gorm:"column:groupName" json:"groupName"`
	ID              int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Index           sql.NullInt64  `gorm:"column:index" json:"index"`
	OptionalRange   sql.NullString `gorm:"column:optionalRange" json:"optionalRange"`
	ParameterName   sql.NullString `gorm:"column:parameterName" json:"parameterName"`
	ParameterType   sql.NullString `gorm:"column:parameterType" json:"parameterType"`
	ParameterValues sql.NullString `gorm:"column:parameterValues" json:"parameterValues"`
}

// TableName sets the insert table name for this struct type
func (s *SystemParameter) TableName() string {
	return "system_parameter"
}
