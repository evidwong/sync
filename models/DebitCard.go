package models

type DebitCard struct {
	Amount       float32 `gorm:"column:amount" json:"amount"`
	Balance      float32 `gorm:"column:balance" json:"balance"`
	Cid          int     `gorm:"column:cid" json:"cid"`
	Customercode string  `gorm:"column:customercode" json:"customercode"`
	Customername string  `gorm:"column:customername" json:"customername"`
	ID           int     `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name         string  `gorm:"column:name" json:"name"`
	Registerno   string  `gorm:"column:registerno" json:"registerno"`
	Status       string  `gorm:"column:status" json:"status"`
	UID          int     `gorm:"column:uid" json:"uid"`
	Valuecardno  string  `gorm:"column:valuecardno" json:"valuecardno"`
}

// TableName sets the insert table name for this struct type
func (d *DebitCard) TableName() string {
	return "debit_card"
}
