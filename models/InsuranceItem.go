package models

import "database/sql"

type InsuranceItem struct {
	AddTime       sql.NullInt64  `gorm:"column:add_time" json:"add_time"`
	Aid           int            `gorm:"column:aid" json:"aid"`
	Auid          int            `gorm:"column:auid" json:"auid"`
	Cid           int            `gorm:"column:cid" json:"cid"`
	ID            int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	InsuredAmount float32        `gorm:"column:insured_amount" json:"insured_amount"`
	Isdel         int            `gorm:"column:isdel" json:"isdel"`
	Islimit       sql.NullInt64  `gorm:"column:islimit" json:"islimit"`
	ItemChannel   int            `gorm:"column:itemChannel" json:"itemChannel"`
	Itemtype      int            `gorm:"column:itemtype" json:"itemtype"`
	Level         int            `gorm:"column:level" json:"level"`
	Name          string         `gorm:"column:name" json:"name"`
	Pid           sql.NullInt64  `gorm:"column:pid" json:"pid"`
	Recommond     sql.NullInt64  `gorm:"column:recommond" json:"recommond"`
	Remark        sql.NullString `gorm:"column:remark" json:"remark"`
	Sort          sql.NullInt64  `gorm:"column:sort" json:"sort"`
	Status        sql.NullInt64  `gorm:"column:status" json:"status"`
	UpdateTime    sql.NullInt64  `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (i *InsuranceItem) TableName() string {
	return "insurance_item"
}
