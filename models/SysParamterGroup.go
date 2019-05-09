package models

import "database/sql"

type SysParamterGroup struct {
	ParamName sql.NullString `gorm:"column:ParamName" json:"ParamName"`
	CreateAt  sql.NullInt64  `gorm:"column:create_at" json:"create_at"`
	ID        int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name      sql.NullString `gorm:"column:name" json:"name"`
	UpdateAt  sql.NullInt64  `gorm:"column:update_at" json:"update_at"`
}

// TableName sets the insert table name for this struct type
func (s *SysParamterGroup) TableName() string {
	return "sys_paramter_group"
}
