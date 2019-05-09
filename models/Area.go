package models

import "database/sql"

type Area struct {
	Areacode  sql.NullInt64  `gorm:"column:areacode" json:"areacode"`
	Areaname  string         `gorm:"column:areaname" json:"areaname"`
	ID        int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Lat       sql.NullString `gorm:"column:lat" json:"lat"`
	Level     int            `gorm:"column:level" json:"level"`
	Lng       sql.NullString `gorm:"column:lng" json:"lng"`
	Parentid  int            `gorm:"column:parentid" json:"parentid"`
	Pinyin    sql.NullString `gorm:"column:pinyin" json:"pinyin"`
	Position  string         `gorm:"column:position" json:"position"`
	Shortname sql.NullString `gorm:"column:shortname" json:"shortname"`
	Sort      sql.NullInt64  `gorm:"column:sort" json:"sort"`
	Zipcode   sql.NullInt64  `gorm:"column:zipcode" json:"zipcode"`
}

// TableName sets the insert table name for this struct type
func (a *Area) TableName() string {
	return "area"
}
