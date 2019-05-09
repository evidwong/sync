package models

type CMembersetsalesmAndVehicle struct {
	RegisterNo  string `gorm:"column:register_no" json:"register_no"`
	SetsaleID   int    `gorm:"column:setsale_id" json:"setsale_id"`
	VehicleCode string `gorm:"column:vehicle_code" json:"vehicle_code"`
}

// TableName sets the insert table name for this struct type
func (c *CMembersetsalesmAndVehicle) TableName() string {
	return "c_membersetsalesm_and_vehicle"
}
