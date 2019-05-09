package models

import (
	"database/sql"
	"time"
)

type TsPicture struct {
	CreateBy           sql.NullInt64  `gorm:"column:create_by" json:"create_by"`
	CreateTime         time.Time      `gorm:"column:create_time" json:"create_time"`
	ID                 int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	IsDisabled         sql.NullInt64  `gorm:"column:is_disabled" json:"is_disabled"`
	LastEditBy         sql.NullInt64  `gorm:"column:last_edit_by" json:"last_edit_by"`
	LastEditTime       time.Time      `gorm:"column:last_edit_time" json:"last_edit_time"`
	PictureAlt         sql.NullString `gorm:"column:picture_alt" json:"picture_alt"`
	PictureDescription sql.NullString `gorm:"column:picture_description" json:"picture_description"`
	PictureLargeURL    sql.NullString `gorm:"column:picture_large_url" json:"picture_large_url"`
	PictureMidURL      sql.NullString `gorm:"column:picture_mid_url" json:"picture_mid_url"`
	PictureServer      sql.NullString `gorm:"column:picture_server" json:"picture_server"`
	PictureSmallURL    sql.NullString `gorm:"column:picture_small_url" json:"picture_small_url"`
	PictureTitle       sql.NullString `gorm:"column:picture_title" json:"picture_title"`
	PictureURL         sql.NullString `gorm:"column:picture_url" json:"picture_url"`
	Remark             sql.NullString `gorm:"column:remark" json:"remark"`
}

// TableName sets the insert table name for this struct type
func (t *TsPicture) TableName() string {
	return "ts_picture"
}
