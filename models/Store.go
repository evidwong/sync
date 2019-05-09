package models

import "database/sql"

type Store struct {
	Address             string         `gorm:"column:address" json:"address"`
	AftersaleTel        sql.NullString `gorm:"column:aftersale_tel" json:"aftersale_tel"`
	Agentid             int            `gorm:"column:agentid" json:"agentid"`
	AloneSettle         sql.NullInt64  `gorm:"column:alone_settle" json:"alone_settle"`
	Bdcode              sql.NullString `gorm:"column:bdcode" json:"bdcode"`
	Branch              string         `gorm:"column:branch" json:"branch"`
	CarsellTel          sql.NullString `gorm:"column:carsell_tel" json:"carsell_tel"`
	Cid                 int            `gorm:"column:cid" json:"cid"`
	City                string         `gorm:"column:city" json:"city"`
	Comno               string         `gorm:"column:comno" json:"comno"`
	County              string         `gorm:"column:county" json:"county"`
	FitTel              sql.NullString `gorm:"column:fit_tel" json:"fit_tel"`
	GdwzID              int            `gorm:"column:gdwz_id" json:"gdwz_id"`
	HealthCard          sql.NullString `gorm:"column:health_card" json:"health_card"`
	Human               string         `gorm:"column:human" json:"human"`
	ID                  int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Image               string         `gorm:"column:image" json:"image"`
	InsureTel           sql.NullString `gorm:"column:insure_tel" json:"insure_tel"`
	Isshow              int            `gorm:"column:isshow" json:"isshow"`
	Lat                 float64        `gorm:"column:lat" json:"lat"`
	Lng                 float64        `gorm:"column:lng" json:"lng"`
	Name                string         `gorm:"column:name" json:"name"`
	Phone               string         `gorm:"column:phone" json:"phone"`
	Pid                 int            `gorm:"column:pid" json:"pid"`
	Province            string         `gorm:"column:province" json:"province"`
	Remark              sql.NullString `gorm:"column:remark" json:"remark"`
	Score               float32        `gorm:"column:score" json:"score"`
	SecondHandCarShopID sql.NullInt64  `gorm:"column:secondHandCar_shop_id" json:"secondHandCar_shop_id"`
	SecondHandCarTel    sql.NullString `gorm:"column:secondHandCar_tel" json:"secondHandCar_tel"`
	ShopPlatformConfig  sql.NullString `gorm:"column:shop_platform_config" json:"shop_platform_config"`
	ShopPlatformID      sql.NullInt64  `gorm:"column:shop_platform_id" json:"shop_platform_id"`
	ShopPlatformName    sql.NullString `gorm:"column:shop_platform_name" json:"shop_platform_name"`
	ShortName           string         `gorm:"column:short_name" json:"short_name"`
	Sort                int            `gorm:"column:sort" json:"sort"`
	Tel                 string         `gorm:"column:tel" json:"tel"`
	Thumb               string         `gorm:"column:thumb" json:"thumb"`
	Updatetime          int            `gorm:"column:updatetime" json:"updatetime"`
	Voter               int            `gorm:"column:voter" json:"voter"`
}

// TableName sets the insert table name for this struct type
func (s *Store) TableName() string {
	return "store"
}
