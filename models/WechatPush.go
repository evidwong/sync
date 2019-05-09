package models

import "database/sql"

type WechatPush struct {
	Cid       int            `gorm:"column:cid" json:"cid"`
	ID        int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Opends    string         `gorm:"column:opends" json:"opends"`
	PushID    int            `gorm:"column:push_id" json:"push_id"`
	PushLog   sql.NullString `gorm:"column:push_log" json:"push_log"`
	PushName  string         `gorm:"column:push_name" json:"push_name"`
	Status    int            `gorm:"column:status" json:"status"`
	StockTime int            `gorm:"column:stock_time" json:"stock_time"`
	Type      int            `gorm:"column:type" json:"type"`
	UserType  int            `gorm:"column:user_type" json:"user_type"`
	WrID      int            `gorm:"column:wr_id" json:"wr_id"`
	WrTitle   string         `gorm:"column:wr_title" json:"wr_title"`
}

// TableName sets the insert table name for this struct type
func (w *WechatPush) TableName() string {
	return "wechat_push"
}
