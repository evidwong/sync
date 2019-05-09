package models

type UserShortcut struct {
	Cid       int `gorm:"column:cid" json:"cid"`
	CreatedAt int `gorm:"column:created_at" json:"created_at"`
	ID        int `gorm:"column:id;primary_key" json:"id;primary_key"`
	RuleID    int `gorm:"column:rule_id" json:"rule_id"`
	Status    int `gorm:"column:status" json:"status"`
	UpdatedAt int `gorm:"column:updated_at" json:"updated_at"`
	UserID    int `gorm:"column:user_id" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (u *UserShortcut) TableName() string {
	return "user_shortcut"
}
