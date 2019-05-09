package models

import (
	"database/sql"
	"time"
)

type InsuranceDemand struct {
	EngineNo             sql.NullString  `gorm:"column:EngineNo" json:"EngineNo"`
	InsuranceCompany     sql.NullInt64   `gorm:"column:InsuranceCompany" json:"InsuranceCompany"`
	InsuranceCompanyName sql.NullString  `gorm:"column:InsuranceCompanyName" json:"InsuranceCompanyName"`
	InsuranceDate        time.Time       `gorm:"column:InsuranceDate" json:"InsuranceDate"`
	InsuranceTime        sql.NullInt64   `gorm:"column:InsuranceTime" json:"InsuranceTime"`
	IntentType           sql.NullInt64   `gorm:"column:IntentType" json:"IntentType"`
	RegisterNo           sql.NullString  `gorm:"column:RegisterNo" json:"RegisterNo"`
	Vin                  sql.NullString  `gorm:"column:Vin" json:"Vin"`
	AddDate              time.Time       `gorm:"column:add_date" json:"add_date"`
	AddTime              int             `gorm:"column:add_time" json:"add_time"`
	Aid                  sql.NullInt64   `gorm:"column:aid" json:"aid"`
	Auid                 sql.NullInt64   `gorm:"column:auid" json:"auid"`
	BillingDate          time.Time       `gorm:"column:billingDate" json:"billingDate"`
	BillingTime          sql.NullInt64   `gorm:"column:billingTime" json:"billingTime"`
	BrandName            sql.NullString  `gorm:"column:brandName" json:"brandName"`
	BuyPrice             sql.NullFloat64 `gorm:"column:buyPrice" json:"buyPrice"`
	BuyTime              sql.NullString  `gorm:"column:buyTime" json:"buyTime"`
	CarModel             sql.NullInt64   `gorm:"column:carModel" json:"carModel"`
	CarModelText         sql.NullString  `gorm:"column:carModelText" json:"carModelText"`
	CarType              sql.NullInt64   `gorm:"column:carType" json:"carType"`
	CarTypeText          sql.NullString  `gorm:"column:carTypeText" json:"carTypeText"`
	Carid                sql.NullInt64   `gorm:"column:carid" json:"carid"`
	Cid                  sql.NullInt64   `gorm:"column:cid" json:"cid"`
	City                 sql.NullInt64   `gorm:"column:city" json:"city"`
	CityName             sql.NullString  `gorm:"column:cityName" json:"cityName"`
	ClerkName            sql.NullString  `gorm:"column:clerk_name" json:"clerk_name"`
	Clerkid              sql.NullInt64   `gorm:"column:clerkid" json:"clerkid"`
	Coid                 sql.NullInt64   `gorm:"column:coid" json:"coid"`
	CommercialEndDate    time.Time       `gorm:"column:commercialEndDate" json:"commercialEndDate"`
	CommercialNo         sql.NullString  `gorm:"column:commercialNo" json:"commercialNo"`
	CommercialStartDate  time.Time       `gorm:"column:commercialStartDate" json:"commercialStartDate"`
	Comno                sql.NullString  `gorm:"column:comno" json:"comno"`
	CompanyName          sql.NullString  `gorm:"column:companyName" json:"companyName"`
	CompulsoryEndDate    time.Time       `gorm:"column:compulsoryEndDate" json:"compulsoryEndDate"`
	CompulsoryNo         sql.NullString  `gorm:"column:compulsoryNo" json:"compulsoryNo"`
	CompulsoryStartDate  time.Time       `gorm:"column:compulsoryStartDate" json:"compulsoryStartDate"`
	DemandCoid           int             `gorm:"column:demand_coid" json:"demand_coid"`
	DemandCompany        string          `gorm:"column:demand_company" json:"demand_company"`
	Flagid               sql.NullInt64   `gorm:"column:flagid" json:"flagid"`
	GetModel             sql.NullInt64   `gorm:"column:getModel" json:"getModel"`
	GetModelText         sql.NullString  `gorm:"column:getModelText" json:"getModelText"`
	GetType              sql.NullInt64   `gorm:"column:getType" json:"getType"`
	GetTypeText          sql.NullString  `gorm:"column:getTypeText" json:"getTypeText"`
	Human                sql.NullString  `gorm:"column:human" json:"human"`
	ID                   int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	IDCardPic            sql.NullString  `gorm:"column:idCardPic" json:"idCardPic"`
	IDCardPicBackPic     sql.NullString  `gorm:"column:idCardPicBackPic" json:"idCardPicBackPic"`
	InsurePic            sql.NullString  `gorm:"column:insurePic" json:"insurePic"`
	Isbilling            int             `gorm:"column:isbilling" json:"isbilling"`
	Isdel                sql.NullInt64   `gorm:"column:isdel" json:"isdel"`
	Ispost               int             `gorm:"column:ispost" json:"ispost"`
	Itemid               sql.NullString  `gorm:"column:itemid" json:"itemid"`
	Orderid              sql.NullString  `gorm:"column:orderid" json:"orderid"`
	PayStatus            int             `gorm:"column:payStatus" json:"payStatus"`
	PayStatusText        string          `gorm:"column:payStatusText" json:"payStatusText"`
	Phone                sql.NullString  `gorm:"column:phone" json:"phone"`
	PostCM               sql.NullString  `gorm:"column:postCM" json:"postCM"`
	PostNo               sql.NullString  `gorm:"column:postNo" json:"postNo"`
	PostTime             time.Time       `gorm:"column:postTime" json:"postTime"`
	Province             sql.NullInt64   `gorm:"column:province" json:"province"`
	ProvinceName         sql.NullString  `gorm:"column:provinceName" json:"provinceName"`
	QuotesTime           int             `gorm:"column:quotesTime" json:"quotesTime"`
	Remark               sql.NullString  `gorm:"column:remark" json:"remark"`
	RunLicCopy           sql.NullString  `gorm:"column:runLicCopy" json:"runLicCopy"`
	RunLicPic            sql.NullString  `gorm:"column:runLicPic" json:"runLicPic"`
	SeriesName           sql.NullString  `gorm:"column:seriesName" json:"seriesName"`
	Status               sql.NullInt64   `gorm:"column:status" json:"status"`
	StatusTxt            sql.NullString  `gorm:"column:statusTxt" json:"statusTxt"`
	StoreID              sql.NullInt64   `gorm:"column:store_id" json:"store_id"`
	StoreName            sql.NullString  `gorm:"column:store_name" json:"store_name"`
	UID                  int             `gorm:"column:uid" json:"uid"`
	UpdateDate           time.Time       `gorm:"column:update_date" json:"update_date"`
	UpdateTime           sql.NullInt64   `gorm:"column:update_time" json:"update_time"`
	Vehicle              sql.NullString  `gorm:"column:vehicle" json:"vehicle"`
	VehicleProvince      sql.NullString  `gorm:"column:vehicleProvince" json:"vehicleProvince"`
	WechatNickname       sql.NullString  `gorm:"column:wechat_nickname" json:"wechat_nickname"`
	Year                 sql.NullString  `gorm:"column:year" json:"year"`
}

// TableName sets the insert table name for this struct type
func (i *InsuranceDemand) TableName() string {
	return "insurance_demand"
}
