package models

import (
	"database/sql"
	"time"
)

type PStockm struct {
	ArriveDays       sql.NullInt64   `gorm:"column:ArriveDays" json:"ArriveDays"`
	Attribute1       sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2       sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3       sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4       sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5       sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6       sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7       sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8       sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	Bin              string          `gorm:"column:Bin" json:"Bin"`
	BorrowCostAmount sql.NullFloat64 `gorm:"column:BorrowCostAmount" json:"BorrowCostAmount"`
	BorrowQty        sql.NullFloat64 `gorm:"column:BorrowQty" json:"BorrowQty"`
	ComNo            string          `gorm:"column:ComNo" json:"ComNo"`
	CostAmount       sql.NullFloat64 `gorm:"column:CostAmount" json:"CostAmount"`
	CostPrice        float64         `gorm:"column:CostPrice" json:"CostPrice"`
	DeliverPeriod    sql.NullInt64   `gorm:"column:DeliverPeriod" json:"DeliverPeriod"`
	EnglishName      sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode            sql.NullString  `gorm:"column:FCode" json:"FCode"`
	FOBCurrency      sql.NullString  `gorm:"column:FOBCurrency" json:"FOBCurrency"`
	FOBPrice         sql.NullFloat64 `gorm:"column:FOBPrice" json:"FOBPrice"`
	FOBRate          sql.NullFloat64 `gorm:"column:FOBRate" json:"FOBRate"`
	FirstInDate      time.Time       `gorm:"column:FirstInDate" json:"FirstInDate"`
	Grade            sql.NullString  `gorm:"column:Grade" json:"Grade"`
	GroupNo          sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	Height           sql.NullFloat64 `gorm:"column:Height" json:"Height"`
	InCost           float64         `gorm:"column:InCost" json:"InCost"`
	InPrice          sql.NullFloat64 `gorm:"column:InPrice" json:"InPrice"`
	IsDel            int             `gorm:"column:IsDel" json:"IsDel"`
	ItemCode         string          `gorm:"column:ItemCode" json:"ItemCode"`
	ItemName         sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	ItemRemarks      sql.NullString  `gorm:"column:ItemRemarks" json:"ItemRemarks"`
	LastInDate       time.Time       `gorm:"column:LastInDate" json:"LastInDate"`
	LastOutDate      time.Time       `gorm:"column:LastOutDate" json:"LastOutDate"`
	Length           string          `gorm:"column:Length" json:"Length"`
	Location         string          `gorm:"column:Location" json:"Location"`
	MaxQty           sql.NullFloat64 `gorm:"column:MaxQty" json:"MaxQty"`
	MinOrderQty      sql.NullFloat64 `gorm:"column:MinOrderQty" json:"MinOrderQty"`
	MinQty           sql.NullFloat64 `gorm:"column:MinQty" json:"MinQty"`
	NoTaxCostAmount  sql.NullFloat64 `gorm:"column:NoTaxCostAmount" json:"NoTaxCostAmount"`
	NoTaxCostPrice   float64         `gorm:"column:NoTaxCostPrice" json:"NoTaxCostPrice"`
	OrderPeriod      sql.NullInt64   `gorm:"column:OrderPeriod" json:"OrderPeriod"`
	OutOrderQty      sql.NullFloat64 `gorm:"column:OutOrderQty" json:"OutOrderQty"`
	PYCode           sql.NullString  `gorm:"column:PYCode" json:"PYCode"`
	PackQty          sql.NullFloat64 `gorm:"column:PackQty" json:"PackQty"`
	PackUnit         sql.NullString  `gorm:"column:PackUnit" json:"PackUnit"`
	PrefPrice1       sql.NullFloat64 `gorm:"column:PrefPrice1" json:"PrefPrice1"`
	PrefPrice10      sql.NullFloat64 `gorm:"column:PrefPrice10" json:"PrefPrice10"`
	PrefPrice2       sql.NullFloat64 `gorm:"column:PrefPrice2" json:"PrefPrice2"`
	PrefPrice3       sql.NullFloat64 `gorm:"column:PrefPrice3" json:"PrefPrice3"`
	PrefPrice4       sql.NullFloat64 `gorm:"column:PrefPrice4" json:"PrefPrice4"`
	PrefPrice5       sql.NullFloat64 `gorm:"column:PrefPrice5" json:"PrefPrice5"`
	PrefPrice6       sql.NullFloat64 `gorm:"column:PrefPrice6" json:"PrefPrice6"`
	PrefPrice7       sql.NullFloat64 `gorm:"column:PrefPrice7" json:"PrefPrice7"`
	PrefPrice8       sql.NullFloat64 `gorm:"column:PrefPrice8" json:"PrefPrice8"`
	PrefPrice9       sql.NullFloat64 `gorm:"column:PrefPrice9" json:"PrefPrice9"`
	Price            sql.NullFloat64 `gorm:"column:Price" json:"Price"`
	RefPrice1        sql.NullFloat64 `gorm:"column:RefPrice1" json:"RefPrice1"`
	RefPrice2        sql.NullFloat64 `gorm:"column:RefPrice2" json:"RefPrice2"`
	RefPrice3        sql.NullFloat64 `gorm:"column:RefPrice3" json:"RefPrice3"`
	RefPrice4        sql.NullFloat64 `gorm:"column:RefPrice4" json:"RefPrice4"`
	RelationNo       sql.NullInt64   `gorm:"column:RelationNo" json:"RelationNo"`
	ReseCostAmount   sql.NullFloat64 `gorm:"column:ReseCostAmount" json:"ReseCostAmount"`
	ReseQty          sql.NullFloat64 `gorm:"column:ReseQty" json:"ReseQty"`
	SafeQty          sql.NullFloat64 `gorm:"column:SafeQty" json:"SafeQty"`
	SalesCostAmount  sql.NullFloat64 `gorm:"column:SalesCostAmount" json:"SalesCostAmount"`
	SalesQty         sql.NullFloat64 `gorm:"column:SalesQty" json:"SalesQty"`
	StockID          int             `gorm:"column:StockID;primary_key" json:"StockID;primary_key"`
	StockQty         sql.NullFloat64 `gorm:"column:StockQty" json:"StockQty"`
	SuitQty          sql.NullFloat64 `gorm:"column:SuitQty" json:"SuitQty"`
	SupplierCode     string          `gorm:"column:SupplierCode" json:"SupplierCode"`
	SupplierName     sql.NullString  `gorm:"column:SupplierName" json:"SupplierName"`
	TmpField1        sql.NullString  `gorm:"column:TmpField1" json:"TmpField1"`
	TradePrice       sql.NullFloat64 `gorm:"column:TradePrice" json:"TradePrice"`
	TransferPrice    sql.NullFloat64 `gorm:"column:TransferPrice" json:"TransferPrice"`
	Type1            sql.NullString  `gorm:"column:Type1" json:"Type1"`
	Type2            sql.NullString  `gorm:"column:Type2" json:"Type2"`
	Type3            sql.NullString  `gorm:"column:Type3" json:"Type3"`
	Type4            sql.NullString  `gorm:"column:Type4" json:"Type4"`
	Type5            sql.NullString  `gorm:"column:Type5" json:"Type5"`
	Unit             sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode      sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	Vol              sql.NullFloat64 `gorm:"column:Vol" json:"Vol"`
	WBCode           sql.NullString  `gorm:"column:WBCode" json:"WBCode"`
	Weight           sql.NullFloat64 `gorm:"column:Weight" json:"Weight"`
	Width            sql.NullFloat64 `gorm:"column:Width" json:"Width"`
	Attributecode    sql.NullString  `gorm:"column:attributecode" json:"attributecode"`
	Cid              int             `gorm:"column:cid;primary_key" json:"cid;primary_key"`
}

// TableName sets the insert table name for this struct type
func (p *PStockm) TableName() string {
	return "p_stockm"
}
