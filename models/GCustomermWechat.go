package models

type GCustomermWechat struct {
	CID      int    `gorm:"column:c_id" json:"c_id"`
	Cid      int    `gorm:"column:cid" json:"cid"`
	ID       int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Nickname string `gorm:"column:nickname" json:"nickname"`
	Openid   string `gorm:"column:openid" json:"openid"`
	Status   int    `gorm:"column:status" json:"status"`
}

// TableName sets the insert table name for this struct type
func (g *GCustomermWechat) TableName() string {
	return "g_customerm_wechat"
}
