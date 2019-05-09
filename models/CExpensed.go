package models

import (
	"database/sql"
	"time"
)

type CExpensed struct {
	Amount             sql.NullFloat64 `gorm:"column:Amount" json:"Amount"`
	AwokeListCode      sql.NullString  `gorm:"column:AwokeListCode" json:"AwokeListCode"`
	CExpenseCode       string          `gorm:"column:CExpenseCode" json:"CExpenseCode"`
	CIWPCode           sql.NullString  `gorm:"column:CIWPCode" json:"CIWPCode"`
	CReceptionCode     sql.NullString  `gorm:"column:CReceptionCode" json:"CReceptionCode"`
	Class              sql.NullString  `gorm:"column:Class" json:"Class"`
	CostPrice          sql.NullFloat64 `gorm:"column:CostPrice" json:"CostPrice"`
	CurServiceCount    sql.NullInt64   `gorm:"column:CurServiceCount" json:"CurServiceCount"`
	DepartmentD        sql.NullString  `gorm:"column:DepartmentD" json:"DepartmentD"`
	DiscountRate       sql.NullFloat64 `gorm:"column:DiscountRate" json:"DiscountRate"`
	EPRemarks          sql.NullString  `gorm:"column:EPRemarks" json:"EPRemarks"`
	FunctionCode       string          `gorm:"column:FunctionCode" json:"FunctionCode"`
	GenAwokeListCode   sql.NullString  `gorm:"column:GenAwokeListCode" json:"GenAwokeListCode"`
	IndexID            int             `gorm:"column:IndexID" json:"IndexID"`
	Integral           sql.NullFloat64 `gorm:"column:Integral" json:"Integral"`
	IntegralProportion sql.NullFloat64 `gorm:"column:IntegralProportion" json:"IntegralProportion"`
	IsCalc             sql.NullInt64   `gorm:"column:IsCalc" json:"IsCalc"`
	IsCanIntegral      sql.NullInt64   `gorm:"column:IsCanIntegral" json:"IsCanIntegral"`
	IsEPFlag           sql.NullInt64   `gorm:"column:IsEPFlag" json:"IsEPFlag"`
	IsFree             sql.NullInt64   `gorm:"column:IsFree" json:"IsFree"`
	IsMemberSet        sql.NullInt64   `gorm:"column:IsMemberSet" json:"IsMemberSet"`
	IsSalesPromotion   sql.NullInt64   `gorm:"column:IsSalesPromotion" json:"IsSalesPromotion"`
	JobCode            sql.NullString  `gorm:"column:JobCode" json:"JobCode"`
	JobName            sql.NullString  `gorm:"column:JobName" json:"JobName"`
	JobPeriod          sql.NullInt64   `gorm:"column:JobPeriod" json:"JobPeriod"`
	JobRemarks1        sql.NullString  `gorm:"column:JobRemarks1" json:"JobRemarks1"`
	JobRemarks2        sql.NullString  `gorm:"column:JobRemarks2" json:"JobRemarks2"`
	JobType            sql.NullString  `gorm:"column:JobType" json:"JobType"`
	MainWorker         sql.NullString  `gorm:"column:MainWorker" json:"MainWorker"`
	MainWorkerCode     sql.NullString  `gorm:"column:MainWorkerCode" json:"MainWorkerCode"`
	MemberPrice        sql.NullFloat64 `gorm:"column:MemberPrice" json:"MemberPrice"`
	MemberSetSalesCode sql.NullString  `gorm:"column:MemberSetSalesCode" json:"MemberSetSalesCode"`
	NextServiceDate    time.Time       `gorm:"column:NextServiceDate" json:"NextServiceDate"`
	OldPrice           sql.NullFloat64 `gorm:"column:OldPrice" json:"OldPrice"`
	OrderID            sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	RecTypeD           string          `gorm:"column:RecTypeD" json:"RecTypeD"`
	SalesLeader        sql.NullString  `gorm:"column:SalesLeader" json:"SalesLeader"`
	SalesPromotionCode sql.NullString  `gorm:"column:SalesPromotionCode" json:"SalesPromotionCode"`
	ScoreAmount        sql.NullFloat64 `gorm:"column:ScoreAmount" json:"ScoreAmount"`
	SubWorker          sql.NullString  `gorm:"column:SubWorker" json:"SubWorker"`
	TimeAmount         sql.NullFloat64 `gorm:"column:TimeAmount" json:"TimeAmount"`
	TimePrice          sql.NullFloat64 `gorm:"column:TimePrice" json:"TimePrice"`
	WorkPlace          sql.NullString  `gorm:"column:WorkPlace" json:"WorkPlace"`
	WorkTime           sql.NullFloat64 `gorm:"column:WorkTime" json:"WorkTime"`
	WorkType           sql.NullString  `gorm:"column:WorkType" json:"WorkType"`
	Cid                int             `gorm:"column:cid" json:"cid"`
	ID                 int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Pid                sql.NullInt64   `gorm:"column:pid" json:"pid"`
	Status             int             `gorm:"column:status" json:"status"`
}

// TableName sets the insert table name for this struct type
func (c *CExpensed) TableName() string {
	return "c_expensed"
}
