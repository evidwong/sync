package models

import "database/sql"

type TbDynamicSystem struct {
	AfterTheMotorMaximumPower         sql.NullString `gorm:"column:afterTheMotorMaximumPower" json:"afterTheMotorMaximumPower"`
	BatteryCapacity                   sql.NullString `gorm:"column:batteryCapacity" json:"batteryCapacity"`
	BatteryChargingTime               sql.NullString `gorm:"column:batteryChargingTime" json:"batteryChargingTime"`
	BatteryChargingTimeIsFastCharging sql.NullString `gorm:"column:batteryChargingTimeIsFastCharging" json:"batteryChargingTimeIsFastCharging"`
	BatteryChargingTimeisSlow         sql.NullString `gorm:"column:batteryChargingTimeisSlow" json:"batteryChargingTimeisSlow"`
	BatteryQualityAssurance           sql.NullString `gorm:"column:batteryQualityAssurance" json:"batteryQualityAssurance"`
	BlockNumber                       sql.NullString `gorm:"column:blockNumber" json:"blockNumber"`
	BrandName                         sql.NullString `gorm:"column:brandName" json:"brandName"`
	CarLine                           sql.NullString `gorm:"column:carLine" json:"carLine"`
	CarModel                          sql.NullString `gorm:"column:carModel" json:"carModel"`
	Compression                       sql.NullString `gorm:"column:compression" json:"compression"`
	CylinderForm                      sql.NullString `gorm:"column:cylinderForm" json:"cylinderForm"`
	CylinderNumber                    sql.NullString `gorm:"column:cylinderNumber" json:"cylinderNumber"`
	Displacement                      sql.NullString `gorm:"column:displacement" json:"displacement"`
	EngineStartAndStop                sql.NullString `gorm:"column:engineStartAndStop" json:"engineStartAndStop"`
	FuelLabeling                      sql.NullString `gorm:"column:fuelLabeling" json:"fuelLabeling"`
	ID                                int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	IntakeForm                        sql.NullString `gorm:"column:intakeForm" json:"intakeForm"`
	Manufacturer                      sql.NullString `gorm:"column:manufacturer" json:"manufacturer"`
	MaximumEndurance                  sql.NullString `gorm:"column:maximumEndurance" json:"maximumEndurance"`
	MaximumHP                         sql.NullString `gorm:"column:maximumHP" json:"maximumHP"`
	MaximumPowerKW                    sql.NullString `gorm:"column:maximumPowerKW" json:"maximumPowerKW"`
	MaximumPowerOfFrontMotor          sql.NullString `gorm:"column:maximumPowerOfFrontMotor" json:"maximumPowerOfFrontMotor"`
	MaximumPowerSpeed                 sql.NullString `gorm:"column:maximumPowerSpeed" json:"maximumPowerSpeed"`
	MaximumTorqueOfFrontMotor         sql.NullString `gorm:"column:maximumTorqueOfFrontMotor" json:"maximumTorqueOfFrontMotor"`
	MaximumTorqueOfRearMotor          sql.NullString `gorm:"column:maximumTorqueOfRearMotor" json:"maximumTorqueOfRearMotor"`
	MaximumTorqueSpeed                sql.NullString `gorm:"column:maximumTorqueSpeed" json:"maximumTorqueSpeed"`
	NaximumTorque                     sql.NullString `gorm:"column:naximumTorque" json:"naximumTorque"`
	OilSupplyMode                     sql.NullString `gorm:"column:oilSupplyMode" json:"oilSupplyMode"`
	PowerConsumption                  sql.NullString `gorm:"column:powerConsumption" json:"powerConsumption"`
	SystemIntegratedPower             sql.NullString `gorm:"column:systemIntegratedPower" json:"systemIntegratedPower"`
	SystemIntegratedTorque            sql.NullString `gorm:"column:systemIntegratedTorque" json:"systemIntegratedTorque"`
	TotalMotorPower                   sql.NullString `gorm:"column:totalMotorPower" json:"totalMotorPower"`
	TotalMotorTorque                  sql.NullString `gorm:"column:totalMotorTorque" json:"totalMotorTorque"`
	TransmissionType                  sql.NullString `gorm:"column:transmissionType" json:"transmissionType"`
}

// TableName sets the insert table name for this struct type
func (t *TbDynamicSystem) TableName() string {
	return "tb_dynamic_system"
}
