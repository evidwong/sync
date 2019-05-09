package models

type CarMember struct {
	Addtime        int     `gorm:"column:addtime" json:"addtime"`
	Comno          string  `gorm:"column:comno" json:"comno"`
	Customercode   string  `gorm:"column:customercode" json:"customercode"`
	Customername   string  `gorm:"column:customername" json:"customername"`
	ID             int     `gorm:"column:id;primary_key" json:"id;primary_key"`
	Integral       float32 `gorm:"column:integral" json:"integral"`
	Joinmemberdate int     `gorm:"column:joinmemberdate" json:"joinmemberdate"`
	Limitdate      int     `gorm:"column:limitdate" json:"limitdate"`
	Membercardno   int     `gorm:"column:membercardno" json:"membercardno"`
	Membercode     string  `gorm:"column:membercode" json:"membercode"`
	Memberyears    int     `gorm:"column:memberyears" json:"memberyears"`
	Registerno     string  `gorm:"column:registerno" json:"registerno"`
	Status         string  `gorm:"column:status" json:"status"`
}

// TableName sets the insert table name for this struct type
func (c *CarMember) TableName() string {
	return "car_member"
}
