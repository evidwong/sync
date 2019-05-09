package models

import (
	"database/sql"
	"time"
)

type SalesServer struct {
	BalanceCount       float64         `gorm:"column:BalanceCount" json:"BalanceCount"`
	BalanceValueAmount sql.NullFloat64 `gorm:"column:BalanceValueAmount" json:"BalanceValueAmount"`
	Class              sql.NullString  `gorm:"column:Class" json:"Class"`
	CustomerCode       sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	JobCode            sql.NullString  `gorm:"column:JobCode" json:"JobCode"`
	JobName            sql.NullString  `gorm:"column:JobName" json:"JobName"`
	JobType            sql.NullString  `gorm:"column:JobType" json:"JobType"`
	LimitDate          time.Time       `gorm:"column:LimitDate" json:"LimitDate"`
	MainWorker         sql.NullString  `gorm:"column:MainWorker" json:"MainWorker"`
	MemberSetComno     string          `gorm:"column:MemberSetComno" json:"MemberSetComno"`
	MemberSetName      sql.NullString  `gorm:"column:MemberSetName" json:"MemberSetName"`
	MemberSetSalesCode sql.NullString  `gorm:"column:MemberSetSalesCode" json:"MemberSetSalesCode"`
	MemberSetType      sql.NullString  `gorm:"column:MemberSetType" json:"MemberSetType"`
	ServiceCount       float64         `gorm:"column:ServiceCount" json:"ServiceCount"`
	SubWorker          sql.NullString  `gorm:"column:SubWorker" json:"SubWorker"`
	SubWorkerCode      sql.NullString  `gorm:"column:SubWorkerCode" json:"SubWorkerCode"`
	UseCount           float64         `gorm:"column:UseCount" json:"UseCount"`
	UseRegisterNo      sql.NullString  `gorm:"column:UseRegisterNo" json:"UseRegisterNo"`
	UseValueAmount     sql.NullFloat64 `gorm:"column:UseValueAmount" json:"UseValueAmount"`
	UseVehicleCode     sql.NullString  `gorm:"column:UseVehicleCode" json:"UseVehicleCode"`
	ValueAmount        sql.NullFloat64 `gorm:"column:ValueAmount" json:"ValueAmount"`
	WorkPlace          sql.NullString  `gorm:"column:WorkPlace" json:"WorkPlace"`
	Cid                sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ConsumePrice       sql.NullFloat64 `gorm:"column:consume_price" json:"consume_price"`
	Costprice          sql.NullFloat64 `gorm:"column:costprice" json:"costprice"`
	ID                 int             `gorm:"column:id" json:"id"`
	Isintegral         sql.NullInt64   `gorm:"column:isintegral" json:"isintegral"`
	MID                int             `gorm:"column:m_id" json:"m_id"`
	Servicecycle       sql.NullInt64   `gorm:"column:servicecycle" json:"servicecycle"`
	StandardPrice      sql.NullFloat64 `gorm:"column:standard_price" json:"standard_price"`
	Workhour           sql.NullInt64   `gorm:"column:workhour" json:"workhour"`
}

// TableName sets the insert table name for this struct type
func (s *SalesServer) TableName() string {
	return "sales_server"
}
