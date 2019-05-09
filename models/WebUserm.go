package models

import (
	"database/sql"
	"time"
)

type WebUserm struct {
	Birthday          time.Time       `gorm:"column:Birthday" json:"Birthday"`
	BonusPercent      sql.NullFloat64 `gorm:"column:BonusPercent" json:"BonusPercent"`
	Class             sql.NullString  `gorm:"column:Class" json:"Class"`
	ComNo             string          `gorm:"column:ComNo" json:"ComNo"`
	ContractDate      time.Time       `gorm:"column:ContractDate" json:"ContractDate"`
	Degree            sql.NullString  `gorm:"column:Degree" json:"Degree"`
	Department        sql.NullString  `gorm:"column:Department" json:"Department"`
	Duty              sql.NullString  `gorm:"column:Duty" json:"Duty"`
	Email             sql.NullString  `gorm:"column:Email" json:"Email"`
	EmergencyContact  sql.NullString  `gorm:"column:EmergencyContact" json:"EmergencyContact"`
	EnterDate         time.Time       `gorm:"column:EnterDate" json:"EnterDate"`
	HandPhone1        sql.NullString  `gorm:"column:HandPhone1" json:"HandPhone1"`
	HandPhone2        sql.NullString  `gorm:"column:HandPhone2" json:"HandPhone2"`
	HomeAddress       sql.NullString  `gorm:"column:HomeAddress" json:"HomeAddress"`
	HomePhone         sql.NullString  `gorm:"column:HomePhone" json:"HomePhone"`
	HomePostCode      sql.NullString  `gorm:"column:HomePostCode" json:"HomePostCode"`
	IDCard            sql.NullString  `gorm:"column:IDCard" json:"IDCard"`
	IsAccountPerson   int             `gorm:"column:IsAccountPerson" json:"IsAccountPerson"`
	IsBuyer           int             `gorm:"column:IsBuyer" json:"IsBuyer"`
	IsChecker         int             `gorm:"column:IsChecker" json:"IsChecker"`
	IsCustomService   int             `gorm:"column:IsCustomService" json:"IsCustomService"`
	IsDayReport       sql.NullInt64   `gorm:"column:IsDayReport" json:"IsDayReport"`
	IsDeliveryman     int             `gorm:"column:IsDeliveryman" json:"IsDeliveryman"`
	IsDriver          int             `gorm:"column:IsDriver" json:"IsDriver"`
	IsInsurance       int             `gorm:"column:IsInsurance" json:"IsInsurance"`
	IsQC              int             `gorm:"column:IsQC" json:"IsQC"`
	IsReceive         sql.NullInt64   `gorm:"column:IsReceive" json:"IsReceive"`
	IsSalesman        int             `gorm:"column:IsSalesman" json:"IsSalesman"`
	IsSecondCarPerson int             `gorm:"column:IsSecondCarPerson" json:"IsSecondCarPerson"`
	IsService         int             `gorm:"column:IsService" json:"IsService"`
	IsTestCar         int             `gorm:"column:IsTestCar" json:"IsTestCar"`
	IsWorkChecker     int             `gorm:"column:IsWorkChecker" json:"IsWorkChecker"`
	IsWorkLeaderName  sql.NullInt64   `gorm:"column:IsWorkLeaderName" json:"IsWorkLeaderName"`
	IsWorker          int             `gorm:"column:IsWorker" json:"IsWorker"`
	JoinWorkDate      time.Time       `gorm:"column:JoinWorkDate" json:"JoinWorkDate"`
	LeaveDate         time.Time       `gorm:"column:LeaveDate" json:"LeaveDate"`
	LeaveReason       sql.NullString  `gorm:"column:LeaveReason" json:"LeaveReason"`
	Liking            sql.NullString  `gorm:"column:Liking" json:"Liking"`
	MSN               sql.NullString  `gorm:"column:MSN" json:"MSN"`
	Marriage          sql.NullString  `gorm:"column:Marriage" json:"Marriage"`
	ModifyDate        time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP          sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode  sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName  sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	NativePlace       sql.NullString  `gorm:"column:NativePlace" json:"NativePlace"`
	OfficePhone       sql.NullString  `gorm:"column:OfficePhone" json:"OfficePhone"`
	PYCode            sql.NullString  `gorm:"column:PYCode" json:"PYCode"`
	QQ                sql.NullString  `gorm:"column:QQ" json:"QQ"`
	Remarks           sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	Sex               sql.NullString  `gorm:"column:Sex" json:"Sex"`
	Skill             sql.NullString  `gorm:"column:Skill" json:"Skill"`
	Station           sql.NullString  `gorm:"column:Station" json:"Station"`
	SubWorker         sql.NullString  `gorm:"column:SubWorker" json:"SubWorker"`
	SubWorkerID       sql.NullString  `gorm:"column:SubWorker_id" json:"SubWorker_id"`
	TechnicalPost     sql.NullString  `gorm:"column:TechnicalPost" json:"TechnicalPost"`
	UserCode          string          `gorm:"column:UserCode" json:"UserCode"`
	UserName          string          `gorm:"column:UserName" json:"UserName"`
	UserPic           sql.NullString  `gorm:"column:UserPic" json:"UserPic"`
	WBCode            sql.NullString  `gorm:"column:WBCode" json:"WBCode"`
	WorkPlace         sql.NullString  `gorm:"column:WorkPlace" json:"WorkPlace"`
	WorkType          sql.NullString  `gorm:"column:WorkType" json:"WorkType"`
	Cid               sql.NullInt64   `gorm:"column:cid" json:"cid"`
	DayReportTime     sql.NullInt64   `gorm:"column:day_report_time" json:"day_report_time"`
	ID                int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel             sql.NullInt64   `gorm:"column:isdel" json:"isdel"`
	Isshow            sql.NullInt64   `gorm:"column:isshow" json:"isshow"`
	Nickname          sql.NullString  `gorm:"column:nickname" json:"nickname"`
	Openid            sql.NullString  `gorm:"column:openid" json:"openid"`
	RestTime          sql.NullString  `gorm:"column:restTime" json:"restTime"`
	Score             sql.NullFloat64 `gorm:"column:score" json:"score"`
	Sort              sql.NullInt64   `gorm:"column:sort" json:"sort"`
	StoreName         sql.NullString  `gorm:"column:store_name" json:"store_name"`
	Storeid           sql.NullInt64   `gorm:"column:storeid" json:"storeid"`
	Thumb             sql.NullString  `gorm:"column:thumb" json:"thumb"`
	Voter             sql.NullInt64   `gorm:"column:voter" json:"voter"`
}

// TableName sets the insert table name for this struct type
func (w *WebUserm) TableName() string {
	return "web_userm"
}
