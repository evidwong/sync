package models

import "database/sql"

type TbInfotainment struct {
	BackRowLiquidCrystalScreen  sql.NullString `gorm:"column:backRowLiquidCrystalScreen" json:"backRowLiquidCrystalScreen"`
	Bluetooth                   sql.NullString `gorm:"column:bluetooth" json:"bluetooth"`
	Bose                        sql.NullString `gorm:"column:bose" json:"bose"`
	BrandName                   sql.NullString `gorm:"column:brandName" json:"brandName"`
	CarLine                     sql.NullString `gorm:"column:carLine" json:"carLine"`
	CarModel                    sql.NullString `gorm:"column:carModel" json:"carModel"`
	CarTV                       sql.NullString `gorm:"column:carTV" json:"carTV"`
	CdDvd                       sql.NullString `gorm:"column:cdDvd" json:"cdDvd"`
	ComputerDisplayScreen       sql.NullString `gorm:"column:computerDisplayScreen" json:"computerDisplayScreen"`
	ControlColorLcdScreen       sql.NullString `gorm:"column:controlColorLcdScreen" json:"controlColorLcdScreen"`
	ExternalInterface           sql.NullString `gorm:"column:externalInterface" json:"externalInterface"`
	GestureControlSystem        sql.NullString `gorm:"column:gestureControlSystem" json:"gestureControlSystem"`
	HUDFlatViewDisplay          sql.NullString `gorm:"column:hUDFlatViewDisplay" json:"hUDFlatViewDisplay"`
	ID                          int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	IntelligentLocation         sql.NullString `gorm:"column:intelligentLocation" json:"intelligentLocation"`
	LcdScreenSize               sql.NullString `gorm:"column:lcdScreenSize" json:"lcdScreenSize"`
	LoudspeakerNumber           sql.NullString `gorm:"column:loudspeakerNumber" json:"loudspeakerNumber"`
	Manufacturer                sql.NullString `gorm:"column:manufacturer" json:"manufacturer"`
	MobileInternet              sql.NullString `gorm:"column:mobileInternet" json:"mobileInternet"`
	MobilePhoneWirelessCharging sql.NullString `gorm:"column:mobilePhoneWirelessCharging" json:"mobilePhoneWirelessCharging"`
	NavigationGPS               sql.NullString `gorm:"column:navigationGPS" json:"navigationGPS"`
	SpeechControl               sql.NullString `gorm:"column:speechControl" json:"speechControl"`
	TheLcdPanel                 sql.NullString `gorm:"column:theLcdPanel" json:"theLcdPanel"`
	VehiclePower                sql.NullString `gorm:"column:vehiclePower" json:"vehiclePower"`
	VehicleRecorder             sql.NullString `gorm:"column:vehicleRecorder" json:"vehicleRecorder"`
}

// TableName sets the insert table name for this struct type
func (t *TbInfotainment) TableName() string {
	return "tb_infotainment"
}
