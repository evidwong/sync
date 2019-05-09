package models

import "database/sql"

type TbBodySize struct {
	AfterTheTireSize            sql.NullString `gorm:"column:afterTheTireSize" json:"afterTheTireSize"`
	BrandName                   sql.NullString `gorm:"column:brandName" json:"brandName"`
	CarHigh                     sql.NullString `gorm:"column:carHigh" json:"carHigh"`
	CarLine                     sql.NullString `gorm:"column:carLine" json:"carLine"`
	CarLong                     sql.NullString `gorm:"column:carLong" json:"carLong"`
	CarModel                    sql.NullString `gorm:"column:carModel" json:"carModel"`
	CarWide                     sql.NullString `gorm:"column:carWide" json:"carWide"`
	FrontTireType               sql.NullString `gorm:"column:frontTireType" json:"frontTireType"`
	FullLoadMass                sql.NullString `gorm:"column:fullLoadMass" json:"fullLoadMass"`
	ID                          int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	KerbMass                    sql.NullString `gorm:"column:kerbMass" json:"kerbMass"`
	Luggage                     sql.NullString `gorm:"column:luggage" json:"luggage"`
	LuggageCompartmentExpansion sql.NullString `gorm:"column:luggageCompartmentExpansion" json:"luggageCompartmentExpansion"`
	Manufacturer                sql.NullString `gorm:"column:manufacturer" json:"manufacturer"`
	MinimumGroundClearance      sql.NullString `gorm:"column:minimumGroundClearance" json:"minimumGroundClearance"`
	MinimumTurningDiameter      sql.NullString `gorm:"column:minimumTurningDiameter" json:"minimumTurningDiameter"`
	Pedestal                    sql.NullString `gorm:"column:pedestal" json:"pedestal"`
	SpareTire                   sql.NullString `gorm:"column:spareTire" json:"spareTire"`
	TankVolume                  sql.NullString `gorm:"column:tankVolume" json:"tankVolume"`
	TireSpecification           sql.NullString `gorm:"column:tireSpecification" json:"tireSpecification"`
	Wheelbase                   sql.NullString `gorm:"column:wheelbase" json:"wheelbase"`
}

// TableName sets the insert table name for this struct type
func (t *TbBodySize) TableName() string {
	return "tb_body_size"
}
