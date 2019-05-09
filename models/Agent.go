package models

import "database/sql"

type Agent struct {
	Addr             string         `gorm:"column:addr" json:"addr"`
	Addtime          sql.NullInt64  `gorm:"column:addtime" json:"addtime"`
	Commission       float32        `gorm:"column:commission" json:"commission"`
	Email            string         `gorm:"column:email" json:"email"`
	ID               int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Intro            sql.NullString `gorm:"column:intro" json:"intro"`
	IsDefault        int            `gorm:"column:is_default" json:"is_default"`
	Name             string         `gorm:"column:name" json:"name"`
	OrtherCommission float32        `gorm:"column:orther_commission" json:"orther_commission"`
	Pic              string         `gorm:"column:pic" json:"pic"`
	Qrcode           sql.NullInt64  `gorm:"column:qrcode" json:"qrcode"`
	Tel              string         `gorm:"column:tel" json:"tel"`
	UID              int            `gorm:"column:uid" json:"uid"`
	Updatetime       sql.NullInt64  `gorm:"column:updatetime" json:"updatetime"`
}

// TableName sets the insert table name for this struct type
func (a *Agent) TableName() string {
	return "agent"
}
