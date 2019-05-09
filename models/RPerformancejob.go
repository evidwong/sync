package models

import (
	"database/sql"
	"time"
)

type RPerformancejob struct {
	COMNo                  string          `gorm:"column:COMNo" json:"COMNo"`
	CReceivableAmount      sql.NullFloat64 `gorm:"column:CReceivableAmount" json:"CReceivableAmount"`
	CReceptionCode         string          `gorm:"column:CReceptionCode" json:"CReceptionCode"`
	CarProcess             sql.NullString  `gorm:"column:CarProcess" json:"CarProcess"`
	InDate                 time.Time       `gorm:"column:InDate" json:"InDate"`
	JobCode                sql.NullString  `gorm:"column:JobCode" json:"JobCode"`
	JobName                sql.NullString  `gorm:"column:JobName" json:"JobName"`
	Profit                 sql.NullFloat64 `gorm:"column:Profit" json:"Profit"`
	RecMoneyDate           time.Time       `gorm:"column:RecMoneyDate" json:"RecMoneyDate"`
	RecPersonName          sql.NullString  `gorm:"column:RecPersonName" json:"RecPersonName"`
	RecPersonNameAmount    sql.NullFloat64 `gorm:"column:RecPersonNameAmount" json:"RecPersonNameAmount"`
	RecPersonNameBonus     sql.NullInt64   `gorm:"column:RecPersonNameBonus" json:"RecPersonNameBonus"`
	RecPersonNamePercent   sql.NullFloat64 `gorm:"column:RecPersonNamePercent" json:"RecPersonNamePercent"`
	RecPersonNamekpiamount sql.NullFloat64 `gorm:"column:RecPersonNamekpiamount" json:"RecPersonNamekpiamount"`
	RecPersonNamekpiqty    sql.NullFloat64 `gorm:"column:RecPersonNamekpiqty" json:"RecPersonNamekpiqty"`
	RecType                string          `gorm:"column:RecType" json:"RecType"`
	RegisterNo             sql.NullString  `gorm:"column:RegisterNo" json:"RegisterNo"`
	SumCostAmount          sql.NullFloat64 `gorm:"column:SumCostAmount" json:"SumCostAmount"`
	Branch                 string          `gorm:"column:branch" json:"branch"`
	Cid                    int             `gorm:"column:cid" json:"cid"`
	Class                  sql.NullString  `gorm:"column:class" json:"class"`
	Company                string          `gorm:"column:company" json:"company"`
	ID                     int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	IsChange               int             `gorm:"column:isChange" json:"isChange"`
}

// TableName sets the insert table name for this struct type
func (r *RPerformancejob) TableName() string {
	return "r_performancejob"
}
