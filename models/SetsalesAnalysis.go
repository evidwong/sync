package models

import (
	"database/sql"
	"time"
)

type SetsalesAnalysis struct {
	AccCardAmount      float64         `gorm:"column:AccCardAmount" json:"AccCardAmount"`
	AccCashAmount      float64         `gorm:"column:AccCashAmount" json:"AccCashAmount"`
	AccValueCardAmount float64         `gorm:"column:AccValueCardAmount" json:"AccValueCardAmount"`
	ArrearAmount       float64         `gorm:"column:ArrearAmount" json:"ArrearAmount"`
	DiscountAmount     float64         `gorm:"column:DiscountAmount" json:"DiscountAmount"`
	MemberSetName      sql.NullString  `gorm:"column:MemberSetName" json:"MemberSetName"`
	MemberSetSalesCode sql.NullString  `gorm:"column:MemberSetSalesCode" json:"MemberSetSalesCode"`
	MemberSetType      sql.NullString  `gorm:"column:MemberSetType" json:"MemberSetType"`
	Price              float64         `gorm:"column:Price" json:"Price"`
	RecMoneyDate       time.Time       `gorm:"column:RecMoneyDate" json:"RecMoneyDate"`
	ReceivableAmount   float64         `gorm:"column:ReceivableAmount" json:"ReceivableAmount"`
	Alipay             sql.NullFloat64 `gorm:"column:alipay" json:"alipay"`
	Cid                sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Comno              sql.NullString  `gorm:"column:comno" json:"comno"`
	Otherpay           sql.NullFloat64 `gorm:"column:otherpay" json:"otherpay"`
	Part               sql.NullFloat64 `gorm:"column:part" json:"part"`
	Server             sql.NullFloat64 `gorm:"column:server" json:"server"`
	Wepay              sql.NullFloat64 `gorm:"column:wepay" json:"wepay"`
}

// TableName sets the insert table name for this struct type
func (s *SetsalesAnalysis) TableName() string {
	return "setsales_analysis"
}
