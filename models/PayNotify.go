package models

type PayNotify struct {
	AddTime  int    `gorm:"column:add_time" json:"add_time"`
	Datas    string `gorm:"column:datas" json:"datas"`
	FromType int    `gorm:"column:from_type" json:"from_type"`
	ID       int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Unid     string `gorm:"column:unid" json:"unid"`
}

// TableName sets the insert table name for this struct type
func (p *PayNotify) TableName() string {
	return "pay_notify"
}
