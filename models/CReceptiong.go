package models

type CReceptiong struct {
	Cid            int    `gorm:"column:cid" json:"cid"`
	CreceptionCode string `gorm:"column:creception_code" json:"creception_code"`
	FunctionCode   string `gorm:"column:function_code" json:"function_code"`
	ID             int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Picture        string `gorm:"column:picture" json:"picture"`
	Status         int    `gorm:"column:status" json:"status"`
	Video          string `gorm:"column:video" json:"video"`
}

// TableName sets the insert table name for this struct type
func (c *CReceptiong) TableName() string {
	return "c_receptiong"
}
