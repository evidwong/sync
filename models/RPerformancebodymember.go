package models

import (
	"database/sql"
	"time"
)

type RPerformancebodymember struct {
	AccountDate          time.Time       `gorm:"column:AccountDate" json:"AccountDate"`
	COMNo                string          `gorm:"column:COMNo" json:"COMNo"`
	Cost                 sql.NullFloat64 `gorm:"column:Cost" json:"Cost"`
	MemberSetCode        sql.NullString  `gorm:"column:MemberSetCode" json:"MemberSetCode"`
	MemberSetName        sql.NullString  `gorm:"column:MemberSetName" json:"MemberSetName"`
	MemberSetSalesCode   sql.NullString  `gorm:"column:MemberSetSalesCode" json:"MemberSetSalesCode"`
	Price                sql.NullFloat64 `gorm:"column:Price" json:"Price"`
	Profit               sql.NullFloat64 `gorm:"column:Profit" json:"Profit"`
	RegisterNo           sql.NullString  `gorm:"column:RegisterNo" json:"RegisterNo"`
	SalesLeaderAmount    sql.NullFloat64 `gorm:"column:SalesLeaderAmount" json:"SalesLeaderAmount"`
	SalesLeaderBonus     sql.NullInt64   `gorm:"column:SalesLeaderBonus" json:"SalesLeaderBonus"`
	SalesLeaderPercent   sql.NullFloat64 `gorm:"column:SalesLeaderPercent" json:"SalesLeaderPercent"`
	SalesLeaderkpiamount sql.NullFloat64 `gorm:"column:SalesLeaderkpiamount" json:"SalesLeaderkpiamount"`
	SalesLeaderkpiqty    sql.NullFloat64 `gorm:"column:SalesLeaderkpiqty" json:"SalesLeaderkpiqty"`
	SalesmanName         sql.NullString  `gorm:"column:SalesmanName" json:"SalesmanName"`
	Branch               string          `gorm:"column:branch" json:"branch"`
	Cid                  int             `gorm:"column:cid" json:"cid"`
	Class                sql.NullString  `gorm:"column:class" json:"class"`
	Company              string          `gorm:"column:company" json:"company"`
	ID                   int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	IsChange             int             `gorm:"column:isChange" json:"isChange"`
}

// TableName sets the insert table name for this struct type
func (r *RPerformancebodymember) TableName() string {
	return "r_performancebodymember"
}
