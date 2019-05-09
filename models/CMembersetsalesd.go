package models

import (
	"database/sql"
	"time"
)

type CMembersetsalesd struct {
	BalanceCount       float64         `gorm:"column:BalanceCount" json:"BalanceCount"`
	ChargeAmount       sql.NullFloat64 `gorm:"column:ChargeAmount" json:"ChargeAmount"`
	Class              sql.NullString  `gorm:"column:Class" json:"Class"`
	ComNo              sql.NullString  `gorm:"column:ComNo" json:"ComNo"`
	IndexID            sql.NullInt64   `gorm:"column:IndexID" json:"IndexID"`
	IntervalDay        sql.NullInt64   `gorm:"column:IntervalDay" json:"IntervalDay"`
	IsAdditional       sql.NullInt64   `gorm:"column:IsAdditional" json:"IsAdditional"`
	IsPresentJob       sql.NullInt64   `gorm:"column:IsPresentJob" json:"IsPresentJob"`
	JobCode            sql.NullString  `gorm:"column:JobCode" json:"JobCode"`
	JobName            sql.NullString  `gorm:"column:JobName" json:"JobName"`
	JobType            sql.NullString  `gorm:"column:JobType" json:"JobType"`
	LastServiceDate    time.Time       `gorm:"column:LastServiceDate" json:"LastServiceDate"`
	LimitDate          time.Time       `gorm:"column:LimitDate" json:"LimitDate"`
	MainWorker         sql.NullString  `gorm:"column:MainWorker" json:"MainWorker"`
	MemberSetSalesCode sql.NullString  `gorm:"column:MemberSetSalesCode" json:"MemberSetSalesCode"`
	OrderID            sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	Remarks            sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	ServiceCount       float64         `gorm:"column:ServiceCount" json:"ServiceCount"`
	SubWorker          sql.NullString  `gorm:"column:SubWorker" json:"SubWorker"`
	SubWorkerCode      sql.NullString  `gorm:"column:SubWorkerCode" json:"SubWorkerCode"`
	UseCount           float64         `gorm:"column:UseCount" json:"UseCount"`
	UseRegisterNo      sql.NullString  `gorm:"column:UseRegisterNo" json:"UseRegisterNo"`
	UseVehicleCode     sql.NullString  `gorm:"column:UseVehicleCode" json:"UseVehicleCode"`
	WorkPlace          sql.NullString  `gorm:"column:WorkPlace" json:"WorkPlace"`
	WorkType           sql.NullString  `gorm:"column:WorkType" json:"WorkType"`
	Cid                int             `gorm:"column:cid" json:"cid"`
	ConsumePrice       sql.NullFloat64 `gorm:"column:consume_price" json:"consume_price"`
	Costprice          sql.NullFloat64 `gorm:"column:costprice" json:"costprice"`
	ID                 int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isintegral         sql.NullInt64   `gorm:"column:isintegral" json:"isintegral"`
	Servicecycle       sql.NullInt64   `gorm:"column:servicecycle" json:"servicecycle"`
	StandardPrice      sql.NullFloat64 `gorm:"column:standard_price" json:"standard_price"`
	Workhour           sql.NullInt64   `gorm:"column:workhour" json:"workhour"`
}

// TableName sets the insert table name for this struct type
func (c *CMembersetsalesd) TableName() string {
	return "c_membersetsalesd"
}
