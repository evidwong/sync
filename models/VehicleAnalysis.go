package models

import (
	"database/sql"
	"time"
)

type VehicleAnalysis struct {
	Brand              sql.NullString  `gorm:"column:Brand" json:"Brand"`
	CModel             sql.NullString  `gorm:"column:CModel" json:"CModel"`
	CarType            sql.NullString  `gorm:"column:CarType" json:"CarType"`
	CreateDate         time.Time       `gorm:"column:CreateDate" json:"CreateDate"`
	CustomerCode       sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName       sql.NullString  `gorm:"column:CustomerName" json:"CustomerName"`
	DriverHandPhone    sql.NullString  `gorm:"column:DriverHandPhone" json:"DriverHandPhone"`
	FirstInFactoryDate time.Time       `gorm:"column:FirstInFactoryDate" json:"FirstInFactoryDate"`
	Insurance2Date     time.Time       `gorm:"column:Insurance2Date" json:"Insurance2Date"`
	InsuranceDate      time.Time       `gorm:"column:InsuranceDate" json:"InsuranceDate"`
	LastInFactoryDate  time.Time       `gorm:"column:LastInFactoryDate" json:"LastInFactoryDate"`
	Linkman            sql.NullString  `gorm:"column:Linkman" json:"Linkman"`
	MemberType         sql.NullString  `gorm:"column:MemberType" json:"MemberType"`
	ModifyDate         time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	PhoneTime          sql.NullString  `gorm:"column:PhoneTime" json:"PhoneTime"`
	RegisterNo         string          `gorm:"column:RegisterNo" json:"RegisterNo"`
	RunLicDate         time.Time       `gorm:"column:RunLicDate" json:"RunLicDate"`
	VehSalesName       sql.NullString  `gorm:"column:VehSalesName" json:"VehSalesName"`
	VehicleCode        sql.NullString  `gorm:"column:VehicleCode" json:"VehicleCode"`
	VehicleLevel       string          `gorm:"column:VehicleLevel" json:"VehicleLevel"`
	Vin                sql.NullString  `gorm:"column:Vin" json:"Vin"`
	WXHistoryAmount    sql.NullFloat64 `gorm:"column:WXHistoryAmount" json:"WXHistoryAmount"`
	WXHistoryCount     sql.NullInt64   `gorm:"column:WXHistoryCount" json:"WXHistoryCount"`
	YearCheckDate      time.Time       `gorm:"column:YearCheckDate" json:"YearCheckDate"`
	Cid                sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ClerkName          sql.NullString  `gorm:"column:clerk_name" json:"clerk_name"`
	Comno              sql.NullString  `gorm:"column:comno" json:"comno"`
	Openid             sql.NullString  `gorm:"column:openid" json:"openid"`
}

// TableName sets the insert table name for this struct type
func (v *VehicleAnalysis) TableName() string {
	return "vehicle_analysis"
}
