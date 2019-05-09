package models

import "database/sql"

type CMembersetb struct {
	ComNo                  sql.NullString  `gorm:"column:ComNo" json:"ComNo"`
	MainWorkerAmount       sql.NullFloat64 `gorm:"column:MainWorkerAmount" json:"MainWorkerAmount"`
	MainWorkerBonus        sql.NullInt64   `gorm:"column:MainWorkerBonus" json:"MainWorkerBonus"`
	MainWorkerPercent      sql.NullFloat64 `gorm:"column:MainWorkerPercent" json:"MainWorkerPercent"`
	MainWorkerkpiProfit    sql.NullFloat64 `gorm:"column:MainWorkerkpiProfit" json:"MainWorkerkpiProfit"`
	MainWorkerkpiamount    sql.NullFloat64 `gorm:"column:MainWorkerkpiamount" json:"MainWorkerkpiamount"`
	MainWorkerkpiqty       sql.NullFloat64 `gorm:"column:MainWorkerkpiqty" json:"MainWorkerkpiqty"`
	MemberSetCode          sql.NullString  `gorm:"column:MemberSetCode" json:"MemberSetCode"`
	MemberSetName          sql.NullString  `gorm:"column:MemberSetName" json:"MemberSetName"`
	RecPersonNameAmount    sql.NullFloat64 `gorm:"column:RecPersonNameAmount" json:"RecPersonNameAmount"`
	RecPersonNameBonus     sql.NullInt64   `gorm:"column:RecPersonNameBonus" json:"RecPersonNameBonus"`
	RecPersonNamePercent   sql.NullFloat64 `gorm:"column:RecPersonNamePercent" json:"RecPersonNamePercent"`
	RecPersonNamekpiProfit sql.NullFloat64 `gorm:"column:RecPersonNamekpiProfit" json:"RecPersonNamekpiProfit"`
	RecPersonNamekpiamount sql.NullFloat64 `gorm:"column:RecPersonNamekpiamount" json:"RecPersonNamekpiamount"`
	RecPersonNamekpiqty    sql.NullFloat64 `gorm:"column:RecPersonNamekpiqty" json:"RecPersonNamekpiqty"`
	SalesLeaderAmount      sql.NullFloat64 `gorm:"column:SalesLeaderAmount" json:"SalesLeaderAmount"`
	SalesLeaderBonus       sql.NullInt64   `gorm:"column:SalesLeaderBonus" json:"SalesLeaderBonus"`
	SalesLeaderPercent     sql.NullFloat64 `gorm:"column:SalesLeaderPercent" json:"SalesLeaderPercent"`
	SalesLeaderkpiProfit   sql.NullFloat64 `gorm:"column:SalesLeaderkpiProfit" json:"SalesLeaderkpiProfit"`
	SalesLeaderkpiamount   sql.NullFloat64 `gorm:"column:SalesLeaderkpiamount" json:"SalesLeaderkpiamount"`
	SalesLeaderkpiqty      sql.NullFloat64 `gorm:"column:SalesLeaderkpiqty" json:"SalesLeaderkpiqty"`
	Cid                    sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ID                     int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Pid                    sql.NullInt64   `gorm:"column:pid" json:"pid"`
}

// TableName sets the insert table name for this struct type
func (c *CMembersetb) TableName() string {
	return "c_membersetb"
}
