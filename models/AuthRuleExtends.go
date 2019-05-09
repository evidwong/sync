package models

import "database/sql"

type AuthRuleExtends struct {
	ActionName string         `gorm:"column:action_name" json:"action_name"`
	AddTime    sql.NullInt64  `gorm:"column:add_time" json:"add_time"`
	ColumnName string         `gorm:"column:column_name" json:"column_name"`
	ID         int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	ModelName  sql.NullString `gorm:"column:model_name" json:"model_name"`
	TableNames  string         `gorm:"column:table_name" json:"table_name"`
	Title      sql.NullString `gorm:"column:title" json:"title"`
}

// TableName sets the insert table name for this struct type
func (a *AuthRuleExtends) TableName() string {
	return "auth_rule_extends"
}
