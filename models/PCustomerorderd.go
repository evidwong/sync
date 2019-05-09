package models

import "database/sql"

type PCustomerorderd struct {
	Attribute1         sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2         sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3         sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4         sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5         sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6         sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7         sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8         sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	Bin                sql.NullString  `gorm:"column:Bin" json:"Bin"`
	CostPrice          sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	DiscountRate       sql.NullFloat64 `gorm:"column:DiscountRate" json:"DiscountRate"`
	EnglishName        sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode              sql.NullString  `gorm:"column:FCode" json:"FCode"`
	FalseAttribute2    sql.NullString  `gorm:"column:FalseAttribute2" json:"FalseAttribute2"`
	FalseAttribute3    sql.NullString  `gorm:"column:FalseAttribute3" json:"FalseAttribute3"`
	FalseItemCode      sql.NullString  `gorm:"column:FalseItemCode" json:"FalseItemCode"`
	FalseItemName      sql.NullString  `gorm:"column:FalseItemName" json:"FalseItemName"`
	GroupNo            sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	IndexID            int             `gorm:"column:IndexID;primary_key" json:"IndexID;primary_key"`
	ItemName           sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	Location           sql.NullString  `gorm:"column:Location" json:"Location"`
	NoTaxCostPrice     sql.NullFloat64 `gorm:"column:NoTaxCostPrice" json:"NoTaxCostPrice"`
	OldPrice           sql.NullFloat64 `gorm:"column:OldPrice" json:"OldPrice"`
	OrderID            sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	OrderPrice         sql.NullFloat64 `gorm:"column:OrderPrice" json:"OrderPrice"`
	OutPriceAmount     sql.NullFloat64 `gorm:"column:OutPriceAmount" json:"OutPriceAmount"`
	OutQty             sql.NullFloat64 `gorm:"column:OutQty" json:"OutQty"`
	PCORemarks         sql.NullString  `gorm:"column:PCORemarks" json:"PCORemarks"`
	PCustomerOrderCode string          `gorm:"column:PCustomerOrderCode;primary_key" json:"PCustomerOrderCode;primary_key"`
	POutCode           sql.NullString  `gorm:"column:POutCode" json:"POutCode"`
	PriceType          sql.NullString  `gorm:"column:PriceType" json:"PriceType"`
	Qty                sql.NullFloat64 `gorm:"column:Qty" json:"Qty"`
	StandardPrice      sql.NullFloat64 `gorm:"column:StandardPrice" json:"StandardPrice"`
	StockID            sql.NullInt64   `gorm:"column:StockID" json:"StockID"`
	Unit               sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode        sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	Cid                sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode           string          `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (p *PCustomerorderd) TableName() string {
	return "p_customerorderd"
}
