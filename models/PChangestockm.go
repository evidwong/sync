package models

import (
	"database/sql"
	"time"
)

type PChangestockm struct {
	ApprovalDate      time.Time       `gorm:"column:ApprovalDate" json:"ApprovalDate"`
	ArriveDays        sql.NullInt64   `gorm:"column:ArriveDays" json:"ArriveDays"`
	Attribute1        sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2        sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3        sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4        sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5        sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6        sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7        sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8        sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	Bin               sql.NullString  `gorm:"column:Bin" json:"Bin"`
	BorrowQty         sql.NullFloat64 `gorm:"column:BorrowQty" json:"BorrowQty"`
	COMNo             string          `gorm:"column:COMNo" json:"COMNo"`
	CostPrice         sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	DeliverPeriod     sql.NullInt64   `gorm:"column:DeliverPeriod" json:"DeliverPeriod"`
	EnglishName       sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode             sql.NullString  `gorm:"column:FCode" json:"FCode"`
	FOBCurrency       sql.NullString  `gorm:"column:FOBCurrency" json:"FOBCurrency"`
	FOBPrice          sql.NullFloat64 `gorm:"column:FOBPrice" json:"FOBPrice"`
	FOBRate           sql.NullFloat64 `gorm:"column:FOBRate" json:"FOBRate"`
	FirstInDate       time.Time       `gorm:"column:FirstInDate" json:"FirstInDate"`
	Grade             sql.NullString  `gorm:"column:Grade" json:"Grade"`
	GroupNo           sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	Height            sql.NullFloat64 `gorm:"column:Height" json:"Height"`
	InFlag1           sql.NullInt64   `gorm:"column:InFlag1" json:"InFlag1"`
	InFlag2           sql.NullInt64   `gorm:"column:InFlag2" json:"InFlag2"`
	InFlag3           sql.NullInt64   `gorm:"column:InFlag3" json:"InFlag3"`
	InPrice           sql.NullFloat64 `gorm:"column:InPrice" json:"InPrice"`
	InRemarks         sql.NullString  `gorm:"column:InRemarks" json:"InRemarks"`
	ItemName          sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	ItemRemarks       sql.NullString  `gorm:"column:ItemRemarks" json:"ItemRemarks"`
	LastInDate        time.Time       `gorm:"column:LastInDate" json:"LastInDate"`
	LastOutDate       time.Time       `gorm:"column:LastOutDate" json:"LastOutDate"`
	Length            sql.NullFloat64 `gorm:"column:Length" json:"Length"`
	Location          sql.NullString  `gorm:"column:Location" json:"Location"`
	MaxQty            sql.NullFloat64 `gorm:"column:MaxQty" json:"MaxQty"`
	MinOrderQty       sql.NullFloat64 `gorm:"column:MinOrderQty" json:"MinOrderQty"`
	MinQty            sql.NullFloat64 `gorm:"column:MinQty" json:"MinQty"`
	NoTaxCostPrice    sql.NullFloat64 `gorm:"column:NoTaxCostPrice" json:"NoTaxCostPrice"`
	OldBorrowQty      sql.NullFloat64 `gorm:"column:OldBorrowQty" json:"OldBorrowQty"`
	OldCostPrice      sql.NullFloat64 `gorm:"column:OldCostPrice" json:"OldCostPrice"`
	OldNoTaxCostPrice sql.NullFloat64 `gorm:"column:OldNoTaxCostPrice" json:"OldNoTaxCostPrice"`
	OldReseQty        sql.NullFloat64 `gorm:"column:OldReseQty" json:"OldReseQty"`
	OldSalesQty       sql.NullFloat64 `gorm:"column:OldSalesQty" json:"OldSalesQty"`
	OperatorCode      sql.NullString  `gorm:"column:OperatorCode" json:"OperatorCode"`
	OperatorName      sql.NullString  `gorm:"column:OperatorName" json:"OperatorName"`
	OrderPeriod       sql.NullInt64   `gorm:"column:OrderPeriod" json:"OrderPeriod"`
	PChangeStockCode  string          `gorm:"column:PChangeStockCode;primary_key" json:"PChangeStockCode;primary_key"`
	PYCode            sql.NullString  `gorm:"column:PYCode" json:"PYCode"`
	PackQty           sql.NullFloat64 `gorm:"column:PackQty" json:"PackQty"`
	PackUnit          sql.NullString  `gorm:"column:PackUnit" json:"PackUnit"`
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
	RefPrice1         sql.NullFloat64 `gorm:"column:RefPrice1" json:"RefPrice1"`
	RefPrice2         sql.NullFloat64 `gorm:"column:RefPrice2" json:"RefPrice2"`
	RefPrice3         sql.NullFloat64 `gorm:"column:RefPrice3" json:"RefPrice3"`
	RefPrice4         sql.NullFloat64 `gorm:"column:RefPrice4" json:"RefPrice4"`
	RelationNo        sql.NullInt64   `gorm:"column:RelationNo" json:"RelationNo"`
	Remarks           sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	ReseQty           sql.NullFloat64 `gorm:"column:ReseQty" json:"ReseQty"`
	SafeQty           sql.NullFloat64 `gorm:"column:SafeQty" json:"SafeQty"`
	SalesQty          sql.NullFloat64 `gorm:"column:SalesQty" json:"SalesQty"`
	StockID           sql.NullInt64   `gorm:"column:StockID" json:"StockID"`
	StockQty          sql.NullFloat64 `gorm:"column:StockQty" json:"StockQty"`
	StockRemarks      sql.NullString  `gorm:"column:StockRemarks" json:"StockRemarks"`
	SuitQty           sql.NullFloat64 `gorm:"column:SuitQty" json:"SuitQty"`
	SupplierCode      sql.NullString  `gorm:"column:SupplierCode" json:"SupplierCode"`
	SupplierName      sql.NullString  `gorm:"column:SupplierName" json:"SupplierName"`
	TmpField1         sql.NullString  `gorm:"column:TmpField1" json:"TmpField1"`
	TmpField2         time.Time       `gorm:"column:TmpField2" json:"TmpField2"`
	TradePrice        sql.NullFloat64 `gorm:"column:TradePrice" json:"TradePrice"`
	TransferPrice     sql.NullFloat64 `gorm:"column:TransferPrice" json:"TransferPrice"`
	Type1             sql.NullString  `gorm:"column:Type1" json:"Type1"`
	Type2             sql.NullString  `gorm:"column:Type2" json:"Type2"`
	Type3             sql.NullString  `gorm:"column:Type3" json:"Type3"`
	Type4             sql.NullString  `gorm:"column:Type4" json:"Type4"`
	Type5             sql.NullString  `gorm:"column:Type5" json:"Type5"`
	Unit              sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode       sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	ValidDate         time.Time       `gorm:"column:ValidDate" json:"ValidDate"`
	Vol               sql.NullFloat64 `gorm:"column:Vol" json:"Vol"`
	WBCode            sql.NullString  `gorm:"column:WBCode" json:"WBCode"`
	Weight            sql.NullFloat64 `gorm:"column:Weight" json:"Weight"`
	Width             sql.NullFloat64 `gorm:"column:Width" json:"Width"`
	Cid               sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode          string          `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (p *PChangestockm) TableName() string {
	return "p_changestockm"
}
