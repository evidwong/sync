package models

import "database/sql"

type TbSafeConfiguration struct {
	Airbag                 sql.NullString `gorm:"column:airbag" json:"airbag"`
	AntiLockBrake          sql.NullString `gorm:"column:antiLockBrake" json:"antiLockBrake"`
	BackCentralAirBag      sql.NullString `gorm:"column:backCentralAirBag" json:"backCentralAirBag"`
	BacksideAirbag         sql.NullString `gorm:"column:backsideAirbag" json:"backsideAirbag"`
	BodyStabilityControl   sql.NullString `gorm:"column:bodyStabilityControl" json:"bodyStabilityControl"`
	BrakeAssistant         sql.NullString `gorm:"column:brakeAssistant" json:"brakeAssistant"`
	BrakeForceDistribution sql.NullString `gorm:"column:brakeForceDistribution" json:"brakeForceDistribution"`
	BrandName              sql.NullString `gorm:"column:brandName" json:"brandName"`
	CarLine                sql.NullString `gorm:"column:carLine" json:"carLine"`
	CarModel               sql.NullString `gorm:"column:carModel" json:"carModel"`
	FrontAirBag            sql.NullString `gorm:"column:frontAirBag" json:"frontAirBag"`
	ID                     int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	KneeAirbag             sql.NullString `gorm:"column:kneeAirbag" json:"kneeAirbag"`
	MainDrivingAirbag      sql.NullString `gorm:"column:mainDrivingAirbag" json:"mainDrivingAirbag"`
	Manufacturer           sql.NullString `gorm:"column:manufacturer" json:"manufacturer"`
	RearChildSeatInterface sql.NullString `gorm:"column:rearChildSeatInterface" json:"rearChildSeatInterface"`
	RunFlats               sql.NullString `gorm:"column:runFlats" json:"runFlats"`
	SafetyBeltAirbag       sql.NullString `gorm:"column:safetyBeltAirbag" json:"safetyBeltAirbag"`
	SideSafetyAirCurtain   sql.NullString `gorm:"column:sideSafetyAirCurtain" json:"sideSafetyAirCurtain"`
	TirePressureMonitor    sql.NullString `gorm:"column:tirePressureMonitor" json:"tirePressureMonitor"`
	TractionControl        sql.NullString `gorm:"column:tractionControl" json:"tractionControl"`
}

// TableName sets the insert table name for this struct type
func (t *TbSafeConfiguration) TableName() string {
	return "tb_safe_configuration"
}
