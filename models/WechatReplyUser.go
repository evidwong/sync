package models

type WechatReplyUser struct {
	Cid     int    `gorm:"column:cid" json:"cid"`
	ID      int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Replyid int    `gorm:"column:replyid" json:"replyid"`
	Type    string `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (w *WechatReplyUser) TableName() string {
	return "wechat_reply_user"
}
