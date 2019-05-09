package models

import "time"

type CarBrand struct {
	Addtime   int       `gorm:"column:addtime" json:"addtime"`
	Alpha     string    `gorm:"column:alpha" json:"alpha"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	ID        int       `gorm:"column:id;primary_key" json:"id;primary_key"`
	Image     string    `gorm:"column:image" json:"image"`
	Name      string    `gorm:"column:name" json:"name"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Weight    int       `gorm:"column:weight" json:"weight"`
}

// TableName sets the insert table name for this struct type
func (c *CarBrand) TableName() string {
	return "car_brand"
}
