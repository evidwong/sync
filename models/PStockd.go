package models

import (
	"database/sql"
	"time"
)

type PStockd struct {
	ArriveDays      sql.NullInt64   `gorm:"column:ArriveDays" json:"ArriveDays"`
	Attribute1      sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2      sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3      sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4      sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5      sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6      sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7      sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8      sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	Bin             string          `gorm:"column:Bin" json:"Bin"`
	BorrowQty       sql.NullFloat64 `gorm:"column:BorrowQty" json:"BorrowQty"`
	CalcStockQty    sql.NullFloat64 `gorm:"column:CalcStockQty" json:"CalcStockQty"`
	ComNo           string          `gorm:"column:ComNo" json:"ComNo"`
	CostPrice       sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	DeliverPeriod   sql.NullInt64   `gorm:"column:DeliverPeriod" json:"DeliverPeriod"`
	EnglishName     sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode           sql.NullString  `gorm:"column:FCode" json:"FCode"`
	FOBCurrency     sql.NullString  `gorm:"column:FOBCurrency" json:"FOBCurrency"`
	FOBPrice        sql.NullFloat64 `gorm:"column:FOBPrice" json:"FOBPrice"`
	FOBRate         sql.NullFloat64 `gorm:"column:FOBRate" json:"FOBRate"`
	FirstInDate     time.Time       `gorm:"column:FirstInDate" json:"FirstInDate"`
	Grade           sql.NullString  `gorm:"column:Grade" json:"Grade"`
	GroupNo         sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	Height          sql.NullFloat64 `gorm:"column:Height" json:"Height"`
	InCost          sql.NullFloat64 `gorm:"column:InCost" json:"InCost"`
	InFlag1         sql.NullInt64   `gorm:"column:InFlag1" json:"InFlag1"`
	InFlag2         sql.NullInt64   `gorm:"column:InFlag2" json:"InFlag2"`
	InFlag3         sql.NullInt64   `gorm:"column:InFlag3" json:"InFlag3"`
	InPrice         sql.NullFloat64 `gorm:"column:InPrice" json:"InPrice"`
	InRemarks       sql.NullString  `gorm:"column:InRemarks" json:"InRemarks"`
	IsDel           int             `gorm:"column:IsDel" json:"IsDel"`
	ItemName        sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	ItemRemarks     sql.NullString  `gorm:"column:ItemRemarks" json:"ItemRemarks"`
	LastInDate      time.Time       `gorm:"column:LastInDate" json:"LastInDate"`
	LastOutDate     time.Time       `gorm:"column:LastOutDate" json:"LastOutDate"`
	Length          string          `gorm:"column:Length" json:"Length"`
	Location        string          `gorm:"column:Location" json:"Location"`
	MaxQty          sql.NullFloat64 `gorm:"column:MaxQty" json:"MaxQty"`
	MinOrderQty     sql.NullFloat64 `gorm:"column:MinOrderQty" json:"MinOrderQty"`
	MinQty          sql.NullFloat64 `gorm:"column:MinQty" json:"MinQty"`
	NoTaxCostPrice  sql.NullFloat64 `gorm:"column:NoTaxCostPrice" json:"NoTaxCostPrice"`
	NotAppQty       sql.NullFloat64 `gorm:"column:NotAppQty" json:"NotAppQty"`
	OrderPeriod     sql.NullInt64   `gorm:"column:OrderPeriod" json:"OrderPeriod"`
	OutOrderQty     sql.NullFloat64 `gorm:"column:OutOrderQty" json:"OutOrderQty"`
	PInCode         sql.NullString  `gorm:"column:PInCode" json:"PInCode"`
	PStockCheckCode sql.NullString  `gorm:"column:PStockCheckCode" json:"PStockCheckCode"`
	PYCode          sql.NullString  `gorm:"column:PYCode" json:"PYCode"`
	PackQty         sql.NullFloat64 `gorm:"column:PackQty" json:"PackQty"`
	PackUnit        sql.NullString  `gorm:"column:PackUnit" json:"PackUnit"`
	PrefPrice1      sql.NullFloat64 `gorm:"column:PrefPrice1" json:"PrefPrice1"`
	PrefPrice10     sql.NullFloat64 `gorm:"column:PrefPrice10" json:"PrefPrice10"`
	PrefPrice2      sql.NullFloat64 `gorm:"column:PrefPrice2" json:"PrefPrice2"`
	PrefPrice3      sql.NullFloat64 `gorm:"column:PrefPrice3" json:"PrefPrice3"`
	PrefPrice4      sql.NullFloat64 `gorm:"column:PrefPrice4" json:"PrefPrice4"`
	PrefPrice5      sql.NullFloat64 `gorm:"column:PrefPrice5" json:"PrefPrice5"`
	PrefPrice6      sql.NullFloat64 `gorm:"column:PrefPrice6" json:"PrefPrice6"`
	PrefPrice7      sql.NullFloat64 `gorm:"column:PrefPrice7" json:"PrefPrice7"`
	PrefPrice8      sql.NullFloat64 `gorm:"column:PrefPrice8" json:"PrefPrice8"`
	PrefPrice9      sql.NullFloat64 `gorm:"column:PrefPrice9" json:"PrefPrice9"`
	Price           sql.NullFloat64 `gorm:"column:Price" json:"Price"`
	RefPrice1       sql.NullFloat64 `gorm:"column:RefPrice1" json:"RefPrice1"`
	RefPrice2       sql.NullFloat64 `gorm:"column:RefPrice2" json:"RefPrice2"`
	RefPrice3       sql.NullFloat64 `gorm:"column:RefPrice3" json:"RefPrice3"`
	RefPrice4       sql.NullFloat64 `gorm:"column:RefPrice4" json:"RefPrice4"`
	RelationCode    sql.NullString  `gorm:"column:RelationCode" json:"RelationCode"`
	RelationNo      sql.NullInt64   `gorm:"column:RelationNo" json:"RelationNo"`
	ReseQty         sql.NullFloat64 `gorm:"column:ReseQty" json:"ReseQty"`
	SafeQty         sql.NullFloat64 `gorm:"column:SafeQty" json:"SafeQty"`
	SalesQty        sql.NullFloat64 `gorm:"column:SalesQty" json:"SalesQty"`
	StockCheckDate  time.Time       `gorm:"column:StockCheckDate" json:"StockCheckDate"`
	StockID         int             `gorm:"column:StockID" json:"StockID"`
	StockQty        sql.NullFloat64 `gorm:"column:StockQty" json:"StockQty"`
	StockRemarks    sql.NullString  `gorm:"column:StockRemarks" json:"StockRemarks"`
	SuitQty         sql.NullFloat64 `gorm:"column:SuitQty" json:"SuitQty"`
	SupplierCode    string          `gorm:"column:SupplierCode" json:"SupplierCode"`
	SupplierName    sql.NullString  `gorm:"column:SupplierName" json:"SupplierName"`
	TmpField1       sql.NullString  `gorm:"column:TmpField1" json:"TmpField1"`
	TmpField2       time.Time       `gorm:"column:TmpField2" json:"TmpField2"`
	TradePrice      sql.NullFloat64 `gorm:"column:TradePrice" json:"TradePrice"`
	TransferPrice   sql.NullFloat64 `gorm:"column:TransferPrice" json:"TransferPrice"`
	Type1           sql.NullString  `gorm:"column:Type1" json:"Type1"`
	Type2           sql.NullString  `gorm:"column:Type2" json:"Type2"`
	Type3           sql.NullString  `gorm:"column:Type3" json:"Type3"`
	Type4           sql.NullString  `gorm:"column:Type4" json:"Type4"`
	Type5           sql.NullString  `gorm:"column:Type5" json:"Type5"`
	Unit            sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode     sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	ValidDate       time.Time       `gorm:"column:ValidDate" json:"ValidDate"`
	Vol             sql.NullFloat64 `gorm:"column:Vol" json:"Vol"`
	WBCode          sql.NullString  `gorm:"column:WBCode" json:"WBCode"`
	Weight          sql.NullFloat64 `gorm:"column:Weight" json:"Weight"`
	Width           sql.NullFloat64 `gorm:"column:Width" json:"Width"`
	Attributecode   sql.NullString  `gorm:"column:attributecode" json:"attributecode"`
	Cid             int             `gorm:"column:cid" json:"cid"`
	Itemcode        string          `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (p *PStockd) TableName() string {
	return "p_stockd"
}
