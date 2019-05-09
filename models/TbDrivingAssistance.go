package models

import "database/sql"

type TbDrivingAssistance struct {
	AutoParking               sql.NullString `gorm:"column:autoParking" json:"autoParking"`
	AutomaticParking          sql.NullString `gorm:"column:automaticParking" json:"automaticParking"`
	AutopilotAuxiliary        sql.NullString `gorm:"column:autopilotAuxiliary" json:"autopilotAuxiliary"`
	Auxiliary                 sql.NullString `gorm:"column:auxiliary" json:"auxiliary"`
	BackBackRadar             sql.NullString `gorm:"column:backBackRadar" json:"backBackRadar"`
	BrandName                 sql.NullString `gorm:"column:brandName" json:"brandName"`
	CarLine                   sql.NullString `gorm:"column:carLine" json:"carLine"`
	CarModel                  sql.NullString `gorm:"column:carModel" json:"carModel"`
	ChoiceOfDrivingMode       sql.NullString `gorm:"column:choiceOfDrivingMode" json:"choiceOfDrivingMode"`
	CollisionAlarm            sql.NullString `gorm:"column:collisionAlarm" json:"collisionAlarm"`
	CruiseControl             sql.NullString `gorm:"column:cruiseControl" json:"cruiseControl"`
	FatigueReminder           sql.NullString `gorm:"column:fatigueReminder" json:"fatigueReminder"`
	ForwardReversingRadar     sql.NullString `gorm:"column:forwardReversingRadar" json:"forwardReversingRadar"`
	Hdc                       sql.NullString `gorm:"column:hdc" json:"hdc"`
	HillStartAssist           sql.NullString `gorm:"column:hillStartAssist" json:"hillStartAssist"`
	ID                        int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	LaneKeeping               sql.NullString `gorm:"column:laneKeeping" json:"laneKeeping"`
	Manufacturer              sql.NullString `gorm:"column:manufacturer" json:"manufacturer"`
	NightVisionSystem         sql.NullString `gorm:"column:nightVisionSystem" json:"nightVisionSystem"`
	PilotedParking            sql.NullString `gorm:"column:pilotedParking" json:"pilotedParking"`
	ReversingImage            sql.NullString `gorm:"column:reversingImage" json:"reversingImage"`
	VariableGearRatioSteering sql.NullString `gorm:"column:variableGearRatioSteering" json:"variableGearRatioSteering"`
}

// TableName sets the insert table name for this struct type
func (t *TbDrivingAssistance) TableName() string {
	return "tb_driving_assistance"
}
