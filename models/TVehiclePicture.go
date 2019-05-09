package models

import "database/sql"

type TVehiclePicture struct {
	ID              int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	IsDisabled      sql.NullInt64  `gorm:"column:is_disabled" json:"is_disabled"`
	PictureCategory sql.NullString `gorm:"column:picture_category" json:"picture_category"`
	PictureID       int            `gorm:"column:picture_id" json:"picture_id"`
	SequenceNumber  sql.NullInt64  `gorm:"column:sequence_number" json:"sequence_number"`
	VehicleID       int            `gorm:"column:vehicle_id" json:"vehicle_id"`
}

// TableName sets the insert table name for this struct type
func (t *TVehiclePicture) TableName() string {
	return "t_vehicle_picture"
}
