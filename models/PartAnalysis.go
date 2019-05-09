package models

import (
	"database/sql"
	"time"
)

type PartAnalysis struct {
	Amount          sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	Attribute1      sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2      sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3      sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute5      sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Brand           sql.NullString  `gorm:"column:Brand" json:"Brand"`
	CReceptionCode  string          `gorm:"column:CReceptionCode" json:"CReceptionCode"`
	CostPrice       float64         `gorm:"column:CostPrice" json:"CostPrice"`
	InvoiceType     sql.NullString  `gorm:"column:InvoiceType" json:"InvoiceType"`
	IsRecMoney      sql.NullInt64   `gorm:"column:IsRecMoney" json:"IsRecMoney"`
	ItemCode        sql.NullString  `gorm:"column:ItemCode" json:"ItemCode"`
	ItemName        sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	JobType         sql.NullString  `gorm:"column:JobType" json:"JobType"`
	MainWorker      sql.NullString  `gorm:"column:MainWorker" json:"MainWorker"`
	Price           float64         `gorm:"column:Price" json:"Price"`
	Qty             float64         `gorm:"column:Qty" json:"Qty"`
	RecMoneyDate    time.Time       `gorm:"column:RecMoneyDate" json:"RecMoneyDate"`
	RecPersonName   sql.NullString  `gorm:"column:RecPersonName" json:"RecPersonName"`
	RecTypeD        sql.NullString  `gorm:"column:RecTypeD" json:"RecTypeD"`
	ReceptionSource sql.NullString  `gorm:"column:ReceptionSource" json:"ReceptionSource"`
	SalesLeader     sql.NullString  `gorm:"column:SalesLeader" json:"SalesLeader"`
	VehicleCode     sql.NullString  `gorm:"column:VehicleCode" json:"VehicleCode"`
	VehicleLevel    sql.NullString  `gorm:"column:VehicleLevel" json:"VehicleLevel"`
	Cid             int             `gorm:"column:cid" json:"cid"`
	Class           sql.NullString  `gorm:"column:class" json:"class"`
	Comno           string          `gorm:"column:comno" json:"comno"`
	Registerno      sql.NullString  `gorm:"column:registerno" json:"registerno"`
}

// TableName sets the insert table name for this struct type
func (p *PartAnalysis) TableName() string {
	return "part_analysis"
}
