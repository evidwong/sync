package models

import (
	"database/sql"
	"time"
)

type CAwoke struct {
	Amount            sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	AwokeListCode     sql.NullString  `gorm:"column:AwokeListCode" json:"AwokeListCode"`
	AwokeType         sql.NullString  `gorm:"column:AwokeType" json:"AwokeType"`
	BonusAmount       sql.NullFloat64 `gorm:"column:BonusAmount" json:"BonusAmount"`
	BookingDate       time.Time       `gorm:"column:BookingDate" json:"BookingDate"`
	Brand             sql.NullString  `gorm:"column:Brand" json:"Brand"`
	BusinessType      sql.NullString  `gorm:"column:BusinessType" json:"BusinessType"`
	CModel            sql.NullString  `gorm:"column:CModel" json:"CModel"`
	ComNo             sql.NullString  `gorm:"column:ComNo" json:"ComNo"`
	CreateDate        time.Time       `gorm:"column:CreateDate" json:"CreateDate"`
	CustomerCode      sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName      sql.NullString  `gorm:"column:CustomerName" json:"CustomerName"`
	DeadDate          time.Time       `gorm:"column:DeadDate" json:"DeadDate"`
	DoneDate          time.Time       `gorm:"column:DoneDate" json:"DoneDate"`
	DoneReceptionCode sql.NullString  `gorm:"column:DoneReceptionCode" json:"DoneReceptionCode"`
	Factory           sql.NullString  `gorm:"column:Factory" json:"Factory"`
	FunctionCode      sql.NullString  `gorm:"column:FunctionCode" json:"FunctionCode"`
	GenReceptionCode  sql.NullString  `gorm:"column:GenReceptionCode" json:"GenReceptionCode"`
	HandPhone         sql.NullString  `gorm:"column:HandPhone" json:"HandPhone"`
	IntentType        sql.NullString  `gorm:"column:IntentType" json:"IntentType"`
	IsBooking         sql.NullInt64   `gorm:"column:IsBooking" json:"IsBooking"`
	IsNeedBooking     sql.NullInt64   `gorm:"column:IsNeedBooking" json:"IsNeedBooking"`
	JobCode           sql.NullString  `gorm:"column:JobCode" json:"JobCode"`
	JobName           sql.NullString  `gorm:"column:JobName" json:"JobName"`
	Linkman           sql.NullString  `gorm:"column:Linkman" json:"Linkman"`
	ModifyDate        time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP          sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode  sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName  sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	Phone             sql.NullString  `gorm:"column:Phone" json:"Phone"`
	PlanVisitContent  sql.NullString  `gorm:"column:PlanVisitContent" json:"PlanVisitContent"`
	PlanVisitDate     time.Time       `gorm:"column:PlanVisitDate" json:"PlanVisitDate"`
	PlanVisitType     sql.NullString  `gorm:"column:PlanVisitType" json:"PlanVisitType"`
	ReceptionSource   sql.NullString  `gorm:"column:ReceptionSource" json:"ReceptionSource"`
	RegisterNo        sql.NullString  `gorm:"column:RegisterNo" json:"RegisterNo"`
	RelaSalesCode     sql.NullString  `gorm:"column:RelaSalesCode" json:"RelaSalesCode"`
	RelaSalesman      sql.NullString  `gorm:"column:RelaSalesman" json:"RelaSalesman"`
	Remarks           sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	Status            sql.NullInt64   `gorm:"column:Status" json:"Status"`
	VehicleCode       sql.NullString  `gorm:"column:VehicleCode" json:"VehicleCode"`
	VehicleLevel      sql.NullString  `gorm:"column:VehicleLevel" json:"VehicleLevel"`
	Visitqty          sql.NullInt64   `gorm:"column:Visitqty" json:"Visitqty"`
	Writer            sql.NullString  `gorm:"column:Writer" json:"Writer"`
	Cid               sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ID                int             `gorm:"column:id;primary_key" json:"id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (c *CAwoke) TableName() string {
	return "c_awoke"
}
