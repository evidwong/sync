package models

type Groups struct {
	Addtime int    `gorm:"column:addtime" json:"addtime"`
	ID      int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name    string `gorm:"column:name" json:"name"`
}

// TableName sets the insert table name for this struct type
func (g *Groups) TableName() string {
	return "groups"
}
