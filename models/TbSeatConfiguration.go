package models

import "database/sql"

type TbSeatConfiguration struct {
	BackCentralArmrest              sql.NullString `gorm:"column:backCentralArmrest" json:"backCentralArmrest"`
	BackSeatFunction                sql.NullString `gorm:"column:backSeatFunction" json:"backSeatFunction"`
	BrandName                       sql.NullString `gorm:"column:brandName" json:"brandName"`
	CarLine                         sql.NullString `gorm:"column:carLine" json:"carLine"`
	CarModel                        sql.NullString `gorm:"column:carModel" json:"carModel"`
	FrontArmrest                    sql.NullString `gorm:"column:frontArmrest" json:"frontArmrest"`
	FrontSeatFunction               sql.NullString `gorm:"column:frontSeatFunction" json:"frontSeatFunction"`
	ID                              int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	MainSeatAdjustment              sql.NullString `gorm:"column:mainSeatAdjustment" json:"mainSeatAdjustment"`
	MainSeatElectricControl         sql.NullString `gorm:"column:mainSeatElectricControl" json:"mainSeatElectricControl"`
	Manufacturer                    sql.NullString `gorm:"column:manufacturer" json:"manufacturer"`
	RearCupHolder                   sql.NullString `gorm:"column:rearCupHolder" json:"rearCupHolder"`
	RearFoldingTable                sql.NullString `gorm:"column:rearFoldingTable" json:"rearFoldingTable"`
	SeatMaterial                    sql.NullString `gorm:"column:seatMaterial" json:"seatMaterial"`
	SeatingWay                      sql.NullString `gorm:"column:seatingWay" json:"seatingWay"`
	SecondRowSeatAdjustmentMode     sql.NullString `gorm:"column:secondRowSeatAdjustmentMode" json:"secondRowSeatAdjustmentMode"`
	SecondRowSeatElectricRegulation sql.NullString `gorm:"column:secondRowSeatElectricRegulation" json:"secondRowSeatElectricRegulation"`
	SportsStyleSeat                 sql.NullString `gorm:"column:sportsStyleSeat" json:"sportsStyleSeat"`
	ThirdRowsOfSeats                sql.NullString `gorm:"column:thirdRowsOfSeats" json:"thirdRowsOfSeats"`
	ViceChairAdjustingWay           sql.NullString `gorm:"column:viceChairAdjustingWay" json:"viceChairAdjustingWay"`
	ViceChairElectricControl        sql.NullString `gorm:"column:viceChairElectricControl" json:"viceChairElectricControl"`
}

// TableName sets the insert table name for this struct type
func (t *TbSeatConfiguration) TableName() string {
	return "tb_seat_configuration"
}
