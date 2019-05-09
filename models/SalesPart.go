package models

import (
	"database/sql"
	"time"
)

type SalesPart struct {
	Attribute2         sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute5         sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	BalanceOutQty      float64         `gorm:"column:BalanceOutQty" json:"BalanceOutQty"`
	BalanceValueAmount sql.NullFloat64 `gorm:"column:BalanceValueAmount" json:"BalanceValueAmount"`
	CanOutQty          float64         `gorm:"column:CanOutQty" json:"CanOutQty"`
	CostPrice          sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	CustomerCode       sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	ItemCode           sql.NullString  `gorm:"column:ItemCode" json:"ItemCode"`
	ItemName           sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	LimitDate          time.Time       `gorm:"column:LimitDate" json:"LimitDate"`
	MemberSetComno     string          `gorm:"column:MemberSetComno" json:"MemberSetComno"`
	MemberSetName      sql.NullString  `gorm:"column:MemberSetName" json:"MemberSetName"`
	MemberSetSalesCode sql.NullString  `gorm:"column:MemberSetSalesCode" json:"MemberSetSalesCode"`
	MemberSetType      sql.NullString  `gorm:"column:MemberSetType" json:"MemberSetType"`
	UseOutQty          float64         `gorm:"column:UseOutQty" json:"UseOutQty"`
	UseRegisterNo      sql.NullString  `gorm:"column:UseRegisterNo" json:"UseRegisterNo"`
	UseValueAmount     sql.NullFloat64 `gorm:"column:UseValueAmount" json:"UseValueAmount"`
	UseVehicleCode     sql.NullString  `gorm:"column:UseVehicleCode" json:"UseVehicleCode"`
	ValueAmount        sql.NullFloat64 `gorm:"column:ValueAmount" json:"ValueAmount"`
	Attributecode      string          `gorm:"column:attributecode" json:"attributecode"`
	Cid                sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ConsumePrice       sql.NullFloat64 `gorm:"column:consume_price" json:"consume_price"`
	ID                 int             `gorm:"column:id" json:"id"`
	Isintegral         sql.NullInt64   `gorm:"column:isintegral" json:"isintegral"`
	MID                int             `gorm:"column:m_id" json:"m_id"`
	StandardPrice      sql.NullFloat64 `gorm:"column:standard_price" json:"standard_price"`
}

// TableName sets the insert table name for this struct type
func (s *SalesPart) TableName() string {
	return "sales_part"
}
