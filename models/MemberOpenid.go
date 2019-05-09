package models

type MemberOpenid struct {
	Cid    int    `gorm:"column:cid" json:"cid"`
	Openid string `gorm:"column:openid" json:"openid"`
	Phone  string `gorm:"column:phone" json:"phone"`
}

// TableName sets the insert table name for this struct type
func (m *MemberOpenid) TableName() string {
	return "member_openid"
}
