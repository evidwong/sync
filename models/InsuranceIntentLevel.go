package models

type InsuranceIntentLevel struct {
	AddTime    int    `gorm:"column:add_time" json:"add_time"`
	Aid        int    `gorm:"column:aid" json:"aid"`
	Auid       int    `gorm:"column:auid" json:"auid"`
	Cid        int    `gorm:"column:cid" json:"cid"`
	EndTime    int    `gorm:"column:end_time" json:"end_time"`
	ID         int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Intval     int    `gorm:"column:intval" json:"intval"`
	Isdel      int    `gorm:"column:isdel" json:"isdel"`
	Name       string `gorm:"column:name" json:"name"`
	StartTime  int    `gorm:"column:start_time" json:"start_time"`
	Status     int    `gorm:"column:status" json:"status"`
	Type       int    `gorm:"column:type" json:"type"`
	UpdateTime int    `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (i *InsuranceIntentLevel) TableName() string {
	return "insurance_intent_level"
}
