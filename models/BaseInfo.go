package models

import "database/sql"

type BaseInfo struct {
	AddTime    sql.NullInt64  `gorm:"column:add_time" json:"add_time"`
	Aid        sql.NullInt64  `gorm:"column:aid" json:"aid"`
	Cid        sql.NullInt64  `gorm:"column:cid" json:"cid"`
	GroupID    sql.NullInt64  `gorm:"column:group_id" json:"group_id"`
	GroupName  sql.NullString `gorm:"column:group_name" json:"group_name"`
	ID         int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel      sql.NullInt64  `gorm:"column:isdel" json:"isdel"`
	Name       sql.NullString `gorm:"column:name" json:"name"`
	Remark     sql.NullString `gorm:"column:remark" json:"remark"`
	Sort       sql.NullInt64  `gorm:"column:sort" json:"sort"`
	Status     sql.NullInt64  `gorm:"column:status" json:"status"`
	Type       sql.NullInt64  `gorm:"column:type" json:"type"`
	UpdateTime sql.NullInt64  `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (b *BaseInfo) TableName() string {
	return "base_info"
}
