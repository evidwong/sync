package models

import (
	"database/sql"
	"time"
)

type ServerAnalysis struct {
	Amount          sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	Brand           sql.NullString  `gorm:"column:Brand" json:"Brand"`
	CReceptionCode  string          `gorm:"column:CReceptionCode" json:"CReceptionCode"`
	Class           sql.NullString  `gorm:"column:Class" json:"Class"`
	CostPrice       float64         `gorm:"column:CostPrice" json:"CostPrice"`
	CurServiceCount float64         `gorm:"column:CurServiceCount" json:"CurServiceCount"`
	InvoiceType     sql.NullString  `gorm:"column:InvoiceType" json:"InvoiceType"`
	IsRecMoney      sql.NullInt64   `gorm:"column:IsRecMoney" json:"IsRecMoney"`
	JobCode         sql.NullString  `gorm:"column:JobCode" json:"JobCode"`
	JobName         sql.NullString  `gorm:"column:JobName" json:"JobName"`
	JobType         sql.NullString  `gorm:"column:JobType" json:"JobType"`
	MainWorker      sql.NullString  `gorm:"column:MainWorker" json:"MainWorker"`
	Price           float64         `gorm:"column:Price" json:"Price"`
	Qty             float64         `gorm:"column:Qty" json:"Qty"`
	RecMoneyDate    time.Time       `gorm:"column:RecMoneyDate" json:"RecMoneyDate"`
	RecPersonName   sql.NullString  `gorm:"column:RecPersonName" json:"RecPersonName"`
	RecTypeD        string          `gorm:"column:RecTypeD" json:"RecTypeD"`
	ReceptionSource sql.NullString  `gorm:"column:ReceptionSource" json:"ReceptionSource"`
	RegisterNo      sql.NullString  `gorm:"column:RegisterNo" json:"RegisterNo"`
	SalesLeader     sql.NullString  `gorm:"column:SalesLeader" json:"SalesLeader"`
	VehicleCode     sql.NullString  `gorm:"column:VehicleCode" json:"VehicleCode"`
	VehicleLevel    sql.NullString  `gorm:"column:VehicleLevel" json:"VehicleLevel"`
	WorkTime        sql.NullFloat64 `gorm:"column:WorkTime" json:"WorkTime"`
	Cid             int             `gorm:"column:cid" json:"cid"`
	Comno           string          `gorm:"column:comno" json:"comno"`
}

// TableName sets the insert table name for this struct type
func (s *ServerAnalysis) TableName() string {
	return "server_analysis"
}
