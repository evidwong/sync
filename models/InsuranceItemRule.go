package models

import (
	"database/sql"
	"time"
)

type InsuranceItemRule struct {
	AddTime       sql.NullInt64  `gorm:"column:add_time" json:"add_time"`
	Aid           sql.NullInt64  `gorm:"column:aid" json:"aid"`
	Auid          sql.NullInt64  `gorm:"column:auid" json:"auid"`
	Cid           sql.NullInt64  `gorm:"column:cid" json:"cid"`
	City          sql.NullInt64  `gorm:"column:city" json:"city"`
	CityName      sql.NullString `gorm:"column:cityName" json:"cityName"`
	Coid          sql.NullInt64  `gorm:"column:coid" json:"coid"`
	CompanyName   sql.NullString `gorm:"column:companyName" json:"companyName"`
	EndTime       sql.NullInt64  `gorm:"column:end_time" json:"end_time"`
	EndTimeDate   time.Time      `gorm:"column:end_time_date" json:"end_time_date"`
	ID            int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel         int            `gorm:"column:isdel" json:"isdel"`
	Item          sql.NullString `gorm:"column:item" json:"item"`
	Name          sql.NullString `gorm:"column:name" json:"name"`
	Province      sql.NullInt64  `gorm:"column:province" json:"province"`
	ProvinceName  sql.NullString `gorm:"column:provinceName" json:"provinceName"`
	Sort          int            `gorm:"column:sort" json:"sort"`
	StartTime     sql.NullInt64  `gorm:"column:start_time" json:"start_time"`
	StartTimeDate time.Time      `gorm:"column:start_time_date" json:"start_time_date"`
	Status        int            `gorm:"column:status" json:"status"`
	Storeid       sql.NullString `gorm:"column:storeid" json:"storeid"`
	UpdateTime    sql.NullInt64  `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (i *InsuranceItemRule) TableName() string {
	return "insurance_item_rule"
}
