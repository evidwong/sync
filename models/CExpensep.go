package models

import "database/sql"

type CExpensep struct {
	Amount              sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	Attribute1          sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2          sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3          sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4          sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5          sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6          sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7          sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8          sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	Bin                 sql.NullString  `gorm:"column:Bin" json:"Bin"`
	CExpenseCode        string          `gorm:"column:CExpenseCode" json:"CExpenseCode"`
	CIWPCode            sql.NullString  `gorm:"column:CIWPCode" json:"CIWPCode"`
	CReceptionCode      sql.NullString  `gorm:"column:CReceptionCode" json:"CReceptionCode"`
	Class               sql.NullString  `gorm:"column:Class" json:"Class"`
	CostPrice           sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	DepartmentD         sql.NullString  `gorm:"column:DepartmentD" json:"DepartmentD"`
	DiscountRate        sql.NullFloat64 `gorm:"column:DiscountRate" json:"DiscountRate"`
	EPRemarks           sql.NullString  `gorm:"column:EPRemarks" json:"EPRemarks"`
	EnglishName         sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode               sql.NullString  `gorm:"column:FCode" json:"FCode"`
	FOBCurrency         sql.NullString  `gorm:"column:FOBCurrency" json:"FOBCurrency"`
	FOBPrice            sql.NullFloat64 `gorm:"column:FOBPrice" json:"FOBPrice"`
	FOBRate             sql.NullFloat64 `gorm:"column:FOBRate" json:"FOBRate"`
	FalseItemName       sql.NullString  `gorm:"column:FalseItemName" json:"FalseItemName"`
	FunctionCode        string          `gorm:"column:FunctionCode" json:"FunctionCode"`
	GroupNo             sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	GroupWareCode       sql.NullString  `gorm:"column:GroupWareCode" json:"GroupWareCode"`
	IndexID             int             `gorm:"column:IndexID" json:"IndexID"`
	Integral            sql.NullFloat64 `gorm:"column:Integral" json:"Integral"`
	IntegralProportion  sql.NullFloat64 `gorm:"column:IntegralProportion" json:"IntegralProportion"`
	IsCanIntegral       sql.NullInt64   `gorm:"column:IsCanIntegral" json:"IsCanIntegral"`
	IsCheckPLC          sql.NullInt64   `gorm:"column:IsCheckPLC" json:"IsCheckPLC"`
	IsCheckPLT          sql.NullInt64   `gorm:"column:IsCheckPLT" json:"IsCheckPLT"`
	IsEPFlag            sql.NullInt64   `gorm:"column:IsEPFlag" json:"IsEPFlag"`
	IsGroupWare         sql.NullInt64   `gorm:"column:IsGroupWare" json:"IsGroupWare"`
	IsMemberSet         sql.NullInt64   `gorm:"column:IsMemberSet" json:"IsMemberSet"`
	IsPresent           sql.NullInt64   `gorm:"column:IsPresent" json:"IsPresent"`
	IsSalesPromotion    sql.NullInt64   `gorm:"column:IsSalesPromotion" json:"IsSalesPromotion"`
	ItemCode            sql.NullString  `gorm:"column:ItemCode" json:"ItemCode"`
	ItemName            sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	JobType             string          `gorm:"column:JobType" json:"JobType"`
	Location            sql.NullString  `gorm:"column:Location" json:"Location"`
	MainWorker          sql.NullString  `gorm:"column:MainWorker" json:"MainWorker"`
	MainWorkerCode      sql.NullString  `gorm:"column:MainWorkerCode" json:"MainWorkerCode"`
	MemberSetDiffAmount float64         `gorm:"column:MemberSetDiffAmount" json:"MemberSetDiffAmount"`
	MemberSetDiffPrice  float64         `gorm:"column:MemberSetDiffPrice" json:"MemberSetDiffPrice"`
	MemberSetItemCode   string          `gorm:"column:MemberSetItemCode" json:"MemberSetItemCode"`
	MemberSetPrice      float64         `gorm:"column:MemberSetPrice" json:"MemberSetPrice"`
	MemberSetRemark     string          `gorm:"column:MemberSetRemark" json:"MemberSetRemark"`
	MemberSetSalesCode  sql.NullString  `gorm:"column:MemberSetSalesCode" json:"MemberSetSalesCode"`
	NoTaxCostPrice      sql.NullFloat64 `gorm:"column:NoTaxCostPrice" json:"NoTaxCostPrice"`
	OldPrice            sql.NullFloat64 `gorm:"column:OldPrice" json:"OldPrice"`
	OrderID             sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	PYCode              sql.NullString  `gorm:"column:PYCode" json:"PYCode"`
	PrefPrice1          sql.NullFloat64 `gorm:"column:PrefPrice1" json:"PrefPrice1"`
	PrefPrice10         sql.NullFloat64 `gorm:"column:PrefPrice10" json:"PrefPrice10"`
	PrefPrice2          sql.NullFloat64 `gorm:"column:PrefPrice2" json:"PrefPrice2"`
	PrefPrice3          sql.NullFloat64 `gorm:"column:PrefPrice3" json:"PrefPrice3"`
	PrefPrice4          sql.NullFloat64 `gorm:"column:PrefPrice4" json:"PrefPrice4"`
	PrefPrice5          sql.NullFloat64 `gorm:"column:PrefPrice5" json:"PrefPrice5"`
	PrefPrice6          sql.NullFloat64 `gorm:"column:PrefPrice6" json:"PrefPrice6"`
	PrefPrice7          sql.NullFloat64 `gorm:"column:PrefPrice7" json:"PrefPrice7"`
	PrefPrice8          sql.NullFloat64 `gorm:"column:PrefPrice8" json:"PrefPrice8"`
	PrefPrice9          sql.NullFloat64 `gorm:"column:PrefPrice9" json:"PrefPrice9"`
	Price               sql.NullFloat64 `gorm:"column:Price" json:"Price"`
	PriceType           sql.NullString  `gorm:"column:PriceType" json:"PriceType"`
	Qty                 sql.NullFloat64 `gorm:"column:Qty" json:"Qty"`
	RecTypeD            string          `gorm:"column:RecTypeD" json:"RecTypeD"`
	RefPrice1           sql.NullFloat64 `gorm:"column:RefPrice1" json:"RefPrice1"`
	RefPrice2           sql.NullFloat64 `gorm:"column:RefPrice2" json:"RefPrice2"`
	RefPrice3           sql.NullFloat64 `gorm:"column:RefPrice3" json:"RefPrice3"`
	RefPrice4           sql.NullFloat64 `gorm:"column:RefPrice4" json:"RefPrice4"`
	SalesLeader         sql.NullString  `gorm:"column:SalesLeader" json:"SalesLeader"`
	SalesPromotionCode  sql.NullString  `gorm:"column:SalesPromotionCode" json:"SalesPromotionCode"`
	StandardPrice       sql.NullFloat64 `gorm:"column:StandardPrice" json:"StandardPrice"`
	StockID             sql.NullInt64   `gorm:"column:StockID" json:"StockID"`
	SupplierCode        sql.NullString  `gorm:"column:SupplierCode" json:"SupplierCode"`
	SupplierName        sql.NullString  `gorm:"column:SupplierName" json:"SupplierName"`
	TradePrice          sql.NullFloat64 `gorm:"column:TradePrice" json:"TradePrice"`
	TransferPrice       sql.NullFloat64 `gorm:"column:TransferPrice" json:"TransferPrice"`
	Unit                sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode         sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	WBCode              sql.NullString  `gorm:"column:WBCode" json:"WBCode"`
	WorkPlace           sql.NullString  `gorm:"column:WorkPlace" json:"WorkPlace"`
	WorkType            sql.NullString  `gorm:"column:WorkType" json:"WorkType"`
	Attributecode       string          `gorm:"column:attributecode" json:"attributecode"`
	Cid                 int             `gorm:"column:cid" json:"cid"`
	ID                  int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Pid                 sql.NullInt64   `gorm:"column:pid" json:"pid"`
	Status              int             `gorm:"column:status" json:"status"`
}

// TableName sets the insert table name for this struct type
func (c *CExpensep) TableName() string {
	return "c_expensep"
}
