package models
//AdminWechatStore ...
type AdminWechatStore struct {
	Openid string `gorm:"column:openid" json:"openid"`
	Status int    `gorm:"column:status" json:"status"`
	UID    int    `gorm:"column:uid;primary_key" json:"uid;primary_key"`
}

// TableName sets the insert table name for this struct type
func (a *AdminWechatStore) TableName() string {
	return "admin_wechat_store"
}

