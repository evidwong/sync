package models

import (
	"database/sql"
	"time"
)

type HCallinm struct {
	Address              sql.NullString  `gorm:"column:Address" json:"Address"`
	Age                  sql.NullInt64   `gorm:"column:Age" json:"Age"`
	Area                 sql.NullString  `gorm:"column:Area" json:"Area"`
	BUnitName            sql.NullString  `gorm:"column:BUnitName" json:"BUnitName"`
	Brand                sql.NullString  `gorm:"column:Brand" json:"Brand"`
	COMNo                string          `gorm:"column:COMNo" json:"COMNo"`
	CallInDate           time.Time       `gorm:"column:CallInDate" json:"CallInDate"`
	CallInType           sql.NullString  `gorm:"column:CallInType" json:"CallInType"`
	CarGroup             sql.NullString  `gorm:"column:CarGroup" json:"CarGroup"`
	City                 sql.NullString  `gorm:"column:City" json:"City"`
	Color                sql.NullString  `gorm:"column:Color" json:"Color"`
	CorpType             sql.NullString  `gorm:"column:CorpType" json:"CorpType"`
	CustomerKind         sql.NullString  `gorm:"column:CustomerKind" json:"CustomerKind"`
	CustomerName         sql.NullString  `gorm:"column:CustomerName" json:"CustomerName"`
	CustomerStatues      sql.NullString  `gorm:"column:CustomerStatues" json:"CustomerStatues"`
	DecisionAddress      sql.NullString  `gorm:"column:DecisionAddress" json:"DecisionAddress"`
	DecisionDuty         sql.NullString  `gorm:"column:DecisionDuty" json:"DecisionDuty"`
	DecisionHandPhone    sql.NullString  `gorm:"column:DecisionHandPhone" json:"DecisionHandPhone"`
	DecisionName         sql.NullString  `gorm:"column:DecisionName" json:"DecisionName"`
	DecisionPhone        sql.NullString  `gorm:"column:DecisionPhone" json:"DecisionPhone"`
	DecisionSex          sql.NullString  `gorm:"column:DecisionSex" json:"DecisionSex"`
	Degree               sql.NullString  `gorm:"column:Degree" json:"Degree"`
	Displacement         sql.NullString  `gorm:"column:Displacement" json:"Displacement"`
	DriveFeel            sql.NullString  `gorm:"column:DriveFeel" json:"DriveFeel"`
	DriveRide            sql.NullString  `gorm:"column:DriveRide" json:"DriveRide"`
	DriveRideDate        time.Time       `gorm:"column:DriveRideDate" json:"DriveRideDate"`
	DriveType            sql.NullString  `gorm:"column:DriveType" json:"DriveType"`
	Duty                 sql.NullString  `gorm:"column:Duty" json:"Duty"`
	Earning              sql.NullString  `gorm:"column:Earning" json:"Earning"`
	Email                sql.NullString  `gorm:"column:Email" json:"Email"`
	ExpectPrice          sql.NullFloat64 `gorm:"column:ExpectPrice" json:"ExpectPrice"`
	Factory              sql.NullString  `gorm:"column:Factory" json:"Factory"`
	Fancy                sql.NullString  `gorm:"column:Fancy" json:"Fancy"`
	Fax                  sql.NullString  `gorm:"column:Fax" json:"Fax"`
	FrontAxle            sql.NullString  `gorm:"column:FrontAxle" json:"FrontAxle"`
	HCOCode              sql.NullString  `gorm:"column:HCOCode" json:"HCOCode"`
	HCallInCode          string          `gorm:"column:HCallInCode" json:"HCallInCode"`
	HDriver              sql.NullString  `gorm:"column:HDriver" json:"HDriver"`
	HDriverAddress       sql.NullString  `gorm:"column:HDriverAddress" json:"HDriverAddress"`
	HDriverLevel         sql.NullString  `gorm:"column:HDriverLevel" json:"HDriverLevel"`
	HDriverPaper         sql.NullString  `gorm:"column:HDriverPaper" json:"HDriverPaper"`
	HModel               sql.NullString  `gorm:"column:HModel" json:"HModel"`
	HandPhone            sql.NullString  `gorm:"column:HandPhone" json:"HandPhone"`
	HandleDate           time.Time       `gorm:"column:HandleDate" json:"HandleDate"`
	IDCard               sql.NullString  `gorm:"column:IDCard" json:"IDCard"`
	Innern               sql.NullString  `gorm:"column:Innern" json:"Innern"`
	IntentType           sql.NullString  `gorm:"column:IntentType" json:"IntentType"`
	Introducer           sql.NullString  `gorm:"column:Introducer" json:"Introducer"`
	IntroducerCompany    sql.NullString  `gorm:"column:IntroducerCompany" json:"IntroducerCompany"`
	IntroducerHandPhone  sql.NullString  `gorm:"column:IntroducerHandPhone" json:"IntroducerHandPhone"`
	IntroducerPhone      sql.NullString  `gorm:"column:IntroducerPhone" json:"IntroducerPhone"`
	IntroducerRegisterNo sql.NullString  `gorm:"column:IntroducerRegisterNo" json:"IntroducerRegisterNo"`
	IntroducerVin        sql.NullString  `gorm:"column:IntroducerVin" json:"IntroducerVin"`
	IsLost               int             `gorm:"column:IsLost" json:"IsLost"`
	LastVisitContent     sql.NullString  `gorm:"column:LastVisitContent" json:"LastVisitContent"`
	LastVisitDate        time.Time       `gorm:"column:LastVisitDate" json:"LastVisitDate"`
	LastVisitType        sql.NullString  `gorm:"column:LastVisitType" json:"LastVisitType"`
	Linkman              sql.NullString  `gorm:"column:Linkman" json:"Linkman"`
	LostDate             time.Time       `gorm:"column:LostDate" json:"LostDate"`
	LostReason           sql.NullString  `gorm:"column:LostReason" json:"LostReason"`
	Marriage             sql.NullString  `gorm:"column:Marriage" json:"Marriage"`
	ModelCode            sql.NullString  `gorm:"column:ModelCode" json:"ModelCode"`
	ModifyDate           time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP             sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode     sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName     sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation    sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	OilCapacity          sql.NullString  `gorm:"column:OilCapacity" json:"OilCapacity"`
	OldBuyYear           sql.NullString  `gorm:"column:OldBuyYear" json:"OldBuyYear"`
	OldHModel            sql.NullString  `gorm:"column:OldHModel" json:"OldHModel"`
	OldSalesAdviser      sql.NullString  `gorm:"column:OldSalesAdviser" json:"OldSalesAdviser"`
	OldSalesAdviserCode  sql.NullString  `gorm:"column:OldSalesAdviserCode" json:"OldSalesAdviserCode"`
	OtherDemand          sql.NullString  `gorm:"column:OtherDemand" json:"OtherDemand"`
	PayType              sql.NullString  `gorm:"column:PayType" json:"PayType"`
	Phone                sql.NullString  `gorm:"column:Phone" json:"Phone"`
	PlanBuyDate          time.Time       `gorm:"column:PlanBuyDate" json:"PlanBuyDate"`
	PlanVisitContent     sql.NullString  `gorm:"column:PlanVisitContent" json:"PlanVisitContent"`
	PlanVisitDate        time.Time       `gorm:"column:PlanVisitDate" json:"PlanVisitDate"`
	PlanVisitType        sql.NullString  `gorm:"column:PlanVisitType" json:"PlanVisitType"`
	PostCode             sql.NullString  `gorm:"column:PostCode" json:"PostCode"`
	Price                sql.NullFloat64 `gorm:"column:Price" json:"Price"`
	Province             sql.NullString  `gorm:"column:Province" json:"Province"`
	Rangerover           sql.NullString  `gorm:"column:Rangerover" json:"Rangerover"`
	ReasonType           sql.NullString  `gorm:"column:ReasonType" json:"ReasonType"`
	ReceptionSource      sql.NullString  `gorm:"column:ReceptionSource" json:"ReceptionSource"`
	Remarks              sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	RidePerson           sql.NullString  `gorm:"column:RidePerson" json:"RidePerson"`
	Saddle               sql.NullString  `gorm:"column:Saddle" json:"Saddle"`
	SalesAdviser         sql.NullString  `gorm:"column:SalesAdviser" json:"SalesAdviser"`
	SalesAdviserCode     sql.NullString  `gorm:"column:SalesAdviserCode" json:"SalesAdviserCode"`
	Seat                 sql.NullInt64   `gorm:"column:Seat" json:"Seat"`
	Sex                  sql.NullString  `gorm:"column:Sex" json:"Sex"`
	SpeedRatio           sql.NullString  `gorm:"column:SpeedRatio" json:"SpeedRatio"`
	Street               sql.NullString  `gorm:"column:Street" json:"Street"`
	TireType             sql.NullString  `gorm:"column:TireType" json:"TireType"`
	TransAxle            sql.NullString  `gorm:"column:TransAxle" json:"TransAxle"`
	UrgeReason           sql.NullString  `gorm:"column:UrgeReason" json:"UrgeReason"`
	UserDefCode          sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	VehicleName          sql.NullString  `gorm:"column:VehicleName" json:"VehicleName"`
	VisitCount           sql.NullInt64   `gorm:"column:VisitCount" json:"VisitCount"`
	VisitWithPerson      sql.NullString  `gorm:"column:VisitWithPerson" json:"VisitWithPerson"`
	WorkAt               sql.NullString  `gorm:"column:WorkAt" json:"WorkAt"`
	Cid                  int             `gorm:"column:cid" json:"cid"`
	ComnoName            sql.NullString  `gorm:"column:comno_name" json:"comno_name"`
	ID                   int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	StoreID              sql.NullInt64   `gorm:"column:store_id" json:"store_id"`
}

// TableName sets the insert table name for this struct type
func (h *HCallinm) TableName() string {
	return "h_callinm"
}
