package models

import (
	"database/sql"
	"time"
)

type PInd struct {
	Attribute1             sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2             sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3             sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4             sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5             sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6             sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7             sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8             sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	Bin                    string          `gorm:"column:Bin" json:"Bin"`
	BoxNo                  sql.NullString  `gorm:"column:BoxNo" json:"BoxNo"`
	ComNo                  sql.NullString  `gorm:"column:ComNo" json:"ComNo"`
	CostAmount             sql.NullFloat64 `gorm:"column:CostAmount" json:"CostAmount"`
	CostPrice              sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	DepartmentD            sql.NullString  `gorm:"column:DepartmentD" json:"DepartmentD"`
	EnglishName            sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode                  sql.NullString  `gorm:"column:FCode" json:"FCode"`
	FOBCurrency            sql.NullString  `gorm:"column:FOBCurrency" json:"FOBCurrency"`
	FOBPrice               sql.NullFloat64 `gorm:"column:FOBPrice" json:"FOBPrice"`
	FOBRate                sql.NullFloat64 `gorm:"column:FOBRate" json:"FOBRate"`
	FastActionID           sql.NullInt64   `gorm:"column:FastActionID" json:"FastActionID"`
	FastInQty              sql.NullFloat64 `gorm:"column:FastInQty" json:"FastInQty"`
	FastOutCodeD           sql.NullString  `gorm:"column:FastOutCodeD" json:"FastOutCodeD"`
	FastOutIndexID         sql.NullInt64   `gorm:"column:FastOutIndexID" json:"FastOutIndexID"`
	Formula1               sql.NullFloat64 `gorm:"column:Formula1" json:"Formula1"`
	Formula2               sql.NullFloat64 `gorm:"column:Formula2" json:"Formula2"`
	Formula3               sql.NullFloat64 `gorm:"column:Formula3" json:"Formula3"`
	GroupNo                sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	InCost                 sql.NullFloat64 `gorm:"column:InCost" json:"InCost"`
	InCostAmount2          sql.NullFloat64 `gorm:"column:InCostAmount2" json:"InCostAmount2"`
	InFlag1                sql.NullInt64   `gorm:"column:InFlag1" json:"InFlag1"`
	InFlag2                sql.NullInt64   `gorm:"column:InFlag2" json:"InFlag2"`
	InFlag3                sql.NullInt64   `gorm:"column:InFlag3" json:"InFlag3"`
	InRemarks              sql.NullString  `gorm:"column:InRemarks" json:"InRemarks"`
	IndexID                int             `gorm:"column:IndexID" json:"IndexID"`
	IsFastD                sql.NullInt64   `gorm:"column:IsFastD" json:"IsFastD"`
	IsNewItemCode          sql.NullInt64   `gorm:"column:IsNewItemCode" json:"IsNewItemCode"`
	IsPresent              sql.NullInt64   `gorm:"column:IsPresent" json:"IsPresent"`
	ItemCode               string          `gorm:"column:ItemCode" json:"ItemCode"`
	ItemName               sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	ItemRemarks            sql.NullString  `gorm:"column:ItemRemarks" json:"ItemRemarks"`
	Location               string          `gorm:"column:Location" json:"Location"`
	MaxQty                 sql.NullFloat64 `gorm:"column:MaxQty" json:"MaxQty"`
	MinQty                 sql.NullFloat64 `gorm:"column:MinQty" json:"MinQty"`
	NoTaxCostAmount        sql.NullFloat64 `gorm:"column:NoTaxCostAmount" json:"NoTaxCostAmount"`
	NoTaxCostPrice         sql.NullFloat64 `gorm:"column:NoTaxCostPrice" json:"NoTaxCostPrice"`
	OldInCost              sql.NullFloat64 `gorm:"column:OldInCost" json:"OldInCost"`
	OldInCostAmount        sql.NullFloat64 `gorm:"column:OldInCostAmount" json:"OldInCostAmount"`
	OrderID                sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	PInCode                string          `gorm:"column:PInCode" json:"PInCode"`
	PPurchaseOrderCode     sql.NullString  `gorm:"column:PPurchaseOrderCode" json:"PPurchaseOrderCode"`
	PPurchaseOrderIndexID  sql.NullInt64   `gorm:"column:PPurchaseOrderIndexID" json:"PPurchaseOrderIndexID"`
	PQualityJudgeCodeD     sql.NullString  `gorm:"column:PQualityJudgeCodeD" json:"PQualityJudgeCodeD"`
	PYCode                 sql.NullString  `gorm:"column:PYCode" json:"PYCode"`
	PackQty                sql.NullFloat64 `gorm:"column:PackQty" json:"PackQty"`
	PackUnit               sql.NullString  `gorm:"column:PackUnit" json:"PackUnit"`
	PackagePieceQty        sql.NullFloat64 `gorm:"column:PackagePieceQty" json:"PackagePieceQty"`
	PackageQty             sql.NullInt64   `gorm:"column:PackageQty" json:"PackageQty"`
	PrefPrice1             sql.NullFloat64 `gorm:"column:PrefPrice1" json:"PrefPrice1"`
	PrefPrice10            sql.NullFloat64 `gorm:"column:PrefPrice10" json:"PrefPrice10"`
	PrefPrice2             sql.NullFloat64 `gorm:"column:PrefPrice2" json:"PrefPrice2"`
	PrefPrice3             sql.NullFloat64 `gorm:"column:PrefPrice3" json:"PrefPrice3"`
	PrefPrice4             sql.NullFloat64 `gorm:"column:PrefPrice4" json:"PrefPrice4"`
	PrefPrice5             sql.NullFloat64 `gorm:"column:PrefPrice5" json:"PrefPrice5"`
	PrefPrice6             sql.NullFloat64 `gorm:"column:PrefPrice6" json:"PrefPrice6"`
	PrefPrice7             sql.NullFloat64 `gorm:"column:PrefPrice7" json:"PrefPrice7"`
	PrefPrice8             sql.NullFloat64 `gorm:"column:PrefPrice8" json:"PrefPrice8"`
	PrefPrice9             sql.NullFloat64 `gorm:"column:PrefPrice9" json:"PrefPrice9"`
	Price                  sql.NullFloat64 `gorm:"column:Price" json:"Price"`
	Qty                    sql.NullFloat64 `gorm:"column:Qty" json:"Qty"`
	RefPrice1              sql.NullFloat64 `gorm:"column:RefPrice1" json:"RefPrice1"`
	RefPrice2              sql.NullFloat64 `gorm:"column:RefPrice2" json:"RefPrice2"`
	RefPrice3              sql.NullFloat64 `gorm:"column:RefPrice3" json:"RefPrice3"`
	RefPrice4              sql.NullFloat64 `gorm:"column:RefPrice4" json:"RefPrice4"`
	ReturnCostProfitAmount sql.NullFloat64 `gorm:"column:ReturnCostProfitAmount" json:"ReturnCostProfitAmount"`
	ReturnQty              sql.NullFloat64 `gorm:"column:ReturnQty" json:"ReturnQty"`
	SafeQty                sql.NullFloat64 `gorm:"column:SafeQty" json:"SafeQty"`
	SourceInCode           sql.NullString  `gorm:"column:SourceInCode" json:"SourceInCode"`
	SourceInCostPrice      sql.NullFloat64 `gorm:"column:SourceInCostPrice" json:"SourceInCostPrice"`
	SourceInIndexID        sql.NullInt64   `gorm:"column:SourceInIndexID" json:"SourceInIndexID"`
	SourceInNoTaxCostPrice sql.NullFloat64 `gorm:"column:SourceInNoTaxCostPrice" json:"SourceInNoTaxCostPrice"`
	SourceInQty            sql.NullFloat64 `gorm:"column:SourceInQty" json:"SourceInQty"`
	StockID                int64           `gorm:"column:StockID" json:"StockID"`
	SuitQty                sql.NullFloat64 `gorm:"column:SuitQty" json:"SuitQty"`
	SupplierCode           sql.NullString  `gorm:"column:SupplierCode" json:"SupplierCode"`
	SupplierName           sql.NullString  `gorm:"column:SupplierName" json:"SupplierName"`
	TmpField1              sql.NullString  `gorm:"column:TmpField1" json:"TmpField1"`
	TradePrice             sql.NullFloat64 `gorm:"column:TradePrice" json:"TradePrice"`
	TranIndexID            sql.NullInt64   `gorm:"column:TranIndexID" json:"TranIndexID"`
	TranOutCode            sql.NullString  `gorm:"column:TranOutCode" json:"TranOutCode"`
	TransferAmountD        sql.NullFloat64 `gorm:"column:TransferAmountD" json:"TransferAmountD"`
	TransferPrice          sql.NullFloat64 `gorm:"column:TransferPrice" json:"TransferPrice"`
	UnStockID              sql.NullInt64   `gorm:"column:UnStockID" json:"UnStockID"`
	Unit                   sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode            sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	ValidDate              time.Time       `gorm:"column:ValidDate" json:"ValidDate"`
	WBCode                 sql.NullString  `gorm:"column:WBCode" json:"WBCode"`
	Attributecode          sql.NullString  `gorm:"column:attributecode" json:"attributecode"`
	Cid                    sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ID                     int64           `gorm:"column:id;primary_key" json:"id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (p *PInd) TableName() string {
	return "p_ind"
}
