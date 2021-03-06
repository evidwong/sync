package models

import (
	"database/sql"
	"time"
)

type GVehicle struct {
	BType                     string          `gorm:"column:BType" json:"BType"`
	BUnitName                 sql.NullString  `gorm:"column:BUnitName" json:"BUnitName"`
	Brand                     sql.NullString  `gorm:"column:Brand" json:"Brand"`
	BusiFromDate              time.Time       `gorm:"column:BusiFromDate" json:"BusiFromDate"`
	BusiToDate                time.Time       `gorm:"column:BusiToDate" json:"BusiToDate"`
	BuyDate                   time.Time       `gorm:"column:BuyDate" json:"BuyDate"`
	CModel                    sql.NullString  `gorm:"column:CModel" json:"CModel"`
	CarType                   sql.NullString  `gorm:"column:CarType" json:"CarType"`
	CarWeight                 sql.NullFloat64 `gorm:"column:CarWeight" json:"CarWeight"`
	Color                     sql.NullString  `gorm:"column:Color" json:"Color"`
	ContractFromDate          time.Time       `gorm:"column:ContractFromDate" json:"ContractFromDate"`
	ContractToDate            time.Time       `gorm:"column:ContractToDate" json:"ContractToDate"`
	CreateDate                time.Time       `gorm:"column:CreateDate" json:"CreateDate"`
	CreateFilePerson          sql.NullString  `gorm:"column:CreateFilePerson" json:"CreateFilePerson"`
	CustomerCode              sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName              sql.NullString  `gorm:"column:CustomerName" json:"CustomerName"`
	DependCompany             sql.NullString  `gorm:"column:DependCompany" json:"DependCompany"`
	DependCompanyCode         sql.NullString  `gorm:"column:DependCompanyCode" json:"DependCompanyCode"`
	DependType                sql.NullString  `gorm:"column:DependType" json:"DependType"`
	Displacement              sql.NullString  `gorm:"column:Displacement" json:"Displacement"`
	DriveLic                  sql.NullString  `gorm:"column:DriveLic" json:"DriveLic"`
	DriveLicDate              time.Time       `gorm:"column:DriveLicDate" json:"DriveLicDate"`
	DriveLicType              sql.NullString  `gorm:"column:DriveLicType" json:"DriveLicType"`
	DriveLicYearCheckDate     time.Time       `gorm:"column:DriveLicYearCheckDate" json:"DriveLicYearCheckDate"`
	Driver                    sql.NullString  `gorm:"column:Driver" json:"Driver"`
	DriverBirthMD             sql.NullString  `gorm:"column:DriverBirthMD" json:"DriverBirthMD"`
	DriverDuty                sql.NullString  `gorm:"column:DriverDuty" json:"DriverDuty"`
	DriverFancy               sql.NullString  `gorm:"column:DriverFancy" json:"DriverFancy"`
	DriverHandPhone           sql.NullString  `gorm:"column:DriverHandPhone" json:"DriverHandPhone"`
	DriverMSN                 sql.NullString  `gorm:"column:DriverMSN" json:"DriverMSN"`
	DriverPhone               sql.NullString  `gorm:"column:DriverPhone" json:"DriverPhone"`
	DriverQQ                  sql.NullString  `gorm:"column:DriverQQ" json:"DriverQQ"`
	DriverSex                 sql.NullString  `gorm:"column:DriverSex" json:"DriverSex"`
	EngineNo                  sql.NullString  `gorm:"column:EngineNo" json:"EngineNo"`
	EngineType                sql.NullString  `gorm:"column:EngineType" json:"EngineType"`
	EnviromentFromDate        time.Time       `gorm:"column:EnviromentFromDate" json:"EnviromentFromDate"`
	EnviromentToDate          time.Time       `gorm:"column:EnviromentToDate" json:"EnviromentToDate"`
	Equip                     sql.NullString  `gorm:"column:Equip" json:"Equip"`
	Factory                   sql.NullString  `gorm:"column:Factory" json:"Factory"`
	FactoryCode               sql.NullString  `gorm:"column:FactoryCode" json:"FactoryCode"`
	FirstInFactoryDate        time.Time       `gorm:"column:FirstInFactoryDate" json:"FirstInFactoryDate"`
	FirstLicDate              time.Time       `gorm:"column:FirstLicDate" json:"FirstLicDate"`
	FirstMainDate             time.Time       `gorm:"column:FirstMainDate" json:"FirstMainDate"`
	FirstMainMileage          sql.NullFloat64 `gorm:"column:FirstMainMileage" json:"FirstMainMileage"`
	FromType                  sql.NullString  `gorm:"column:FromType" json:"FromType"`
	HeadPic                   sql.NullString  `gorm:"column:HeadPic" json:"HeadPic"`
	Innern                    sql.NullString  `gorm:"column:Innern" json:"Innern"`
	Insurance2Date            time.Time       `gorm:"column:Insurance2Date" json:"Insurance2Date"`
	InsuranceBillNo           sql.NullString  `gorm:"column:InsuranceBillNo" json:"InsuranceBillNo"`
	InsuranceBillNo2          sql.NullString  `gorm:"column:InsuranceBillNo2" json:"InsuranceBillNo2"`
	InsuranceCode             sql.NullString  `gorm:"column:InsuranceCode" json:"InsuranceCode"`
	InsuranceCode2            sql.NullString  `gorm:"column:InsuranceCode2" json:"InsuranceCode2"`
	InsuranceCompany          sql.NullString  `gorm:"column:InsuranceCompany" json:"InsuranceCompany"`
	InsuranceCompany2         sql.NullString  `gorm:"column:InsuranceCompany2" json:"InsuranceCompany2"`
	InsuranceDate             time.Time       `gorm:"column:InsuranceDate" json:"InsuranceDate"`
	InsuranceEffect2Date      time.Time       `gorm:"column:InsuranceEffect2Date" json:"InsuranceEffect2Date"`
	InsuranceEffectDate       time.Time       `gorm:"column:InsuranceEffectDate" json:"InsuranceEffectDate"`
	InsuranceKind             sql.NullString  `gorm:"column:InsuranceKind" json:"InsuranceKind"`
	InsuranceKind2            sql.NullString  `gorm:"column:InsuranceKind2" json:"InsuranceKind2"`
	IsOwnSales                sql.NullInt64   `gorm:"column:IsOwnSales" json:"IsOwnSales"`
	IsWXApproval              sql.NullInt64   `gorm:"column:IsWXApproval" json:"IsWXApproval"`
	LastInFactoryDate         time.Time       `gorm:"column:LastInFactoryDate" json:"LastInFactoryDate"`
	LastInFactoryEntryCode    sql.NullString  `gorm:"column:LastInFactoryEntryCode" json:"LastInFactoryEntryCode"`
	LastInFactoryFunctionCode sql.NullString  `gorm:"column:LastInFactoryFunctionCode" json:"LastInFactoryFunctionCode"`
	LastInMileage             sql.NullFloat64 `gorm:"column:LastInMileage" json:"LastInMileage"`
	LastOutFactoryDate        time.Time       `gorm:"column:LastOutFactoryDate" json:"LastOutFactoryDate"`
	LastOutMileage            sql.NullFloat64 `gorm:"column:LastOutMileage" json:"LastOutMileage"`
	MainCalcDate              time.Time       `gorm:"column:MainCalcDate" json:"MainCalcDate"`
	MainCalcEntryCode         sql.NullString  `gorm:"column:MainCalcEntryCode" json:"MainCalcEntryCode"`
	MainCalcMileage           sql.NullFloat64 `gorm:"column:MainCalcMileage" json:"MainCalcMileage"`
	MainWorker                string          `gorm:"column:MainWorker" json:"MainWorker"`
	MainWorkerCode            string          `gorm:"column:MainWorkerCode" json:"MainWorkerCode"`
	MaintainType              sql.NullString  `gorm:"column:MaintainType" json:"MaintainType"`
	MemberCardNo              string          `gorm:"column:MemberCardNo" json:"MemberCardNo"`
	MemberID                  int             `gorm:"column:MemberID" json:"MemberID"`
	MemberType                string          `gorm:"column:MemberType" json:"MemberType"`
	Mileage                   sql.NullFloat64 `gorm:"column:Mileage" json:"Mileage"`
	ModelCode                 sql.NullString  `gorm:"column:ModelCode" json:"ModelCode"`
	ModifyDate                time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP                  sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode          sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName          sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation         sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	MonthMileage              sql.NullFloat64 `gorm:"column:MonthMileage" json:"MonthMileage"`
	OldMainCalcDate           time.Time       `gorm:"column:OldMainCalcDate" json:"OldMainCalcDate"`
	OldMainCalcMileage        float32         `gorm:"column:OldMainCalcMileage" json:"OldMainCalcMileage"`
	OperatorCode              sql.NullString  `gorm:"column:OperatorCode" json:"OperatorCode"`
	OperatorName              sql.NullString  `gorm:"column:OperatorName" json:"OperatorName"`
	PlanVisitContent          sql.NullString  `gorm:"column:PlanVisitContent" json:"PlanVisitContent"`
	PlanVisitDate             time.Time       `gorm:"column:PlanVisitDate" json:"PlanVisitDate"`
	PlanVisitType             sql.NullString  `gorm:"column:PlanVisitType" json:"PlanVisitType"`
	ProFactory                sql.NullString  `gorm:"column:ProFactory" json:"ProFactory"`
	ProductYear               sql.NullString  `gorm:"column:ProductYear" json:"ProductYear"`
	PutInfo                   sql.NullString  `gorm:"column:PutInfo" json:"PutInfo"`
	Rangerover                sql.NullString  `gorm:"column:Rangerover" json:"Rangerover"`
	RegisterNo                string          `gorm:"column:RegisterNo" json:"RegisterNo"`
	RepairFromDate            time.Time       `gorm:"column:RepairFromDate" json:"RepairFromDate"`
	RepairToDate              time.Time       `gorm:"column:RepairToDate" json:"RepairToDate"`
	RunLicDate                time.Time       `gorm:"column:RunLicDate" json:"RunLicDate"`
	RunLicType                sql.NullString  `gorm:"column:RunLicType" json:"RunLicType"`
	ScrapDate                 time.Time       `gorm:"column:ScrapDate" json:"ScrapDate"`
	SeasonCheckContent        sql.NullString  `gorm:"column:SeasonCheckContent" json:"SeasonCheckContent"`
	SeasonCheckDate           time.Time       `gorm:"column:SeasonCheckDate" json:"SeasonCheckDate"`
	Seat                      sql.NullInt64   `gorm:"column:Seat" json:"Seat"`
	TransferNo                sql.NullString  `gorm:"column:TransferNo" json:"TransferNo"`
	TurnInDate                time.Time       `gorm:"column:TurnInDate" json:"TurnInDate"`
	TurnInMileage             sql.NullFloat64 `gorm:"column:TurnInMileage" json:"TurnInMileage"`
	TyreType                  string          `gorm:"column:TyreType" json:"TyreType"`
	UserDefCode               sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	VehSalesCode              string          `gorm:"column:VehSalesCode" json:"VehSalesCode"`
	VehSalesName              sql.NullString  `gorm:"column:VehSalesName" json:"VehSalesName"`
	VehicleAttr1              sql.NullString  `gorm:"column:VehicleAttr1" json:"VehicleAttr1"`
	VehicleAttr2              sql.NullString  `gorm:"column:VehicleAttr2" json:"VehicleAttr2"`
	VehicleCOMNo              sql.NullString  `gorm:"column:VehicleCOMNo" json:"VehicleCOMNo"`
	VehicleCode               sql.NullString  `gorm:"column:VehicleCode" json:"VehicleCode"`
	VehicleLevel              string          `gorm:"column:VehicleLevel" json:"VehicleLevel"`
	VehicleName               sql.NullString  `gorm:"column:VehicleName" json:"VehicleName"`
	VehicleRemarks            sql.NullString  `gorm:"column:VehicleRemarks" json:"VehicleRemarks"`
	Vin                       sql.NullString  `gorm:"column:Vin" json:"Vin"`
	WXHistoryAmount           sql.NullFloat64 `gorm:"column:WXHistoryAmount" json:"WXHistoryAmount"`
	WXHistoryCount            sql.NullInt64   `gorm:"column:WXHistoryCount" json:"WXHistoryCount"`
	YearCheckDate             time.Time       `gorm:"column:YearCheckDate" json:"YearCheckDate"`
	YearEntry                 sql.NullString  `gorm:"column:YearEntry" json:"YearEntry"`
	YearEntryDate             time.Time       `gorm:"column:YearEntryDate" json:"YearEntryDate"`
	AddTime                   sql.NullInt64   `gorm:"column:add_time" json:"add_time"`
	Aid                       sql.NullInt64   `gorm:"column:aid" json:"aid"`
	Auid                      sql.NullInt64   `gorm:"column:auid" json:"auid"`
	CarPic                    sql.NullString  `gorm:"column:carPic" json:"carPic"`
	Cid                       sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ClerkCode                 string          `gorm:"column:clerk_code" json:"clerk_code"`
	ClerkName                 sql.NullString  `gorm:"column:clerk_name" json:"clerk_name"`
	Clerkid                   sql.NullInt64   `gorm:"column:clerkid" json:"clerkid"`
	Coid                      sql.NullInt64   `gorm:"column:coid" json:"coid"`
	CommercialEndDate         time.Time       `gorm:"column:commercialEndDate" json:"commercialEndDate"`
	CommercialNo              sql.NullString  `gorm:"column:commercialNo" json:"commercialNo"`
	CommercialStartDate       time.Time       `gorm:"column:commercialStartDate" json:"commercialStartDate"`
	Comno                     sql.NullString  `gorm:"column:comno" json:"comno"`
	Company                   sql.NullString  `gorm:"column:company" json:"company"`
	Companycode               sql.NullInt64   `gorm:"column:companycode" json:"companycode"`
	CompulsoryEndDate         time.Time       `gorm:"column:compulsoryEndDate" json:"compulsoryEndDate"`
	CompulsoryNo              sql.NullString  `gorm:"column:compulsoryNo" json:"compulsoryNo"`
	CompulsoryStartDate       time.Time       `gorm:"column:compulsoryStartDate" json:"compulsoryStartDate"`
	Flagid                    int             `gorm:"column:flagid" json:"flagid"`
	ID                        int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	IDCardPic                 sql.NullString  `gorm:"column:idCardPic" json:"idCardPic"`
	IDCardPicBackPic          sql.NullString  `gorm:"column:idCardPicBackPic" json:"idCardPicBackPic"`
	InsurePic                 sql.NullString  `gorm:"column:insurePic" json:"insurePic"`
	Isdel                     int             `gorm:"column:isdel" json:"isdel"`
	LastQuotesDate            time.Time       `gorm:"column:lastQuotesDate" json:"lastQuotesDate"`
	LastQuotesTime            sql.NullInt64   `gorm:"column:lastQuotesTime" json:"lastQuotesTime"`
	Modelid                   sql.NullInt64   `gorm:"column:modelid" json:"modelid"`
	RunLicCopy                sql.NullString  `gorm:"column:runLicCopy" json:"runLicCopy"`
	RunLicPic                 sql.NullString  `gorm:"column:runLicPic" json:"runLicPic"`
	Status                    int             `gorm:"column:status" json:"status"`
	StatusTxt                 sql.NullString  `gorm:"column:statusTxt" json:"statusTxt"`
	StoreID                   sql.NullInt64   `gorm:"column:store_id" json:"store_id"`
	StoreName                 sql.NullString  `gorm:"column:store_name" json:"store_name"`
	StyleID                   sql.NullString  `gorm:"column:style_id" json:"style_id"`
	UID                       sql.NullInt64   `gorm:"column:uid" json:"uid"`
	UpdateTime                sql.NullInt64   `gorm:"column:update_time" json:"update_time"`
	WechatUser                sql.NullInt64   `gorm:"column:wechat_user" json:"wechat_user"`
}

// TableName sets the insert table name for this struct type
func (g *GVehicle) TableName() string {
	return "g_vehicle"
}
