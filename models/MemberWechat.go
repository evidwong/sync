package models

import "database/sql"

type MemberWechat struct {
	Cid        int           `gorm:"column:cid" json:"cid"`
	ID         int           `gorm:"column:id;primary_key" json:"id;primary_key"`
	Openid     string        `gorm:"column:openid" json:"openid"`
	Opid       int           `gorm:"column:opid" json:"opid"`
	Status     int           `gorm:"column:status" json:"status"`
	UID        int           `gorm:"column:uid" json:"uid"`
	UpdateTime sql.NullInt64 `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (m *MemberWechat) TableName() string {
	return "member_wechat"
}
