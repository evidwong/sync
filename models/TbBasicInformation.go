package models

import "database/sql"

type TbBasicInformation struct {
	AccelerationTime                      sql.NullString `gorm:"column:accelerationTime" json:"accelerationTime"`
	AnEngine                              sql.NullString `gorm:"column:anEngine" json:"anEngine"`
	AppearanceColor                       sql.NullString `gorm:"column:appearanceColor" json:"appearanceColor"`
	BatteryChargingTime                   sql.NullString `gorm:"column:batteryChargingTime" json:"batteryChargingTime"`
	BatteryWarrantyPolicy                 sql.NullString `gorm:"column:batteryWarrantyPolicy" json:"batteryWarrantyPolicy"`
	Blend                                 sql.NullString `gorm:"column:blend" json:"blend"`
	BodyType                              sql.NullString `gorm:"column:bodyType" json:"bodyType"`
	BrandName                             sql.NullString `gorm:"column:brandName" json:"brandName"`
	CarLine                               sql.NullString `gorm:"column:carLine" json:"carLine"`
	CarModel                              sql.NullString `gorm:"column:carModel" json:"carModel"`
	DynamicType                           sql.NullString `gorm:"column:dynamicType" json:"dynamicType"`
	GearboxType                           sql.NullString `gorm:"column:gearboxType" json:"gearboxType"`
	GuidePrice                            sql.NullString `gorm:"column:guidePrice" json:"guidePrice"`
	ID                                    int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Manufacturer                          sql.NullString `gorm:"column:manufacturer" json:"manufacturer"`
	MaximumPower                          sql.NullString `gorm:"column:maximumPower" json:"maximumPower"`
	ModelLevel                            sql.NullString `gorm:"column:modelLevel" json:"modelLevel"`
	NationalSubsidiesForNewEnergyVehicles sql.NullString `gorm:"column:nationalSubsidiesForNewEnergyVehicles" json:"nationalSubsidiesForNewEnergyVehicles"`
	PureElectricMaximumRange              sql.NullString `gorm:"column:pureElectricMaximumRange" json:"pureElectricMaximumRange"`
	ReferencePrice                        sql.NullString `gorm:"column:referencePrice" json:"referencePrice"`
	Standard                              sql.NullString `gorm:"column:standard" json:"standard"`
	TimeToMarket                          sql.NullString `gorm:"column:timeToMarket" json:"timeToMarket"`
	TopSpeed                              sql.NullString `gorm:"column:topSpeed" json:"topSpeed"`
	TotalMotorPower                       sql.NullString `gorm:"column:totalMotorPower" json:"totalMotorPower"`
	WarrantyPolicy                        sql.NullString `gorm:"column:warrantyPolicy" json:"warrantyPolicy"`
}

// TableName sets the insert table name for this struct type
func (t *TbBasicInformation) TableName() string {
	return "tb_basic_information"
}
