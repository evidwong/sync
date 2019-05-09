package models

import "database/sql"

type Logs struct {
	Action      string         `gorm:"column:action" json:"action"`
	ID          int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	NodeID      int            `gorm:"column:node_id" json:"node_id"`
	QueryString sql.NullString `gorm:"column:query_string" json:"query_string"`
	Time        int            `gorm:"column:time" json:"time"`
	UID         int            `gorm:"column:uid" json:"uid"`
	URL         string         `gorm:"column:url" json:"url"`
}

// TableName sets the insert table name for this struct type
func (l *Logs) TableName() string {
	return "logs"
}
