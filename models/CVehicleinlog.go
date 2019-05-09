package models

import (
	"database/sql"
	"time"
)

type CVehicleinlog struct {
	CModel          sql.NullString `gorm:"column:CModel" json:"CModel"`
	CarProcess      string         `gorm:"column:CarProcess" json:"CarProcess"`
	CustomerName    sql.NullString `gorm:"column:CustomerName" json:"CustomerName"`
	DriverHandPhone sql.NullString `gorm:"column:DriverHandPhone" json:"DriverHandPhone"`
	EngineNo        string         `gorm:"column:EngineNo" json:"EngineNo"`
	InDate          time.Time      `gorm:"column:InDate" json:"InDate"`
	InsuranceDate   time.Time      `gorm:"column:InsuranceDate" json:"InsuranceDate"`
	ReceptionSource sql.NullString `gorm:"column:ReceptionSource" json:"ReceptionSource"`
	RegisterNo      sql.NullString `gorm:"column:RegisterNo" json:"RegisterNo"`
	RelaSalesman    string         `gorm:"column:RelaSalesman" json:"RelaSalesman"`
	RunLicDate      time.Time      `gorm:"column:RunLicDate" json:"RunLicDate"`
	VehicleCode     string         `gorm:"column:VehicleCode" json:"VehicleCode"`
	YearCheckDate   time.Time      `gorm:"column:YearCheckDate" json:"YearCheckDate"`
	Camera          sql.NullString `gorm:"column:camera" json:"camera"`
	Cid             int            `gorm:"column:cid" json:"cid"`
	Comno           sql.NullString `gorm:"column:comno" json:"comno"`
	CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at"`
	ID              int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Pic             sql.NullString `gorm:"column:pic" json:"pic"`
	PicBase         sql.NullString `gorm:"column:pic_base" json:"pic_base"`
	StoreID         int            `gorm:"column:store_id" json:"store_id"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at"`
	Vin             sql.NullString `gorm:"column:vin" json:"vin"`
}

// TableName sets the insert table name for this struct type
func (c *CVehicleinlog) TableName() string {
	return "c_vehicleinlog"
}
