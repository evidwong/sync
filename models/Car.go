package models

import "database/sql"

type Car struct {
	AddTime           int            `gorm:"column:add_time" json:"add_time"`
	BuyTime           int            `gorm:"column:buy_time" json:"buy_time"`
	CarBrand          int            `gorm:"column:car_brand" json:"car_brand"`
	CarColor          int            `gorm:"column:car_color" json:"car_color"`
	CarModel          sql.NullInt64  `gorm:"column:car_model" json:"car_model"`
	CarSeries         int            `gorm:"column:car_series" json:"car_series"`
	CarType           string         `gorm:"column:car_type" json:"car_type"`
	Driverhandphone   sql.NullString `gorm:"column:driverhandphone" json:"driverhandphone"`
	Driverm           sql.NullString `gorm:"column:driverm" json:"driverm"`
	Engine            sql.NullString `gorm:"column:engine" json:"engine"`
	EngineVerified    int            `gorm:"column:engine__verified" json:"engine__verified"`
	FFrame            sql.NullString `gorm:"column:f_frame" json:"f_frame"`
	Factory           sql.NullString `gorm:"column:factory" json:"factory"`
	Frame             sql.NullString `gorm:"column:frame" json:"frame"`
	FrameVerified     int            `gorm:"column:frame_verified" json:"frame_verified"`
	ID                int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Iscom             int            `gorm:"column:iscom" json:"iscom"`
	Isowner           int            `gorm:"column:isowner" json:"isowner"`
	Mileage           sql.NullInt64  `gorm:"column:mileage" json:"mileage"`
	Province          string         `gorm:"column:province" json:"province"`
	Status            int            `gorm:"column:status" json:"status"`
	Surplus           sql.NullString `gorm:"column:surplus" json:"surplus"`
	Travel            sql.NullString `gorm:"column:travel" json:"travel"`
	UID               int            `gorm:"column:uid" json:"uid"`
	UpdateTime        int            `gorm:"column:update_time" json:"update_time"`
	ValidTime         int            `gorm:"column:valid_time" json:"valid_time"`
	Vehicle           string         `gorm:"column:vehicle" json:"vehicle"`
	Vehiclecode       sql.NullString `gorm:"column:vehiclecode" json:"vehiclecode"`
	Vin               sql.NullString `gorm:"column:vin" json:"vin"`
	ViolateUpdateTime int            `gorm:"column:violate_update_time" json:"violate_update_time"`
	Yearcheckdate     sql.NullInt64  `gorm:"column:yearcheckdate" json:"yearcheckdate"`
}

// TableName sets the insert table name for this struct type
func (c *Car) TableName() string {
	return "car"
}
