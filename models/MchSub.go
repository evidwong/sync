package models

import "time"

type MchSub struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	ID        int       `gorm:"column:id;primary_key" json:"id;primary_key"`
	MchID     int       `gorm:"column:mch_id" json:"mch_id"`
	Status    int       `gorm:"column:status" json:"status"`
	SubMchID  int       `gorm:"column:sub_mch_id" json:"sub_mch_id"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (m *MchSub) TableName() string {
	return "mch_sub"
}
