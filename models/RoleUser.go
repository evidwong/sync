package models

import "database/sql"

type RoleUser struct {
	RoleID sql.NullInt64  `gorm:"column:role_id" json:"role_id"`
	UserID sql.NullString `gorm:"column:user_id" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (r *RoleUser) TableName() string {
	return "role_user"
}
