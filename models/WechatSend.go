package models

import "time"

type WechatSend struct {
	Cid      int       `gorm:"column:cid" json:"cid"`
	Content  string    `gorm:"column:content" json:"content"`
	ID       int       `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name     string    `gorm:"column:name" json:"name"`
	SendTime time.Time `gorm:"column:send_time" json:"send_time"`
	Status   int       `gorm:"column:status" json:"status"`
	ToHImg   string    `gorm:"column:to_h_img" json:"to_h_img"`
	ToUserid string    `gorm:"column:to_userid" json:"to_userid"`
	UserHImg string    `gorm:"column:user_h_img" json:"user_h_img"`
	UserID   string    `gorm:"column:user_id" json:"user_id"`
	UserIP   string    `gorm:"column:user_ip" json:"user_ip"`
}

// TableName sets the insert table name for this struct type
func (w *WechatSend) TableName() string {
	return "wechat_send"
}
