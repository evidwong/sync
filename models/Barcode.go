package models

import "database/sql"

type Barcode struct {
	BarCodeFontSize sql.NullInt64   `gorm:"column:barCode_fontSize" json:"barCode_fontSize"`
	BarCodeHeight   sql.NullInt64   `gorm:"column:barCode_height" json:"barCode_height"`
	BarCodeMargin   sql.NullInt64   `gorm:"column:barCode_margin" json:"barCode_margin"`
	BarCodeStep     sql.NullInt64   `gorm:"column:barCode_step" json:"barCode_step"`
	BarCodeWidth    sql.NullInt64   `gorm:"column:barCode_width" json:"barCode_width"`
	Cid             sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Comno           sql.NullString  `gorm:"column:comno" json:"comno"`
	ID              int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	LongRatio       sql.NullFloat64 `gorm:"column:long_ratio" json:"long_ratio"`
	PaperWidth      sql.NullInt64   `gorm:"column:paper_width" json:"paper_width"`
	TitleFontSize   sql.NullInt64   `gorm:"column:title_fontSize" json:"title_fontSize"`
	TitleHeight     sql.NullInt64   `gorm:"column:title_height" json:"title_height"`
	TitleLineHeight sql.NullInt64   `gorm:"column:title_lineHeight" json:"title_lineHeight"`
	WidthRatio      sql.NullFloat64 `gorm:"column:width_ratio" json:"width_ratio"`
}

// TableName sets the insert table name for this struct type
func (b *Barcode) TableName() string {
	return "barcode"
}
