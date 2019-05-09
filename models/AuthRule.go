package models

import "database/sql"

type AuthRule struct {
	Cls          string         `gorm:"column:cls" json:"cls"`
	Condition    string         `gorm:"column:condition" json:"condition"`
	Display      int            `gorm:"column:display" json:"display"`
	FastCid      int            `gorm:"column:fast_cid" json:"fast_cid"`
	FastClick    int            `gorm:"column:fast_click" json:"fast_click"`
	ID           int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Indexid      sql.NullInt64  `gorm:"column:indexid" json:"indexid"`
	Ismenu       int            `gorm:"column:ismenu" json:"ismenu"`
	MenuModel    sql.NullInt64  `gorm:"column:menu_model" json:"menu_model"`
	MenuPid      int            `gorm:"column:menu_pid" json:"menu_pid"`
	Name         string         `gorm:"column:name" json:"name"`
	Pid          int            `gorm:"column:pid" json:"pid"`
	RuleType     sql.NullInt64  `gorm:"column:rule_type" json:"rule_type"`
	ShortcutIcon sql.NullString `gorm:"column:shortcut_icon" json:"shortcut_icon"`
	Sort         int            `gorm:"column:sort" json:"sort"`
	Status       int            `gorm:"column:status" json:"status"`
	Title        string         `gorm:"column:title" json:"title"`
	Type         int            `gorm:"column:type" json:"type"`
	URL          string         `gorm:"column:url" json:"url"`
}

// TableName sets the insert table name for this struct type
func (a *AuthRule) TableName() string {
	return "auth_rule"
}
