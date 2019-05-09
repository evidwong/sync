package models

import "database/sql"

type AuthGroup struct {
	AddTime          int            `gorm:"column:add_time" json:"add_time"`
	Cid              int            `gorm:"column:cid" json:"cid"`
	ExtendPermission sql.NullString `gorm:"column:extend_permission" json:"extend_permission"`
	ID               int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Pid              int            `gorm:"column:pid" json:"pid"`
	Rules            sql.NullString `gorm:"column:rules" json:"rules"`
	Sort             int            `gorm:"column:sort" json:"sort"`
	Status           int            `gorm:"column:status" json:"status"`
	StoreID          sql.NullString `gorm:"column:store_id" json:"store_id"`
	Tagid            int            `gorm:"column:tagid" json:"tagid"`
	Title            string         `gorm:"column:title" json:"title"`
}

// TableName sets the insert table name for this struct type
func (a *AuthGroup) TableName() string {
	return "auth_group"
}
