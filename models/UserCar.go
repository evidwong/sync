package models

type UserCar struct {
	Addtime    int    `gorm:"column:addtime" json:"addtime"`
	BuyTime    int    `gorm:"column:buy_time" json:"buy_time"`
	CarType    string `gorm:"column:car_type" json:"car_type"`
	Carid      int    `gorm:"column:carid" json:"carid"`
	Engine     string `gorm:"column:engine" json:"engine"`
	Exist      int    `gorm:"column:exist" json:"exist"`
	FFrame     string `gorm:"column:f_frame" json:"f_frame"`
	Frame      string `gorm:"column:frame" json:"frame"`
	ID         int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Mileage    int    `gorm:"column:mileage" json:"mileage"`
	Status     int    `gorm:"column:status" json:"status"`
	UID        int    `gorm:"column:uid" json:"uid"`
	Updatetime int    `gorm:"column:updatetime" json:"updatetime"`
	Vin        string `gorm:"column:vin" json:"vin"`
}

// TableName sets the insert table name for this struct type
func (u *UserCar) TableName() string {
	return "user_car"
}
