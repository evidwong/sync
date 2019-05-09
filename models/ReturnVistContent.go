package models

import (
	"database/sql"
	"time"
)

type ReturnVistContent struct {
	ComNo             sql.NullString `gorm:"column:ComNo" json:"ComNo"`
	CustomerCode      sql.NullString `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName      sql.NullString `gorm:"column:CustomerName" json:"CustomerName"`
	RegisterNo        sql.NullString `gorm:"column:RegisterNo" json:"RegisterNo"`
	VehicleCode       sql.NullString `gorm:"column:VehicleCode" json:"VehicleCode"`
	Vehicleid         sql.NullInt64  `gorm:"column:Vehicleid" json:"Vehicleid"`
	AccessPerson      sql.NullString `gorm:"column:accessPerson" json:"accessPerson"`
	AddDate           time.Time      `gorm:"column:add_date" json:"add_date"`
	Aid               sql.NullInt64  `gorm:"column:aid" json:"aid"`
	Auid              sql.NullInt64  `gorm:"column:auid" json:"auid"`
	Awokeid           sql.NullInt64  `gorm:"column:awokeid" json:"awokeid"`
	Cid               sql.NullInt64  `gorm:"column:cid" json:"cid"`
	ClerkName         sql.NullString `gorm:"column:clerkName" json:"clerkName"`
	Clerkid           sql.NullInt64  `gorm:"column:clerkid" json:"clerkid"`
	FailContent       sql.NullString `gorm:"column:failContent" json:"failContent"`
	Feedback          sql.NullString `gorm:"column:feedback" json:"feedback"`
	Human             sql.NullString `gorm:"column:human" json:"human"`
	ID                int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel             sql.NullInt64  `gorm:"column:isdel" json:"isdel"`
	Isfail            sql.NullInt64  `gorm:"column:isfail" json:"isfail"`
	Isover            sql.NullInt64  `gorm:"column:isover" json:"isover"`
	Phone             sql.NullString `gorm:"column:phone" json:"phone"`
	Planid            sql.NullInt64  `gorm:"column:planid" json:"planid"`
	Remarks           sql.NullString `gorm:"column:remarks" json:"remarks"`
	ReturnContent     sql.NullString `gorm:"column:returnContent" json:"returnContent"`
	ReturnDate        time.Time      `gorm:"column:returnDate" json:"returnDate"`
	ReturnIntent      sql.NullString `gorm:"column:returnIntent" json:"returnIntent"`
	ReturnPerson      sql.NullString `gorm:"column:returnPerson" json:"returnPerson"`
	ReturnPlanContent sql.NullString `gorm:"column:returnPlanContent" json:"returnPlanContent"`
	ReturnPlanDate    time.Time      `gorm:"column:returnPlanDate" json:"returnPlanDate"`
	ReturnPlanType    sql.NullInt64  `gorm:"column:returnPlanType" json:"returnPlanType"`
	ReturnType        sql.NullString `gorm:"column:returnType" json:"returnType"`
	Status            sql.NullInt64  `gorm:"column:status" json:"status"`
	UID               sql.NullInt64  `gorm:"column:uid" json:"uid"`
	UpdateDate        time.Time      `gorm:"column:update_date" json:"update_date"`
	VisitResult       sql.NullString `gorm:"column:visitResult" json:"visitResult"`
	VisitResultID     sql.NullInt64  `gorm:"column:visitResultId" json:"visitResultId"`
}

// TableName sets the insert table name for this struct type
func (r *ReturnVistContent) TableName() string {
	return "return_vist_content"
}
