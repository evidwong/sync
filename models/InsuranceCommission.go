package models

type InsuranceCommission struct {
	AddTime       int     `gorm:"column:add_time" json:"add_time"`
	Coid          int     `gorm:"column:coid" json:"coid"`
	EndTime       int     `gorm:"column:end_time" json:"end_time"`
	Itemid        int     `gorm:"column:itemid" json:"itemid"`
	RebatePercent float32 `gorm:"column:rebate_percent" json:"rebate_percent"`
	Sid           int     `gorm:"column:sid" json:"sid"`
	StartTime     int     `gorm:"column:start_time" json:"start_time"`
	UpdateTime    int     `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (i *InsuranceCommission) TableName() string {
	return "insurance_commission"
}
