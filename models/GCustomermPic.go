package models

type GCustomermPic struct {
	CID    int    `gorm:"column:c_id" json:"c_id"`
	Cid    int    `gorm:"column:cid" json:"cid"`
	ID     int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Pic    string `gorm:"column:pic" json:"pic"`
	Status int    `gorm:"column:status" json:"status"`
}

// TableName sets the insert table name for this struct type
func (g *GCustomermPic) TableName() string {
	return "g_customerm_pic"
}
