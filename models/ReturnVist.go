package models

type ReturnVist struct {
	Flag       int    `gorm:"column:flag" json:"flag"`
	ID         int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Orderid    string `gorm:"column:orderid" json:"orderid"`
	UID        int    `gorm:"column:uid" json:"uid"`
	Updatetime int    `gorm:"column:updatetime" json:"updatetime"`
}

// TableName sets the insert table name for this struct type
func (r *ReturnVist) TableName() string {
	return "return_vist"
}
