package models

import "database/sql"

type PMonthsalesstat struct {
	ActionID             int             `gorm:"column:ActionID;primary_key" json:"ActionID;primary_key"`
	BillCount            sql.NullInt64   `gorm:"column:BillCount" json:"BillCount"`
	COMNo                string          `gorm:"column:COMNo" json:"COMNo"`
	ProfitAmount         sql.NullFloat64 `gorm:"column:ProfitAmount" json:"ProfitAmount"`
	ProfitPercent        sql.NullFloat64 `gorm:"column:ProfitPercent" json:"ProfitPercent"`
	RequireQty           sql.NullFloat64 `gorm:"column:RequireQty" json:"RequireQty"`
	TotalCostAmount      sql.NullFloat64 `gorm:"column:TotalCostAmount" json:"TotalCostAmount"`
	TotalNoTaxCostAmount sql.NullFloat64 `gorm:"column:TotalNoTaxCostAmount" json:"TotalNoTaxCostAmount"`
	TotalPriceAmount     sql.NullFloat64 `gorm:"column:TotalPriceAmount" json:"TotalPriceAmount"`
	TotalQty             sql.NullFloat64 `gorm:"column:TotalQty" json:"TotalQty"`
	YearMonth            string          `gorm:"column:YearMonth" json:"YearMonth"`
	Cid                  sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode             string          `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (p *PMonthsalesstat) TableName() string {
	return "p_monthsalesstat"
}
