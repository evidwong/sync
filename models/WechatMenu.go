package models

import "database/sql"

type WechatMenu struct {
	Cid        int            `gorm:"column:cid" json:"cid"`
	Circletime sql.NullString `gorm:"column:circletime" json:"circletime"`
	ID         int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isshow     int            `gorm:"column:isshow" json:"isshow"`
	Menuid     sql.NullString `gorm:"column:menuid" json:"menuid"`
	Postmenu   string         `gorm:"column:postmenu" json:"postmenu"`
	Sqlmenu    string         `gorm:"column:sqlmenu" json:"sqlmenu"`
	Status     int            `gorm:"column:status" json:"status"`
	Tagid      int            `gorm:"column:tagid" json:"tagid"`
	Title      string         `gorm:"column:title" json:"title"`
	UpdateTime int            `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (w *WechatMenu) TableName() string {
	return "wechat_menu"
}
