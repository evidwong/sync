package models

import (
	"database/sql"
	"time"
)

type TsVehicleSerial struct {
	CreateBy            sql.NullInt64  `gorm:"column:create_by" json:"create_by"`
	CreateTime          time.Time      `gorm:"column:create_time" json:"create_time"`
	ID                  int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	IsDisabled          sql.NullInt64  `gorm:"column:is_disabled" json:"is_disabled"`
	LastEditBy          sql.NullInt64  `gorm:"column:last_edit_by" json:"last_edit_by"`
	LastEditTime        time.Time      `gorm:"column:last_edit_time" json:"last_edit_time"`
	Remark              sql.NullString `gorm:"column:remark" json:"remark"`
	VehicleBrandID      int            `gorm:"column:vehicle_brand_id" json:"vehicle_brand_id"`
	VehicleBrandName    sql.NullString `gorm:"column:vehicle_brand_name" json:"vehicle_brand_name"`
	VehicleBrandNameSub sql.NullString `gorm:"column:vehicle_brand_name_sub" json:"vehicle_brand_name_sub"`
	VehicleSerialName   sql.NullString `gorm:"column:vehicle_serial_name" json:"vehicle_serial_name"`
}

// TableName sets the insert table name for this struct type
func (t *TsVehicleSerial) TableName() string {
	return "ts_vehicle_serial"
}
