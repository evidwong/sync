package models

import (
	"database/sql"
	"time"
)

type RPerformancebody struct {
	Amount               sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	COMNo                string          `gorm:"column:COMNo" json:"COMNo"`
	CReceptionCode       string          `gorm:"column:CReceptionCode" json:"CReceptionCode"`
	CarProcess           sql.NullString  `gorm:"column:CarProcess" json:"CarProcess"`
	Cost                 sql.NullFloat64 `gorm:"column:Cost" json:"Cost"`
	CurBalanceCount      float64         `gorm:"column:CurBalanceCount" json:"CurBalanceCount"`
	InDate               time.Time       `gorm:"column:InDate" json:"InDate"`
	JobItemCode          sql.NullString  `gorm:"column:JobItemCode" json:"JobItemCode"`
	JobItemName          sql.NullString  `gorm:"column:JobItemName" json:"JobItemName"`
	MainWorker           sql.NullString  `gorm:"column:MainWorker" json:"MainWorker"`
	MainWorkerAmount     sql.NullFloat64 `gorm:"column:MainWorkerAmount" json:"MainWorkerAmount"`
	MainWorkerBonus      sql.NullInt64   `gorm:"column:MainWorkerBonus" json:"MainWorkerBonus"`
	MainWorkerPercent    sql.NullFloat64 `gorm:"column:MainWorkerPercent" json:"MainWorkerPercent"`
	MainWorkerkpiamount  sql.NullFloat64 `gorm:"column:MainWorkerkpiamount" json:"MainWorkerkpiamount"`
	MainWorkerkpiqty     sql.NullFloat64 `gorm:"column:MainWorkerkpiqty" json:"MainWorkerkpiqty"`
	Profit               sql.NullFloat64 `gorm:"column:Profit" json:"Profit"`
	RecMoneyDate         time.Time       `gorm:"column:RecMoneyDate" json:"RecMoneyDate"`
	RecType              sql.NullString  `gorm:"column:RecType" json:"RecType"`
	RegisterNo           sql.NullString  `gorm:"column:RegisterNo" json:"RegisterNo"`
	SalesLeader          sql.NullString  `gorm:"column:SalesLeader" json:"SalesLeader"`
	SalesLeaderAmount    sql.NullFloat64 `gorm:"column:SalesLeaderAmount" json:"SalesLeaderAmount"`
	SalesLeaderBonus     sql.NullInt64   `gorm:"column:SalesLeaderBonus" json:"SalesLeaderBonus"`
	SalesLeaderPercent   sql.NullFloat64 `gorm:"column:SalesLeaderPercent" json:"SalesLeaderPercent"`
	SalesLeaderkpiamount sql.NullFloat64 `gorm:"column:SalesLeaderkpiamount" json:"SalesLeaderkpiamount"`
	SalesLeaderkpiqty    sql.NullFloat64 `gorm:"column:SalesLeaderkpiqty" json:"SalesLeaderkpiqty"`
	Type                 sql.NullString  `gorm:"column:Type" json:"Type"`
	Branch               string          `gorm:"column:branch" json:"branch"`
	Cid                  int             `gorm:"column:cid" json:"cid"`
	Class                sql.NullString  `gorm:"column:class" json:"class"`
	Company              string          `gorm:"column:company" json:"company"`
	ID                   int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	IsChange             int             `gorm:"column:isChange" json:"isChange"`
}

// TableName sets the insert table name for this struct type
func (r *RPerformancebody) TableName() string {
	return "r_performancebody"
}
