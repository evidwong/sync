package models

import "database/sql"

type UserStoreBind struct {
	Bdcode   sql.NullString `gorm:"column:bdcode" json:"bdcode"`
	BindTime int            `gorm:"column:bind_time" json:"bind_time"`
	Cid      int            `gorm:"column:cid" json:"cid"`
	ID       int64          `gorm:"column:id;primary_key" json:"id;primary_key"`
	Nickname sql.NullString `gorm:"column:nickname" json:"nickname"`
	Openid   string         `gorm:"column:openid" json:"openid"`
	Status   int            `gorm:"column:status" json:"status"`
	StoreID  int            `gorm:"column:store_id" json:"store_id"`
}

// TableName sets the insert table name for this struct type
func (u *UserStoreBind) TableName() string {
	return "user_store_bind"
}
