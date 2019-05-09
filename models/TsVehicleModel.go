package models

import (
	"database/sql"
	"time"
)

type TsVehicleModel struct {
	CarType           sql.NullString  `gorm:"column:car_type" json:"car_type"`
	CreateBy          sql.NullInt64   `gorm:"column:create_by" json:"create_by"`
	CreateTime        time.Time       `gorm:"column:create_time" json:"create_time"`
	ID                int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	IsDisabled        sql.NullInt64   `gorm:"column:is_disabled" json:"is_disabled"`
	IsOnSell          sql.NullInt64   `gorm:"column:is_on_sell" json:"is_on_sell"`
	LastEditBy        sql.NullInt64   `gorm:"column:last_edit_by" json:"last_edit_by"`
	LastEditTime      time.Time       `gorm:"column:last_edit_time" json:"last_edit_time"`
	Remark            sql.NullString  `gorm:"column:remark" json:"remark"`
	VehicleModelName  sql.NullString  `gorm:"column:vehicle_model_name" json:"vehicle_model_name"`
	VehicleModelPrice sql.NullFloat64 `gorm:"column:vehicle_model_price" json:"vehicle_model_price"`
	VehicleModelType  sql.NullString  `gorm:"column:vehicle_model_type" json:"vehicle_model_type"`
	VehicleModelYear  sql.NullInt64   `gorm:"column:vehicle_model_year" json:"vehicle_model_year"`
	VehicleSerialID   int             `gorm:"column:vehicle_serial_id" json:"vehicle_serial_id"`
	VehicleSerialName sql.NullString  `gorm:"column:vehicle_serial_name" json:"vehicle_serial_name"`
}

// TableName sets the insert table name for this struct type
func (t *TsVehicleModel) TableName() string {
	return "ts_vehicle_model"
}
