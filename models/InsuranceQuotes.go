package models

import (
	"database/sql"
	"time"
)

type InsuranceQuotes struct {
	CIFee                float32         `gorm:"column:CIFee" json:"CIFee"`
	EngineNo             sql.NullString  `gorm:"column:EngineNo" json:"EngineNo"`
	InsuranceCompany     sql.NullInt64   `gorm:"column:InsuranceCompany" json:"InsuranceCompany"`
	InsuranceCompanyName sql.NullString  `gorm:"column:InsuranceCompanyName" json:"InsuranceCompanyName"`
	InsuranceDate        time.Time       `gorm:"column:InsuranceDate" json:"InsuranceDate"`
	InsuranceTime        sql.NullInt64   `gorm:"column:InsuranceTime" json:"InsuranceTime"`
	IntentType           sql.NullInt64   `gorm:"column:IntentType" json:"IntentType"`
	RegisterNo           sql.NullString  `gorm:"column:RegisterNo" json:"RegisterNo"`
	VCIFee               float32         `gorm:"column:VCIFee" json:"VCIFee"`
	VVTFee               float32         `gorm:"column:VVTFee" json:"VVTFee"`
	Vin                  sql.NullString  `gorm:"column:Vin" json:"Vin"`
	AddDate              time.Time       `gorm:"column:add_date" json:"add_date"`
	AddTime              sql.NullInt64   `gorm:"column:add_time" json:"add_time"`
	Aid                  sql.NullInt64   `gorm:"column:aid" json:"aid"`
	Auid                 sql.NullInt64   `gorm:"column:auid" json:"auid"`
	BillingDate          time.Time       `gorm:"column:billingDate" json:"billingDate"`
	BillingTime          sql.NullInt64   `gorm:"column:billingTime" json:"billingTime"`
	BrandName            sql.NullString  `gorm:"column:brandName" json:"brandName"`
	BuyPrice             sql.NullFloat64 `gorm:"column:buyPrice" json:"buyPrice"`
	BuyTime              sql.NullInt64   `gorm:"column:buyTime" json:"buyTime"`
	CarModel             sql.NullInt64   `gorm:"column:carModel" json:"carModel"`
	CarModelText         sql.NullString  `gorm:"column:carModelText" json:"carModelText"`
	CarType              sql.NullInt64   `gorm:"column:carType" json:"carType"`
	CarTypeText          sql.NullString  `gorm:"column:carTypeText" json:"carTypeText"`
	Carid                sql.NullInt64   `gorm:"column:carid" json:"carid"`
	Cid                  sql.NullInt64   `gorm:"column:cid" json:"cid"`
	City                 sql.NullInt64   `gorm:"column:city" json:"city"`
	CityName             sql.NullString  `gorm:"column:cityName" json:"cityName"`
	ClerkFee             sql.NullFloat64 `gorm:"column:clerkFee" json:"clerkFee"`
	ClerkPercent         sql.NullFloat64 `gorm:"column:clerkPercent" json:"clerkPercent"`
	ClerkName            sql.NullString  `gorm:"column:clerk_name" json:"clerk_name"`
	Clerkid              sql.NullInt64   `gorm:"column:clerkid" json:"clerkid"`
	Coid                 sql.NullInt64   `gorm:"column:coid" json:"coid"`
	CommercialEndDate    time.Time       `gorm:"column:commercialEndDate" json:"commercialEndDate"`
	CommercialNo         sql.NullString  `gorm:"column:commercialNo" json:"commercialNo"`
	CommercialStartDate  time.Time       `gorm:"column:commercialStartDate" json:"commercialStartDate"`
	CompanyName          sql.NullString  `gorm:"column:companyName" json:"companyName"`
	CompulsoryEndDate    time.Time       `gorm:"column:compulsoryEndDate" json:"compulsoryEndDate"`
	CompulsoryNo         sql.NullString  `gorm:"column:compulsoryNo" json:"compulsoryNo"`
	CompulsoryStartDate  time.Time       `gorm:"column:compulsoryStartDate" json:"compulsoryStartDate"`
	Cost                 sql.NullFloat64 `gorm:"column:cost" json:"cost"`
	DeductibleFee        float32         `gorm:"column:deductibleFee" json:"deductibleFee"`
	DemandOrderid        sql.NullString  `gorm:"column:demandOrderid" json:"demandOrderid"`
	DemandCoid           sql.NullInt64   `gorm:"column:demand_coid" json:"demand_coid"`
	DemandCompany        sql.NullString  `gorm:"column:demand_company" json:"demand_company"`
	Demandid             sql.NullInt64   `gorm:"column:demandid" json:"demandid"`
	FailContent          sql.NullString  `gorm:"column:failContent" json:"failContent"`
	Flagid               sql.NullInt64   `gorm:"column:flagid" json:"flagid"`
	GetModel             sql.NullInt64   `gorm:"column:getModel" json:"getModel"`
	GetModelText         sql.NullString  `gorm:"column:getModelText" json:"getModelText"`
	GetType              sql.NullInt64   `gorm:"column:getType" json:"getType"`
	GetTypeText          sql.NullString  `gorm:"column:getTypeText" json:"getTypeText"`
	Gift                 sql.NullString  `gorm:"column:gift" json:"gift"`
	Human                sql.NullString  `gorm:"column:human" json:"human"`
	ID                   int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	IDCardPic            sql.NullString  `gorm:"column:idCardPic" json:"idCardPic"`
	InsurePic            sql.NullString  `gorm:"column:insurePic" json:"insurePic"`
	InvoiceFee           sql.NullFloat64 `gorm:"column:invoiceFee" json:"invoiceFee"`
	Isaccounting         int             `gorm:"column:isaccounting" json:"isaccounting"`
	Isbilling            int             `gorm:"column:isbilling" json:"isbilling"`
	Isdel                int             `gorm:"column:isdel" json:"isdel"`
	Isfail               sql.NullInt64   `gorm:"column:isfail" json:"isfail"`
	Ispost               int             `gorm:"column:ispost" json:"ispost"`
	ItemContent          sql.NullString  `gorm:"column:item_content" json:"item_content"`
	Itemid               sql.NullString  `gorm:"column:itemid" json:"itemid"`
	OfferFee             sql.NullFloat64 `gorm:"column:offerFee" json:"offerFee"`
	OfferPercent         sql.NullFloat64 `gorm:"column:offerPercent" json:"offerPercent"`
	Orderid              string          `gorm:"column:orderid" json:"orderid"`
	OrgCIFee             float32         `gorm:"column:orgCIFee" json:"orgCIFee"`
	OrgDeductibleFee     float32         `gorm:"column:orgDeductibleFee" json:"orgDeductibleFee"`
	OrgVCIFee            float32         `gorm:"column:orgVCIFee" json:"orgVCIFee"`
	OrgVVTFee            float32         `gorm:"column:orgVVTFee" json:"orgVVTFee"`
	PaidAmount           sql.NullFloat64 `gorm:"column:paid_amount" json:"paid_amount"`
	PayStatus            int             `gorm:"column:payStatus" json:"payStatus"`
	PayStatusText        string          `gorm:"column:payStatusText" json:"payStatusText"`
	Phone                sql.NullString  `gorm:"column:phone" json:"phone"`
	PostCM               sql.NullString  `gorm:"column:postCM" json:"postCM"`
	PostNo               sql.NullString  `gorm:"column:postNo" json:"postNo"`
	PostTime             time.Time       `gorm:"column:postTime" json:"postTime"`
	Province             sql.NullInt64   `gorm:"column:province" json:"province"`
	ProvinceName         sql.NullString  `gorm:"column:provinceName" json:"provinceName"`
	PushSms              sql.NullInt64   `gorm:"column:push_sms" json:"push_sms"`
	PushSmsCount         sql.NullInt64   `gorm:"column:push_sms_count" json:"push_sms_count"`
	PushWechat           sql.NullInt64   `gorm:"column:push_wechat" json:"push_wechat"`
	PushWechatCount      sql.NullInt64   `gorm:"column:push_wechat_count" json:"push_wechat_count"`
	QuotesTime           int             `gorm:"column:quotesTime" json:"quotesTime"`
	RebateFee            sql.NullFloat64 `gorm:"column:rebateFee" json:"rebateFee"`
	RebatePercent        sql.NullFloat64 `gorm:"column:rebatePercent" json:"rebatePercent"`
	Remark               sql.NullString  `gorm:"column:remark" json:"remark"`
	Ruleid               sql.NullInt64   `gorm:"column:ruleid" json:"ruleid"`
	RunLicCopy           sql.NullString  `gorm:"column:runLicCopy" json:"runLicCopy"`
	RunLicPic            sql.NullString  `gorm:"column:runLicPic" json:"runLicPic"`
	SeriesName           sql.NullString  `gorm:"column:seriesName" json:"seriesName"`
	Status               sql.NullInt64   `gorm:"column:status" json:"status"`
	StatusTxt            sql.NullString  `gorm:"column:statusTxt" json:"statusTxt"`
	StoreFee             sql.NullFloat64 `gorm:"column:storeFee" json:"storeFee"`
	StorePercent         sql.NullFloat64 `gorm:"column:storePercent" json:"storePercent"`
	StoreID              sql.NullInt64   `gorm:"column:store_id" json:"store_id"`
	StoreName            sql.NullString  `gorm:"column:store_name" json:"store_name"`
	TaxFee               float32         `gorm:"column:taxFee" json:"taxFee"`
	TaxPercent           float32         `gorm:"column:taxPercent" json:"taxPercent"`
	UID                  sql.NullInt64   `gorm:"column:uid" json:"uid"`
	UpdateDate           time.Time       `gorm:"column:update_date" json:"update_date"`
	UpdateTime           sql.NullInt64   `gorm:"column:update_time" json:"update_time"`
	Vehicle              sql.NullString  `gorm:"column:vehicle" json:"vehicle"`
	VehicleProvince      sql.NullString  `gorm:"column:vehicleProvince" json:"vehicleProvince"`
	WechatNickname       sql.NullString  `gorm:"column:wechat_nickname" json:"wechat_nickname"`
	Year                 sql.NullInt64   `gorm:"column:year" json:"year"`
}

// TableName sets the insert table name for this struct type
func (i *InsuranceQuotes) TableName() string {
	return "insurance_quotes"
}
