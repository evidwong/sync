package models

type OpenUserGroup struct {
	Cid        int    `gorm:"column:cid" json:"cid"`
	ID         int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name       string `gorm:"column:name" json:"name"`
	Permission string `gorm:"column:permission" json:"permission"`
	Status     int    `gorm:"column:status" json:"status"`
	Updatetime int    `gorm:"column:updatetime" json:"updatetime"`
}

// TableName sets the insert table name for this struct type
func (o *OpenUserGroup) TableName() string {
	return "open_user_group"
}
