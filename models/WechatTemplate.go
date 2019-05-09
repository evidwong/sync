package models

type WechatTemplate struct {
	Addtime    int    `gorm:"column:addtime" json:"addtime"`
	Cid        int    `gorm:"column:cid" json:"cid"`
	ID         int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Keyname    string `gorm:"column:keyname" json:"keyname"`
	Name       string `gorm:"column:name" json:"name"`
	Status     int    `gorm:"column:status" json:"status"`
	Templateid string `gorm:"column:templateid" json:"templateid"`
}

// TableName sets the insert table name for this struct type
func (w *WechatTemplate) TableName() string {
	return "wechat_template"
}
