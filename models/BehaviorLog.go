package models

import (
	"database/sql"
	"time"
)

type BehaviorLog struct {
	Area             sql.NullString `gorm:"column:area" json:"area"`
	Cid              sql.NullInt64  `gorm:"column:cid" json:"cid"`
	ControllerAction sql.NullString `gorm:"column:controller_action" json:"controller_action"`
	ControllerName   sql.NullString `gorm:"column:controller_name" json:"controller_name"`
	Country          sql.NullString `gorm:"column:country" json:"country"`
	Date             time.Time      `gorm:"column:date" json:"date"`
	ID               int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	IP               sql.NullString `gorm:"column:ip" json:"ip"`
	JSONData         sql.NullString `gorm:"column:json_data" json:"json_data"`
	UID              int            `gorm:"column:uid" json:"uid"`
	UserAgent        sql.NullString `gorm:"column:user_agent" json:"user_agent"`
	Username         sql.NullString `gorm:"column:username" json:"username"`
}

// TableName sets the insert table name for this struct type
func (b *BehaviorLog) TableName() string {
	return "behavior_log"
}
