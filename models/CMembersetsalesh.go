package models

import (
	"database/sql"
	"time"
)

type CMembersetsalesh struct {
	BalanceAmount      float64        `gorm:"column:BalanceAmount" json:"BalanceAmount"`
	BalanceCount       float64        `gorm:"column:BalanceCount" json:"BalanceCount"`
	BusinessType       string         `gorm:"column:BusinessType" json:"BusinessType"`
	COMNo              string         `gorm:"column:COMNo" json:"COMNo"`
	CustomerCode       string         `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName       sql.NullString `gorm:"column:CustomerName" json:"CustomerName"`
	FunctionCode       string         `gorm:"column:FunctionCode" json:"FunctionCode"`
	MemberSetSalesCode string         `gorm:"column:MemberSetSalesCode" json:"MemberSetSalesCode"`
	Name               string         `gorm:"column:Name" json:"Name"`
	OperationType      string         `gorm:"column:OperationType" json:"OperationType"`
	OrderCode          string         `gorm:"column:OrderCode" json:"OrderCode"`
	RecPersonCode      sql.NullString `gorm:"column:RecPersonCode" json:"RecPersonCode"`
	RecPersonName      sql.NullString `gorm:"column:RecPersonName" json:"RecPersonName"`
	RegisterNo         string         `gorm:"column:RegisterNo" json:"RegisterNo"`
	SourceFunction     string         `gorm:"column:SourceFunction" json:"SourceFunction"`
	Type               int            `gorm:"column:Type" json:"Type"`
	UseAmount          float64        `gorm:"column:UseAmount" json:"UseAmount"`
	UseCount           float64        `gorm:"column:UseCount" json:"UseCount"`
	UseDate            time.Time      `gorm:"column:UseDate" json:"UseDate"`
	VehicleCode        string         `gorm:"column:VehicleCode" json:"VehicleCode"`
	Cid                int            `gorm:"column:cid" json:"cid"`
	ID                 int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Pid                int            `gorm:"column:pid" json:"pid"`
}

// TableName sets the insert table name for this struct type
func (c *CMembersetsalesh) TableName() string {
	return "c_membersetsalesh"
}
