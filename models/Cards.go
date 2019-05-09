package models

import (
	"database/sql"
	"time"
)

type Cards struct {
	CardPics  sql.NullString `gorm:"column:card_pics" json:"card_pics"`
	Cid       sql.NullInt64  `gorm:"column:cid" json:"cid"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	ID        int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name      sql.NullString `gorm:"column:name" json:"name"`
	No        sql.NullString `gorm:"column:no" json:"no"`
	Phone     sql.NullString `gorm:"column:phone" json:"phone"`
	Type      sql.NullString `gorm:"column:type" json:"type"`
	UID       sql.NullInt64  `gorm:"column:uid" json:"uid"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (c *Cards) TableName() string {
	return "cards"
}
