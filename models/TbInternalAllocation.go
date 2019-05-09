package models

import "database/sql"

type TbInternalAllocation struct {
	ActiveNoise                sql.NullString `gorm:"column:activeNoise" json:"activeNoise"`
	AirPurification            sql.NullString `gorm:"column:airPurification" json:"airPurification"`
	AmbientLighting            sql.NullString `gorm:"column:ambientLighting" json:"ambientLighting"`
	BrandName                  sql.NullString `gorm:"column:brandName" json:"brandName"`
	CarLine                    sql.NullString `gorm:"column:carLine" json:"carLine"`
	CarModel                   sql.NullString `gorm:"column:carModel" json:"carModel"`
	FrontRowAirConditioner     sql.NullString `gorm:"column:frontRowAirConditioner" json:"frontRowAirConditioner"`
	HeatingWheel               sql.NullString `gorm:"column:heatingWheel" json:"heatingWheel"`
	ID                         int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	InteriorMaterial           sql.NullString `gorm:"column:interiorMaterial" json:"interiorMaterial"`
	Manufacturer               sql.NullString `gorm:"column:manufacturer" json:"manufacturer"`
	MultifunctionSteeringWheel sql.NullString `gorm:"column:multifunctionSteeringWheel" json:"multifunctionSteeringWheel"`
	RagranceSystem             sql.NullString `gorm:"column:ragranceSystem" json:"ragranceSystem"`
	RearAirConditioning        sql.NullString `gorm:"column:rearAirConditioning" json:"rearAirConditioning"`
	VehicleRefrigerator        sql.NullString `gorm:"column:vehicleRefrigerator" json:"vehicleRefrigerator"`
	VisorMirror                sql.NullString `gorm:"column:visorMirror" json:"visorMirror"`
	WheelAdjustment            sql.NullString `gorm:"column:wheelAdjustment" json:"wheelAdjustment"`
	WheelMateria               sql.NullString `gorm:"column:wheelMateria" json:"wheelMateria"`
	WheelShift                 sql.NullString `gorm:"column:wheelShift" json:"wheelShift"`
}

// TableName sets the insert table name for this struct type
func (t *TbInternalAllocation) TableName() string {
	return "tb_internal_allocation"
}
