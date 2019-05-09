package models

import (
	"database/sql"
	"time"
)

type ReturnVistPlan struct {
	CustomerCode      sql.NullString `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName      sql.NullString `gorm:"column:CustomerName" json:"CustomerName"`
	RegisterNo        sql.NullString `gorm:"column:RegisterNo" json:"RegisterNo"`
	VehicleCode       sql.NullString `gorm:"column:VehicleCode" json:"VehicleCode"`
	Vehicleid         sql.NullInt64  `gorm:"column:Vehicleid" json:"Vehicleid"`
	Vin               sql.NullString `gorm:"column:Vin" json:"Vin"`
	Cid               sql.NullInt64  `gorm:"column:cid" json:"cid"`
	ClerkName         sql.NullString `gorm:"column:clerkName" json:"clerkName"`
	Clerkid           sql.NullInt64  `gorm:"column:clerkid" json:"clerkid"`
	FailContent       sql.NullString `gorm:"column:failContent" json:"failContent"`
	Human             sql.NullString `gorm:"column:human" json:"human"`
	ID                int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel             sql.NullInt64  `gorm:"column:isdel" json:"isdel"`
	Isfail            sql.NullInt64  `gorm:"column:isfail" json:"isfail"`
	Isover            sql.NullInt64  `gorm:"column:isover" json:"isover"`
	Phone             sql.NullString `gorm:"column:phone" json:"phone"`
	ReturnPerson      sql.NullString `gorm:"column:returnPerson" json:"returnPerson"`
	ReturnPlanContent sql.NullString `gorm:"column:returnPlanContent" json:"returnPlanContent"`
	ReturnPlanDate    time.Time      `gorm:"column:returnPlanDate" json:"returnPlanDate"`
	ReturnPlanType    sql.NullInt64  `gorm:"column:returnPlanType" json:"returnPlanType"`
	Status            sql.NullInt64  `gorm:"column:status" json:"status"`
	UID               sql.NullInt64  `gorm:"column:uid" json:"uid"`
	WorkerID          int            `gorm:"column:worker_id" json:"worker_id"`
}

// TableName sets the insert table name for this struct type
func (r *ReturnVistPlan) TableName() string {
	return "return_vist_plan"
}
