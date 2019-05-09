package models

import "database/sql"

type FilmQuotesList struct {
	AddTime        int             `gorm:"column:add_time" json:"add_time"`
	BrandID        int             `gorm:"column:brand_id" json:"brand_id"`
	BrandName      string          `gorm:"column:brand_name" json:"brand_name"`
	Cid            int             `gorm:"column:cid" json:"cid"`
	ColorName      sql.NullString  `gorm:"column:color_name" json:"color_name"`
	Colorid        int             `gorm:"column:colorid" json:"colorid"`
	Difficulty     int             `gorm:"column:difficulty" json:"difficulty"`
	GoodsBrand     sql.NullInt64   `gorm:"column:goods_brand" json:"goods_brand"`
	GoodsBrandName sql.NullString  `gorm:"column:goods_brand_name" json:"goods_brand_name"`
	GoodsID        int             `gorm:"column:goods_id" json:"goods_id"`
	GoodsName      string          `gorm:"column:goods_name" json:"goods_name"`
	ID             int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	MarketPrice    sql.NullFloat64 `gorm:"column:market_price" json:"market_price"`
	ModelID        int             `gorm:"column:model_id" json:"model_id"`
	ModelName      string          `gorm:"column:model_name" json:"model_name"`
	NeedTime       float32         `gorm:"column:need_time" json:"need_time"`
	Num            int             `gorm:"column:num" json:"num"`
	Place          int             `gorm:"column:place" json:"place"`
	PlaceAct       string          `gorm:"column:place_act" json:"place_act"`
	PlaceName      string          `gorm:"column:place_name" json:"place_name"`
	Price          float32         `gorm:"column:price" json:"price"`
	PriceType      int             `gorm:"column:price_type" json:"price_type"`
	PriceTypeName  string          `gorm:"column:price_type_name" json:"price_type_name"`
	Remark         string          `gorm:"column:remark" json:"remark"`
	Seat           int             `gorm:"column:seat" json:"seat"`
	SeatName       string          `gorm:"column:seat_name" json:"seat_name"`
	SeriesID       int             `gorm:"column:series_id" json:"series_id"`
	SeriesName     string          `gorm:"column:series_name" json:"series_name"`
	UpdateTime     int             `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (f *FilmQuotesList) TableName() string {
	return "film_quotes_list"
}
