package models

import "database/sql"

type UserAddress struct {
	Address       string         `gorm:"column:address" json:"address"`
	AddressDetail sql.NullString `gorm:"column:address_detail" json:"address_detail"`
	Addtime       int            `gorm:"column:addtime" json:"addtime"`
	AreaCode      sql.NullString `gorm:"column:area_code" json:"area_code"`
	Cid           int            `gorm:"column:cid" json:"cid"`
	City          sql.NullString `gorm:"column:city" json:"city"`
	County        sql.NullString `gorm:"column:county" json:"county"`
	ID            int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	IsDel         int            `gorm:"column:is_del" json:"is_del"`
	Isdefault     int            `gorm:"column:isdefault" json:"isdefault"`
	Name          string         `gorm:"column:name" json:"name"`
	PostNum       sql.NullString `gorm:"column:post_num" json:"post_num"`
	Province      sql.NullString `gorm:"column:province" json:"province"`
	Tel           string         `gorm:"column:tel" json:"tel"`
	UID           int            `gorm:"column:uid" json:"uid"`
}

// TableName sets the insert table name for this struct type
func (u *UserAddress) TableName() string {
	return "user_address"
}
