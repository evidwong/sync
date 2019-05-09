package models

import "database/sql"

type AuthRuleBak struct {
	Cls       string        `gorm:"column:cls" json:"cls"`
	Condition string        `gorm:"column:condition" json:"condition"`
	Display   int           `gorm:"column:display" json:"display"`
	ID        int           `gorm:"column:id;primary_key" json:"id;primary_key"`
	Ismenu    sql.NullInt64 `gorm:"column:ismenu" json:"ismenu"`
	MenuModel sql.NullInt64 `gorm:"column:menu_model" json:"menu_model"`
	MenuPid   sql.NullInt64 `gorm:"column:menu_pid" json:"menu_pid"`
	Name      string        `gorm:"column:name" json:"name"`
	Pid       int           `gorm:"column:pid" json:"pid"`
	RuleType  sql.NullInt64 `gorm:"column:rule_type" json:"rule_type"`
	Sort      int           `gorm:"column:sort" json:"sort"`
	Status    int           `gorm:"column:status" json:"status"`
	Title     string        `gorm:"column:title" json:"title"`
	Type      int           `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (a *AuthRuleBak) TableName() string {
	return "auth_rule_bak"
}
