package models

import (
	"database/sql"
	"time"
)

type RepairAnalysis struct {
	Amount          sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	Brand           sql.NullString  `gorm:"column:Brand" json:"Brand"`
	CostPrice       float64         `gorm:"column:CostPrice" json:"CostPrice"`
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
	RecTypeD        sql.NullString  `gorm:"column:RecTypeD" json:"RecTypeD"`
	ReceptionSource sql.NullString  `gorm:"column:ReceptionSource" json:"ReceptionSource"`
	VehicleCode     sql.NullString  `gorm:"column:VehicleCode" json:"VehicleCode"`
	VehicleLevel    sql.NullString  `gorm:"column:VehicleLevel" json:"VehicleLevel"`
	Cid             int             `gorm:"column:cid" json:"cid"`
	Class           sql.NullString  `gorm:"column:class" json:"class"`
	Comno           string          `gorm:"column:comno" json:"comno"`
	Registerno      sql.NullString  `gorm:"column:registerno" json:"registerno"`
	Type            int64           `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (r *RepairAnalysis) TableName() string {
	return "repair_analysis"
}
