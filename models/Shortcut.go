package models

type Shortcut struct {
	Cid        int    `gorm:"column:cid" json:"cid"`
	ID         int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	RoleID     int    `gorm:"column:role_id" json:"role_id"`
	Title      string `gorm:"column:title" json:"title"`
	UpdateTime int    `gorm:"column:update_time" json:"update_time"`
	URL        string `gorm:"column:url" json:"url"`
	UserID     int    `gorm:"column:user_id" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (s *Shortcut) TableName() string {
	return "shortcut"
}
