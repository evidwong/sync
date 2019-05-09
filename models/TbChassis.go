package models

import "database/sql"

type TbChassis struct {
	AdjustableSuspension  sql.NullString `gorm:"column:adjustableSuspension" json:"adjustableSuspension"`
	BodyStructure         sql.NullString `gorm:"column:bodyStructure" json:"bodyStructure"`
	BrandName             sql.NullString `gorm:"column:brandName" json:"brandName"`
	CarLine               sql.NullString `gorm:"column:carLine" json:"carLine"`
	CarModel              sql.NullString `gorm:"column:carModel" json:"carModel"`
	DrivingMode           sql.NullString `gorm:"column:drivingMode" json:"drivingMode"`
	ID                    int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	LimitSlipDifferential sql.NullString `gorm:"column:limitSlipDifferential" json:"limitSlipDifferential"`
	Manufacturer          sql.NullString `gorm:"column:manufacturer" json:"manufacturer"`
	ParkingBrakeType      sql.NullString `gorm:"column:parkingBrakeType" json:"parkingBrakeType"`
	RearBrakeType         sql.NullString `gorm:"column:rearBrakeType" json:"rearBrakeType"`
	RearSuspensionType    sql.NullString `gorm:"column:rearSuspensionType" json:"rearSuspensionType"`
	SuspensionType        sql.NullString `gorm:"column:suspensionType" json:"suspensionType"`
	WheelBrakeType        sql.NullString `gorm:"column:wheelBrakeType" json:"wheelBrakeType"`
}

// TableName sets the insert table name for this struct type
func (t *TbChassis) TableName() string {
	return "tb_chassis"
}
