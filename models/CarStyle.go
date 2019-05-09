package models

type CarStyle struct {
	Addtime  int    `gorm:"column:addtime" json:"addtime"`
	Del      int    `gorm:"column:del" json:"del"`
	ID       int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Maker    string `gorm:"column:maker" json:"maker"`
	Name     string `gorm:"column:name" json:"name"`
	SeriesID int    `gorm:"column:series_id" json:"series_id"`
	Weight   int    `gorm:"column:weight" json:"weight"`
	Year     int    `gorm:"column:year" json:"year"`
}

// TableName sets the insert table name for this struct type
func (c *CarStyle) TableName() string {
	return "car_style"
}
