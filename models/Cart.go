package models

type Cart struct {
	Cid        int     `gorm:"column:cid" json:"cid"`
	ID         int     `gorm:"column:id;primary_key" json:"id;primary_key"`
	Item       string  `gorm:"column:item" json:"item"`
	Itemid     int     `gorm:"column:itemid" json:"itemid"`
	Num        int     `gorm:"column:num" json:"num"`
	Ordertime  int     `gorm:"column:ordertime" json:"ordertime"`
	Price      float32 `gorm:"column:price" json:"price"`
	UID        int     `gorm:"column:uid" json:"uid"`
	Updatetime int     `gorm:"column:updatetime" json:"updatetime"`
}

// TableName sets the insert table name for this struct type
func (c *Cart) TableName() string {
	return "cart"
}
