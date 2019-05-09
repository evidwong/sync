package models

type WechatRule struct {
	ID        int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Keyid     int    `gorm:"column:keyid" json:"keyid"`
	ReplyType string `gorm:"column:reply_type" json:"reply_type"`
	Time      int    `gorm:"column:time" json:"time"`
}

// TableName sets the insert table name for this struct type
func (w *WechatRule) TableName() string {
	return "wechat_rule"
}
