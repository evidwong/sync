package models

import (
	"database/sql"
	"time"
)
//TsShop ...
type TsShop struct {
	Cid             sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Contact         sql.NullString `gorm:"column:contact" json:"contact"`
	ContactPhone    sql.NullString `gorm:"column:contact_phone" json:"contact_phone"`
	CreateBy        sql.NullInt64  `gorm:"column:create_by" json:"create_by"`
	CreateTime      time.Time      `gorm:"column:create_time" json:"create_time"`
	FullAddress     sql.NullString `gorm:"column:full_address" json:"full_address"`
	ID              int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	IsShow          sql.NullInt64  `gorm:"column:isShow" json:"isShow"`
	IsDisabled      sql.NullInt64  `gorm:"column:is_disabled" json:"is_disabled"`
	LastEditBy      sql.NullInt64  `gorm:"column:last_edit_by" json:"last_edit_by"`
	LastEditTime    time.Time      `gorm:"column:last_edit_time" json:"last_edit_time"`
	OrgnizationID   int            `gorm:"column:orgnization_id" json:"orgnization_id"`
	Remark          sql.NullString `gorm:"column:remark" json:"remark"`
	ShopAddress     sql.NullString `gorm:"column:shop_address" json:"shop_address"`
	ShopArea        sql.NullInt64  `gorm:"column:shop_area" json:"shop_area"`
	ShopCode        string         `gorm:"column:shop_code" json:"shop_code"`
	ShopLogo        sql.NullString `gorm:"column:shop_logo" json:"shop_logo"`
	ShopName        string         `gorm:"column:shop_name" json:"shop_name"`
	ShopState       string         `gorm:"column:shop_state" json:"shop_state"`
	ShopUserAccount string         `gorm:"column:shop_user_account" json:"shop_user_account"`
}

// TableName sets the insert table name for this struct type
func (t *TsShop) TableName() string {
	return "ts_shop"
}
