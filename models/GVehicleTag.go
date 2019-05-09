package models

import "time"

type GVehicleTag struct {
	AddTime time.Time `gorm:"column:add_time" json:"add_time"`
	CID     int       `gorm:"column:c_id" json:"c_id"`
	Cid     int       `gorm:"column:cid" json:"cid"`
	ID      int       `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name    string    `gorm:"column:name" json:"name"`
	Remark  string    `gorm:"column:remark" json:"remark"`
	Status  int       `gorm:"column:status" json:"status"`
	Tag     string    `gorm:"column:tag" json:"tag"`
}

// TableName sets the insert table name for this struct type
func (g *GVehicleTag) TableName() string {
	return "g_vehicle_tag"
}
