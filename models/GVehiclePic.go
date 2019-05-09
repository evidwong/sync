package models

import "database/sql"

type GVehiclePic struct {
	Cid    int           `gorm:"column:cid" json:"cid"`
	Code   string        `gorm:"column:code" json:"code"`
	ID     int           `gorm:"column:id;primary_key" json:"id;primary_key"`
	Pic    string        `gorm:"column:pic" json:"pic"`
	Status int           `gorm:"column:status" json:"status"`
	UID    sql.NullInt64 `gorm:"column:u_id" json:"u_id"`
}

// TableName sets the insert table name for this struct type
func (g *GVehiclePic) TableName() string {
	return "g_vehicle_pic"
}
