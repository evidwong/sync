package models

import "database/sql"

type Config struct {
	Addtime    int            `gorm:"column:addtime" json:"addtime"`
	Cid        int            `gorm:"column:cid" json:"cid"`
	Comno      sql.NullString `gorm:"column:comno" json:"comno"`
	ID         int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Keyname    string         `gorm:"column:keyname" json:"keyname"`
	Keyvalue   sql.NullString `gorm:"column:keyvalue" json:"keyvalue"`
	Status     int            `gorm:"column:status" json:"status"`
	StoreID    int            `gorm:"column:store_id" json:"store_id"`
	Type       int            `gorm:"column:type" json:"type"`
	Updatetime int            `gorm:"column:updatetime" json:"updatetime"`
}

// TableName sets the insert table name for this struct type
func (c *Config) TableName() string {
	return "config"
}
