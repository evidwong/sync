package models

type MaintOrder struct {
	Completedate int    `gorm:"column:completedate" json:"completedate"`
	ID           int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Indate       int    `gorm:"column:indate" json:"indate"`
	Orderid      string `gorm:"column:orderid" json:"orderid"`
	Outdate      int    `gorm:"column:outdate" json:"outdate"`
	Province     string `gorm:"column:province" json:"province"`
	Recdate      int    `gorm:"column:recdate" json:"recdate"`
	Recmoneydate int    `gorm:"column:recmoneydate" json:"recmoneydate"`
	Rectype      string `gorm:"column:rectype" json:"rectype"`
	Registerno   string `gorm:"column:registerno" json:"registerno"`
	Settledate   int    `gorm:"column:settledate" json:"settledate"`
	Status       string `gorm:"column:status" json:"status"`
	UID          int    `gorm:"column:uid" json:"uid"`
	Updatetime   int    `gorm:"column:updatetime" json:"updatetime"`
	Vehicle      string `gorm:"column:vehicle" json:"vehicle"`
}

// TableName sets the insert table name for this struct type
func (m *MaintOrder) TableName() string {
	return "maint_order"
}
