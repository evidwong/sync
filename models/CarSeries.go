package models

type CarSeries struct {
	Addtime int    `gorm:"column:addtime" json:"addtime"`
	Alpha   string `gorm:"column:alpha" json:"alpha"`
	BrandID int    `gorm:"column:brand_id" json:"brand_id"`
	ID      int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Image   string `gorm:"column:image" json:"image"`
	Name    string `gorm:"column:name" json:"name"`
	Weight  int    `gorm:"column:weight" json:"weight"`
}

// TableName sets the insert table name for this struct type
func (c *CarSeries) TableName() string {
	return "car_series"
}
