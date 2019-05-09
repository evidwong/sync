package models

import (
	"database/sql"
	"time"
)

type CReceptiond struct {
	Amount             sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	AwokeListCode      sql.NullString  `gorm:"column:AwokeListCode" json:"AwokeListCode"`
	BalanceCount       float64         `gorm:"column:BalanceCount" json:"BalanceCount"`
	BonusAmount        sql.NullFloat64 `gorm:"column:BonusAmount" json:"BonusAmount"`
	BonusPercent       sql.NullFloat64 `gorm:"column:BonusPercent" json:"BonusPercent"`
	CIWPCode           sql.NullString  `gorm:"column:CIWPCode" json:"CIWPCode"`
	CReceptionCode     string          `gorm:"column:CReceptionCode" json:"CReceptionCode"`
	ChargeAmount       sql.NullFloat64 `gorm:"column:ChargeAmount" json:"ChargeAmount"`
	Class              sql.NullString  `gorm:"column:Class" json:"Class"`
	CostPrice          sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	CurBalanceCount    float64         `gorm:"column:CurBalanceCount" json:"CurBalanceCount"`
	CurServiceCount    float64         `gorm:"column:CurServiceCount" json:"CurServiceCount"`
	DepartmentD        sql.NullString  `gorm:"column:DepartmentD" json:"DepartmentD"`
	DiscountRate       sql.NullFloat64 `gorm:"column:DiscountRate" json:"DiscountRate"`
	FunctionCode       string          `gorm:"column:FunctionCode" json:"FunctionCode"`
	GenAwokeListCode   sql.NullString  `gorm:"column:GenAwokeListCode" json:"GenAwokeListCode"`
	IndexID            int             `gorm:"column:IndexID" json:"IndexID"`
	Integral           float64         `gorm:"column:Integral" json:"Integral"`
	IntegralProportion float64         `gorm:"column:IntegralProportion" json:"IntegralProportion"`
	IntervalDay        sql.NullInt64   `gorm:"column:IntervalDay" json:"IntervalDay"`
	IsCalc             sql.NullInt64   `gorm:"column:IsCalc" json:"IsCalc"`
	IsCanIntegral      sql.NullInt64   `gorm:"column:IsCanIntegral" json:"IsCanIntegral"`
	IsFree             sql.NullInt64   `gorm:"column:IsFree" json:"IsFree"`
	IsMemberSet        sql.NullInt64   `gorm:"column:IsMemberSet" json:"IsMemberSet"`
	IsSalesPromotion   sql.NullInt64   `gorm:"column:IsSalesPromotion" json:"IsSalesPromotion"`
	JobCode            sql.NullString  `gorm:"column:JobCode" json:"JobCode"`
	JobName            sql.NullString  `gorm:"column:JobName" json:"JobName"`
	JobPeriod          sql.NullInt64   `gorm:"column:JobPeriod" json:"JobPeriod"`
	JobRemarks1        sql.NullString  `gorm:"column:JobRemarks1" json:"JobRemarks1"`
	JobRemarks2        sql.NullString  `gorm:"column:JobRemarks2" json:"JobRemarks2"`
	JobRemarks3        sql.NullString  `gorm:"column:JobRemarks3" json:"JobRemarks3"`
	JobStatus          sql.NullString  `gorm:"column:JobStatus" json:"JobStatus"`
	JobType            sql.NullString  `gorm:"column:JobType" json:"JobType"`
	LastServiceDate    time.Time       `gorm:"column:LastServiceDate" json:"LastServiceDate"`
	MainWorker         sql.NullString  `gorm:"column:MainWorker" json:"MainWorker"`
	MainWorkerCode     sql.NullString  `gorm:"column:MainWorkerCode" json:"MainWorkerCode"`
	MemberPrice        sql.NullFloat64 `gorm:"column:MemberPrice" json:"MemberPrice"`
	MemberSetIndex     sql.NullInt64   `gorm:"column:MemberSetIndex" json:"MemberSetIndex"`
	MemberSetSalesCode sql.NullString  `gorm:"column:MemberSetSalesCode" json:"MemberSetSalesCode"`
	NextServiceDate    time.Time       `gorm:"column:NextServiceDate" json:"NextServiceDate"`
	OldPrice           sql.NullFloat64 `gorm:"column:OldPrice" json:"OldPrice"`
	OldProPrice        sql.NullFloat64 `gorm:"column:OldProPrice" json:"OldProPrice"`
	OrderID            sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	RecTypeD           string          `gorm:"column:RecTypeD" json:"RecTypeD"`
	SalesLeader        sql.NullString  `gorm:"column:SalesLeader" json:"SalesLeader"`
	SalesLeaderCode    sql.NullString  `gorm:"column:SalesLeaderCode" json:"SalesLeaderCode"`
	SalesPromotionCode sql.NullString  `gorm:"column:SalesPromotionCode" json:"SalesPromotionCode"`
	ScheduleStartDate  time.Time       `gorm:"column:ScheduleStartDate" json:"ScheduleStartDate"`
	ScoreAmount        sql.NullFloat64 `gorm:"column:ScoreAmount" json:"ScoreAmount"`
	ServiceCount       float64         `gorm:"column:ServiceCount" json:"ServiceCount"`
	SubWorker          sql.NullString  `gorm:"column:SubWorker" json:"SubWorker"`
	SubWorkerCode      sql.NullString  `gorm:"column:SubWorkerCode" json:"SubWorkerCode"`
	TimeAmount         sql.NullFloat64 `gorm:"column:TimeAmount" json:"TimeAmount"`
	TimePrice          sql.NullFloat64 `gorm:"column:TimePrice" json:"TimePrice"`
	UseCount           float64         `gorm:"column:UseCount" json:"UseCount"`
	WorkAmount         sql.NullFloat64 `gorm:"column:WorkAmount" json:"WorkAmount"`
	WorkEndDate        time.Time       `gorm:"column:WorkEndDate" json:"WorkEndDate"`
	WorkLeaderDCode    sql.NullString  `gorm:"column:WorkLeaderDCode" json:"WorkLeaderDCode"`
	WorkLeaderDName    sql.NullString  `gorm:"column:WorkLeaderDName" json:"WorkLeaderDName"`
	WorkPlace          sql.NullString  `gorm:"column:WorkPlace" json:"WorkPlace"`
	WorkStartDate      time.Time       `gorm:"column:WorkStartDate" json:"WorkStartDate"`
	WorkTime           sql.NullFloat64 `gorm:"column:WorkTime" json:"WorkTime"`
	WorkType           sql.NullString  `gorm:"column:WorkType" json:"WorkType"`
	Cid                sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ID                 int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Pid                sql.NullInt64   `gorm:"column:pid" json:"pid"`
	Status             int             `gorm:"column:status" json:"status"`
}

// TableName sets the insert table name for this struct type
func (c *CReceptiond) TableName() string {
	return "c_receptiond"
}
