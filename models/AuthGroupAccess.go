package models

type AuthGroupAccess struct {
	GroupID int `gorm:"column:group_id;primary_key" json:"group_id;primary_key"`
	UID     int `gorm:"column:uid;primary_key" json:"uid;primary_key"`
}

// TableName sets the insert table name for this struct type
func (a *AuthGroupAccess) TableName() string {
	return "auth_group_access"
}
