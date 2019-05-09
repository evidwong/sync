package models

import (
	"database/sql"
	"time"
)

type SReserve struct {
	Code            string         `gorm:"column:Code" json:"Code"`
	FunctionCode    string         `gorm:"column:FunctionCode" json:"FunctionCode"`
	AddTime         time.Time      `gorm:"column:add_time" json:"add_time"`
	Cid             sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Cmodel          sql.NullString `gorm:"column:cmodel" json:"cmodel"`
	ComeFrom        sql.NullString `gorm:"column:comeFrom" json:"comeFrom"`
	Comno           sql.NullString `gorm:"column:comno" json:"comno"`
	CustomerName    sql.NullString `gorm:"column:customerName" json:"customerName"`
	HandPhone       sql.NullString `gorm:"column:handPhone" json:"handPhone"`
	ID              int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Itemid          sql.NullString `gorm:"column:itemid" json:"itemid"`
	Pics            sql.NullString `gorm:"column:pics" json:"pics"`
	RegisterNo      sql.NullString `gorm:"column:registerNo" json:"registerNo"`
	Remark          sql.NullString `gorm:"column:remark" json:"remark"`
	ReserveTime     time.Time      `gorm:"column:reserveTime" json:"reserveTime"`
	ServiceAdviser  sql.NullString `gorm:"column:serviceAdviser" json:"serviceAdviser"`
	ServiceItems    sql.NullString `gorm:"column:serviceItems" json:"serviceItems"`
	ServicePersonID sql.NullInt64  `gorm:"column:servicePersonId" json:"servicePersonId"`
	Status          sql.NullInt64  `gorm:"column:status" json:"status"`
	StatusTxt       sql.NullString `gorm:"column:statusTxt" json:"statusTxt"`
	StoreID         sql.NullInt64  `gorm:"column:storeId" json:"storeId"`
	UpdateTime      time.Time      `gorm:"column:update_time" json:"update_time"`
	WechatNotify    int            `gorm:"column:wechat_notify" json:"wechat_notify"`
}

// TableName sets the insert table name for this struct type
func (s *SReserve) TableName() string {
	return "s_reserve"
}
