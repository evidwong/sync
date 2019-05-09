package models

import (
	"database/sql"
	"time"
)

type InsuranceOrder struct {
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
	AccountingTime       int             `gorm:"column:accountingTime" json:"accountingTime"`
	AccountingTimeDate   time.Time       `gorm:"column:accountingTimeDate" json:"accountingTimeDate"`
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
	ClerkFee             float32         `gorm:"column:clerkFee" json:"clerkFee"`
	ClerkID              sql.NullInt64   `gorm:"column:clerkId" json:"clerkId"`
	ClerkPercent         float32         `gorm:"column:clerkPercent" json:"clerkPercent"`
	ClerkName            sql.NullString  `gorm:"column:clerk_name" json:"clerk_name"`
	Coid                 sql.NullInt64   `gorm:"column:coid" json:"coid"`
	CommercialEndDate    time.Time       `gorm:"column:commercialEndDate" json:"commercialEndDate"`
	CommercialNo         sql.NullString  `gorm:"column:commercialNo" json:"commercialNo"`
	CommercialStartDate  time.Time       `gorm:"column:commercialStartDate" json:"commercialStartDate"`
	CompanyName          sql.NullString  `gorm:"column:companyName" json:"companyName"`
	CompulsoryEndDate    time.Time       `gorm:"column:compulsoryEndDate" json:"compulsoryEndDate"`
	CompulsoryNo         sql.NullString  `gorm:"column:compulsoryNo" json:"compulsoryNo"`
	CompulsoryStartDate  time.Time       `gorm:"column:compulsoryStartDate" json:"compulsoryStartDate"`
	Cost                 float32         `gorm:"column:cost" json:"cost"`
	DeductibleFee        float32         `gorm:"column:deductibleFee" json:"deductibleFee"`
	DemandOrderid        sql.NullString  `gorm:"column:demandOrderid" json:"demandOrderid"`
	DemandCoid           sql.NullInt64   `gorm:"column:demand_coid" json:"demand_coid"`
	DemandCompany        sql.NullString  `gorm:"column:demand_company" json:"demand_company"`
	Demandid             sql.NullInt64   `gorm:"column:demandid" json:"demandid"`
	FailContent          sql.NullString  `gorm:"column:failContent" json:"failContent"`
	Fee                  float32         `gorm:"column:fee" json:"fee"`
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
	InvoiceFee           float32         `gorm:"column:invoiceFee" json:"invoiceFee"`
	Isaccounting         int             `gorm:"column:isaccounting" json:"isaccounting"`
	Isbilling            int             `gorm:"column:isbilling" json:"isbilling"`
	Isdel                int             `gorm:"column:isdel" json:"isdel"`
	Isfail               int             `gorm:"column:isfail" json:"isfail"`
	Ispost               int             `gorm:"column:ispost" json:"ispost"`
	ItemContent          sql.NullString  `gorm:"column:item_content" json:"item_content"`
	Itemid               sql.NullString  `gorm:"column:itemid" json:"itemid"`
	OfferFee             float32         `gorm:"column:offerFee" json:"offerFee"`
	OfferPercent         float32         `gorm:"column:offerPercent" json:"offerPercent"`
	Orderid              sql.NullString  `gorm:"column:orderid" json:"orderid"`
	OrgCIFee             float32         `gorm:"column:orgCIFee" json:"orgCIFee"`
	OrgDeductibleFee     float32         `gorm:"column:orgDeductibleFee" json:"orgDeductibleFee"`
	OrgVCIFee            float32         `gorm:"column:orgVCIFee" json:"orgVCIFee"`
	OrgVVTFee            float32         `gorm:"column:orgVVTFee" json:"orgVVTFee"`
	PaidAmount           sql.NullFloat64 `gorm:"column:paid_amount" json:"paid_amount"`
	PayFee               float32         `gorm:"column:payFee" json:"payFee"`
	PayModel             sql.NullInt64   `gorm:"column:payModel" json:"payModel"`
	PayModelText         sql.NullString  `gorm:"column:payModelText" json:"payModelText"`
	PayNickname          sql.NullString  `gorm:"column:payNickname" json:"payNickname"`
	PayNo                sql.NullString  `gorm:"column:payNo" json:"payNo"`
	PayNotify            sql.NullString  `gorm:"column:payNotify" json:"payNotify"`
	PayOpenid            sql.NullString  `gorm:"column:payOpenid" json:"payOpenid"`
	PayStatus            int             `gorm:"column:payStatus" json:"payStatus"`
	PayStatusText        string          `gorm:"column:payStatusText" json:"payStatusText"`
	PayTime              sql.NullInt64   `gorm:"column:payTime" json:"payTime"`
	PayTimeText          time.Time       `gorm:"column:payTimeText" json:"payTimeText"`
	PayType              sql.NullInt64   `gorm:"column:payType" json:"payType"`
	PayTypeText          sql.NullString  `gorm:"column:payTypeText" json:"payTypeText"`
	Phone                sql.NullString  `gorm:"column:phone" json:"phone"`
	PostCM               sql.NullString  `gorm:"column:postCM" json:"postCM"`
	PostNo               sql.NullString  `gorm:"column:postNo" json:"postNo"`
	PostTime             time.Time       `gorm:"column:postTime" json:"postTime"`
	Province             sql.NullInt64   `gorm:"column:province" json:"province"`
	ProvinceName         sql.NullString  `gorm:"column:provinceName" json:"provinceName"`
	PushSms              int             `gorm:"column:push_sms" json:"push_sms"`
	PushSmsCount         int             `gorm:"column:push_sms_count" json:"push_sms_count"`
	PushWechat           int             `gorm:"column:push_wechat" json:"push_wechat"`
	PushWechatCount      int             `gorm:"column:push_wechat_count" json:"push_wechat_count"`
	QuotesOrderid        sql.NullString  `gorm:"column:quotesOrderid" json:"quotesOrderid"`
	QuotesTime           sql.NullInt64   `gorm:"column:quotesTime" json:"quotesTime"`
	QuotesTimeDate       time.Time       `gorm:"column:quotesTimeDate" json:"quotesTimeDate"`
	Quotesid             int             `gorm:"column:quotesid" json:"quotesid"`
	RebateFee            float32         `gorm:"column:rebateFee" json:"rebateFee"`
	RebatePercent        float32         `gorm:"column:rebatePercent" json:"rebatePercent"`
	Remark               sql.NullString  `gorm:"column:remark" json:"remark"`
	Ruleid               sql.NullInt64   `gorm:"column:ruleid" json:"ruleid"`
	RunLicCopy           sql.NullString  `gorm:"column:runLicCopy" json:"runLicCopy"`
	RunLicPic            sql.NullString  `gorm:"column:runLicPic" json:"runLicPic"`
	SeriesName           sql.NullString  `gorm:"column:seriesName" json:"seriesName"`
	SettlementTime       int             `gorm:"column:settlementTime" json:"settlementTime"`
	Status               sql.NullInt64   `gorm:"column:status" json:"status"`
	StatusTxt            sql.NullString  `gorm:"column:statusTxt" json:"statusTxt"`
	StoreFee             sql.NullFloat64 `gorm:"column:storeFee" json:"storeFee"`
	StorePercent         sql.NullFloat64 `gorm:"column:storePercent" json:"storePercent"`
	StoreID              sql.NullInt64   `gorm:"column:store_id" json:"store_id"`
	StoreName            sql.NullString  `gorm:"column:store_name" json:"store_name"`
	TaxFee               float32         `gorm:"column:taxFee" json:"taxFee"`
	TaxPercent           float32         `gorm:"column:taxPercent" json:"taxPercent"`
	TransactionID        sql.NullString  `gorm:"column:transaction_id" json:"transaction_id"`
	UID                  sql.NullInt64   `gorm:"column:uid" json:"uid"`
	UpdateDate           time.Time       `gorm:"column:update_date" json:"update_date"`
	UpdateTime           sql.NullInt64   `gorm:"column:update_time" json:"update_time"`
	Vehicle              sql.NullString  `gorm:"column:vehicle" json:"vehicle"`
	VehicleProvince      sql.NullString  `gorm:"column:vehicleProvince" json:"vehicleProvince"`
	WechatNickname       sql.NullString  `gorm:"column:wechat_nickname" json:"wechat_nickname"`
	Year                 sql.NullInt64   `gorm:"column:year" json:"year"`
}

// TableName sets the insert table name for this struct type
func (i *InsuranceOrder) TableName() string {
	return "insurance_order"
}
