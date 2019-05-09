package models

import (
	"database/sql"
	"time"
)

type WebGoods struct {
	Cid             sql.NullInt64  `gorm:"column:cid" json:"cid"`
	CreateTime      time.Time      `gorm:"column:create_time" json:"create_time"`
	CreateUser      string         `gorm:"column:create_user" json:"create_user"`
	GoodsContent    string         `gorm:"column:goods_content" json:"goods_content"`
	GoodsDescribe   string         `gorm:"column:goods_describe" json:"goods_describe"`
	GoodsImage      string         `gorm:"column:goods_image" json:"goods_image"`
	GoodsImages     sql.NullString `gorm:"column:goods_images" json:"goods_images"`
	GoodsName       string         `gorm:"column:goods_name" json:"goods_name"`
	GoodsPrice      float64        `gorm:"column:goods_price" json:"goods_price"`
	GoodsThumb      sql.NullString `gorm:"column:goods_thumb" json:"goods_thumb"`
	GoodscategoryID int            `gorm:"column:goodscategory_id" json:"goodscategory_id"`
	ID              int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	UpdateTime      time.Time      `gorm:"column:update_time" json:"update_time"`
	UpdateUser      string         `gorm:"column:update_user" json:"update_user"`
}

// TableName sets the insert table name for this struct type
func (w *WebGoods) TableName() string {
	return "web_goods"
}
