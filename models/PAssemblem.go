package models

import (
	"database/sql"
	"time"
)

type PAssemblem struct {
	ApprovalDate      time.Time       `gorm:"column:ApprovalDate" json:"ApprovalDate"`
	ApproverCode      sql.NullString  `gorm:"column:ApproverCode" json:"ApproverCode"`
	ApproverName      sql.NullString  `gorm:"column:ApproverName" json:"ApproverName"`
	Attribute1        sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2        sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3        sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4        sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5        sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6        sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7        sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8        sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	Bin               sql.NullString  `gorm:"column:Bin" json:"Bin"`
	COMNo             string          `gorm:"column:COMNo" json:"COMNo"`
	CostPrice         sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	DocDate           time.Time       `gorm:"column:DocDate" json:"DocDate"`
	EnglishName       sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode             sql.NullString  `gorm:"column:FCode" json:"FCode"`
	FOBCurrency       sql.NullString  `gorm:"column:FOBCurrency" json:"FOBCurrency"`
	FOBPrice          sql.NullFloat64 `gorm:"column:FOBPrice" json:"FOBPrice"`
	FOBRate           sql.NullFloat64 `gorm:"column:FOBRate" json:"FOBRate"`
	GroupNo           sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	ItemName          sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	Location          sql.NullString  `gorm:"column:Location" json:"Location"`
	ModifyDate        time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP          sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode  sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName  sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	NoTaxCostPrice    sql.NullFloat64 `gorm:"column:NoTaxCostPrice" json:"NoTaxCostPrice"`
	OperatorCode      sql.NullString  `gorm:"column:OperatorCode" json:"OperatorCode"`
	OperatorName      sql.NullString  `gorm:"column:OperatorName" json:"OperatorName"`
	PAssembleCode     string          `gorm:"column:PAssembleCode;primary_key" json:"PAssembleCode;primary_key"`
	PYCode            sql.NullString  `gorm:"column:PYCode" json:"PYCode"`
	PrefPrice1        sql.NullFloat64 `gorm:"column:PrefPrice1" json:"PrefPrice1"`
	PrefPrice10       sql.NullFloat64 `gorm:"column:PrefPrice10" json:"PrefPrice10"`
	PrefPrice2        sql.NullFloat64 `gorm:"column:PrefPrice2" json:"PrefPrice2"`
	PrefPrice3        sql.NullFloat64 `gorm:"column:PrefPrice3" json:"PrefPrice3"`
	PrefPrice4        sql.NullFloat64 `gorm:"column:PrefPrice4" json:"PrefPrice4"`
	PrefPrice5        sql.NullFloat64 `gorm:"column:PrefPrice5" json:"PrefPrice5"`
	PrefPrice6        sql.NullFloat64 `gorm:"column:PrefPrice6" json:"PrefPrice6"`
	PrefPrice7        sql.NullFloat64 `gorm:"column:PrefPrice7" json:"PrefPrice7"`
	PrefPrice8        sql.NullFloat64 `gorm:"column:PrefPrice8" json:"PrefPrice8"`
	PrefPrice9        sql.NullFloat64 `gorm:"column:PrefPrice9" json:"PrefPrice9"`
	Price             sql.NullFloat64 `gorm:"column:Price" json:"Price"`
	Qty               sql.NullFloat64 `gorm:"column:Qty" json:"Qty"`
	RefPrice1         sql.NullFloat64 `gorm:"column:RefPrice1" json:"RefPrice1"`
	RefPrice2         sql.NullFloat64 `gorm:"column:RefPrice2" json:"RefPrice2"`
	RefPrice3         sql.NullFloat64 `gorm:"column:RefPrice3" json:"RefPrice3"`
	RefPrice4         sql.NullFloat64 `gorm:"column:RefPrice4" json:"RefPrice4"`
	Remarks           sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	StockID           sql.NullInt64   `gorm:"column:StockID" json:"StockID"`
	SupplierCode      sql.NullString  `gorm:"column:SupplierCode" json:"SupplierCode"`
	SupplierName      sql.NullString  `gorm:"column:SupplierName" json:"SupplierName"`
	TmpField1         sql.NullString  `gorm:"column:TmpField1" json:"TmpField1"`
	TradePrice        sql.NullFloat64 `gorm:"column:TradePrice" json:"TradePrice"`
	TransferPrice     sql.NullFloat64 `gorm:"column:TransferPrice" json:"TransferPrice"`
	Unit              sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode       sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	ValidMDate        time.Time       `gorm:"column:ValidMDate" json:"ValidMDate"`
	WBCode            sql.NullString  `gorm:"column:WBCode" json:"WBCode"`
	Cid               sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode          sql.NullString  `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (p *PAssemblem) TableName() string {
	return "p_assemblem"
}
