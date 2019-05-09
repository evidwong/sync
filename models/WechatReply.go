package models

type WechatReply struct {
	Cid     int    `gorm:"column:cid" json:"cid"`
	ID      int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel   int    `gorm:"column:isdel" json:"isdel"`
	MediaID string `gorm:"column:media_id" json:"media_id"`
	Reply   string `gorm:"column:reply" json:"reply"`
	Status  int    `gorm:"column:status" json:"status"`
	Time    int    `gorm:"column:time" json:"time"`
	TimeStr string `gorm:"column:time_str" json:"time_str"`
	Title   string `gorm:"column:title" json:"title"`
	Type    string `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (w *WechatReply) TableName() string {
	return "wechat_reply"
}
