package models

import (
	"database/sql"
	"time"
)

type CMember struct {
	AccType           sql.NullString  `gorm:"column:AccType" json:"AccType"`
	BalanceIntegral   sql.NullFloat64 `gorm:"column:BalanceIntegral" json:"BalanceIntegral"`
	BonusAmount       sql.NullFloat64 `gorm:"column:BonusAmount" json:"BonusAmount"`
	COMNo             sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	CancelReason      sql.NullString  `gorm:"column:CancelReason" json:"CancelReason"`
	CustomerCode      sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName      sql.NullString  `gorm:"column:CustomerName" json:"CustomerName"`
	Integral          sql.NullFloat64 `gorm:"column:Integral" json:"Integral"`
	Introducer        sql.NullString  `gorm:"column:Introducer" json:"Introducer"`
	IsCancel          sql.NullInt64   `gorm:"column:IsCancel" json:"IsCancel"`
	JoinMemberDate    time.Time       `gorm:"column:JoinMemberDate" json:"JoinMemberDate"`
	LimitDate         time.Time       `gorm:"column:LimitDate" json:"LimitDate"`
	MemberAmount      sql.NullFloat64 `gorm:"column:MemberAmount" json:"MemberAmount"`
	MemberCardNo      sql.NullString  `gorm:"column:MemberCardNo" json:"MemberCardNo"`
	MemberCode        sql.NullString  `gorm:"column:MemberCode" json:"MemberCode"`
	MemberExitDate    time.Time       `gorm:"column:MemberExitDate" json:"MemberExitDate"`
	MemberRemarks     sql.NullString  `gorm:"column:MemberRemarks" json:"MemberRemarks"`
	MemberType        sql.NullString  `gorm:"column:MemberType" json:"MemberType"`
	MemberYears       sql.NullInt64   `gorm:"column:MemberYears" json:"MemberYears"`
	ModifyDate        time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP          sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode  sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName  sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	RegisterNo        sql.NullString  `gorm:"column:RegisterNo" json:"RegisterNo"`
	SalesCode         sql.NullString  `gorm:"column:SalesCode" json:"SalesCode"`
	SalesName         sql.NullString  `gorm:"column:SalesName" json:"SalesName"`
	Status            sql.NullString  `gorm:"column:Status" json:"Status"`
	UseIntegral       sql.NullFloat64 `gorm:"column:UseIntegral" json:"UseIntegral"`
	VehicleCode       sql.NullString  `gorm:"column:VehicleCode" json:"VehicleCode"`
	Cid               sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ID                int             `gorm:"column:id;primary_key" json:"id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (c *CMember) TableName() string {
	return "c_member"
}
