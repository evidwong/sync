package models

import "database/sql"

type TbExternalAllocation struct {
	AutomaticAntiGlareOfTheRearviewMirror sql.NullString `gorm:"column:automaticAntiGlareOfTheRearviewMirror" json:"automaticAntiGlareOfTheRearviewMirror"`
	AutomaticAntiGlareRearviewMirror      sql.NullString `gorm:"column:automaticAntiGlareRearviewMirror" json:"automaticAntiGlareRearviewMirror"`
	AutomaticHeadlight                    sql.NullString `gorm:"column:automaticHeadlight" json:"automaticHeadlight"`
	BrandName                             sql.NullString `gorm:"column:brandName" json:"brandName"`
	CarLine                               sql.NullString `gorm:"column:carLine" json:"carLine"`
	CarModel                              sql.NullString `gorm:"column:carModel" json:"carModel"`
	ControlLock                           sql.NullString `gorm:"column:controlLock" json:"controlLock"`
	ElectricLuggage                       sql.NullString `gorm:"column:electricLuggage" json:"electricLuggage"`
	ElectricSideslip                      sql.NullString `gorm:"column:electricSideslip" json:"electricSideslip"`
	ElectricSuctionDoor                   sql.NullString `gorm:"column:electricSuctionDoor" json:"electricSuctionDoor"`
	FrontElectricWindow                   sql.NullString `gorm:"column:frontElectricWindow" json:"frontElectricWindow"`
	FrontFogLamp                          sql.NullString `gorm:"column:frontFogLamp" json:"frontFogLamp"`
	FrontWiper                            sql.NullString `gorm:"column:frontWiper" json:"frontWiper"`
	Headlight                             sql.NullString `gorm:"column:headlight" json:"headlight"`
	HeadlightFunction                     sql.NullString `gorm:"column:headlightFunction" json:"headlightFunction"`
	ID                                    int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	IntelligentKey                        sql.NullString `gorm:"column:intelligentKey" json:"intelligentKey"`
	LedDaytimeTrafficLights               sql.NullString `gorm:"column:ledDaytimeTrafficLights" json:"ledDaytimeTrafficLights"`
	Manufacturer                          sql.NullString `gorm:"column:manufacturer" json:"manufacturer"`
	PrivacyGlass                          sql.NullString `gorm:"column:privacyGlass" json:"privacyGlass"`
	RearElectricWindow                    sql.NullString `gorm:"column:rearElectricWindow" json:"rearElectricWindow"`
	RearShade                             sql.NullString `gorm:"column:rearShade" json:"rearShade"`
	RearSideShade                         sql.NullString `gorm:"column:rearSideShade" json:"rearSideShade"`
	RearWiper                             sql.NullString `gorm:"column:rearWiper" json:"rearWiper"`
	RearviewMirrorElectricControl         sql.NullString `gorm:"column:rearviewMirrorElectricControl" json:"rearviewMirrorElectricControl"`
	RemoteControlFunction                 sql.NullString `gorm:"column:remoteControlFunction" json:"remoteControlFunction"`
	RoofLuggageRack                       sql.NullString `gorm:"column:roofLuggageRack" json:"roofLuggageRack"`
	SkylightType                          sql.NullString `gorm:"column:skylightType" json:"skylightType"`
	SportsAppearanceSuite                 sql.NullString `gorm:"column:sportsAppearanceSuite" json:"sportsAppearanceSuite"`
	StreamingMediaMirror                  sql.NullString `gorm:"column:streamingMediaMirror" json:"streamingMediaMirror"`
	Tail                                  sql.NullString `gorm:"column:tail" json:"tail"`
}

// TableName sets the insert table name for this struct type
func (t *TbExternalAllocation) TableName() string {
	return "tb_external_allocation"
}
