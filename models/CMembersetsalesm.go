package models

import (
	"database/sql"
	"time"
)

type CMembersetsalesm struct {
	AccCardAmount          sql.NullFloat64 `gorm:"column:AccCardAmount" json:"AccCardAmount"`
	AccCashAmount          sql.NullFloat64 `gorm:"column:AccCashAmount" json:"AccCashAmount"`
	AccCheckAmount         sql.NullFloat64 `gorm:"column:AccCheckAmount" json:"AccCheckAmount"`
	AccOtherAmount         sql.NullFloat64 `gorm:"column:AccOtherAmount" json:"AccOtherAmount"`
	AccPaperAccount        sql.NullFloat64 `gorm:"column:AccPaperAccount" json:"AccPaperAccount"`
	AccPaperNo             sql.NullString  `gorm:"column:AccPaperNo" json:"AccPaperNo"`
	AccValueCardAmount     sql.NullFloat64 `gorm:"column:AccValueCardAmount" json:"AccValueCardAmount"`
	AccountDate            time.Time       `gorm:"column:AccountDate" json:"AccountDate"`
	AccountPersonCode      sql.NullString  `gorm:"column:AccountPersonCode" json:"AccountPersonCode"`
	AccountPersonName      sql.NullString  `gorm:"column:AccountPersonName" json:"AccountPersonName"`
	AccountRemarks         sql.NullString  `gorm:"column:AccountRemarks" json:"AccountRemarks"`
	ActPresentIntegral     sql.NullFloat64 `gorm:"column:ActPresentIntegral" json:"ActPresentIntegral"`
	ActionName             sql.NullString  `gorm:"column:ActionName" json:"ActionName"`
	ActivateTime           time.Time       `gorm:"column:ActivateTime" json:"ActivateTime"`
	ActiveDate             time.Time       `gorm:"column:ActiveDate" json:"ActiveDate"`
	AllFee                 sql.NullFloat64 `gorm:"column:AllFee" json:"AllFee"`
	ArrearAmount           sql.NullFloat64 `gorm:"column:ArrearAmount" json:"ArrearAmount"`
	ArrearDiscountAmount   sql.NullFloat64 `gorm:"column:ArrearDiscountAmount" json:"ArrearDiscountAmount"`
	ArrearPreReceiveAmount sql.NullFloat64 `gorm:"column:ArrearPreReceiveAmount" json:"ArrearPreReceiveAmount"`
	ArrearReceiveAmount    sql.NullFloat64 `gorm:"column:ArrearReceiveAmount" json:"ArrearReceiveAmount"`
	BadDebtsAmount         sql.NullFloat64 `gorm:"column:BadDebtsAmount" json:"BadDebtsAmount"`
	BadDebtsDate           time.Time       `gorm:"column:BadDebtsDate" json:"BadDebtsDate"`
	BadDebtsPersonCode     sql.NullString  `gorm:"column:BadDebtsPersonCode" json:"BadDebtsPersonCode"`
	BadDebtsPersonName     sql.NullString  `gorm:"column:BadDebtsPersonName" json:"BadDebtsPersonName"`
	BadDebtsReason         sql.NullString  `gorm:"column:BadDebtsReason" json:"BadDebtsReason"`
	BalanceAmount          sql.NullFloat64 `gorm:"column:BalanceAmount" json:"BalanceAmount"`
	BalanceValueAmount     sql.NullFloat64 `gorm:"column:BalanceValueAmount" json:"BalanceValueAmount"`
	Bank                   sql.NullString  `gorm:"column:Bank" json:"Bank"`
	BankAccount            sql.NullString  `gorm:"column:BankAccount" json:"BankAccount"`
	BonusAmount            sql.NullFloat64 `gorm:"column:BonusAmount" json:"BonusAmount"`
	CModel                 sql.NullString  `gorm:"column:CModel" json:"CModel"`
	COMNo                  sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	CanIntegralAmount      float64         `gorm:"column:CanIntegralAmount" json:"CanIntegralAmount"`
	CardBank               sql.NullString  `gorm:"column:CardBank" json:"CardBank"`
	CardBankAccount        sql.NullString  `gorm:"column:CardBankAccount" json:"CardBankAccount"`
	CardJSONData           sql.NullString  `gorm:"column:CardJsonData" json:"CardJsonData"`
	CardMenu               sql.NullString  `gorm:"column:CardMenu" json:"CardMenu"`
	CardSelect             sql.NullString  `gorm:"column:CardSelect" json:"CardSelect"`
	CardValue              sql.NullString  `gorm:"column:CardValue" json:"CardValue"`
	CouponCardTimeID       string          `gorm:"column:CouponCardTimeID" json:"CouponCardTimeID"`
	CouponFee              sql.NullFloat64 `gorm:"column:CouponFee" json:"CouponFee"`
	CustomerCode           sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName           sql.NullString  `gorm:"column:CustomerName" json:"CustomerName"`
	CustomerSource         sql.NullString  `gorm:"column:CustomerSource" json:"CustomerSource"`
	DiscountAmount         sql.NullFloat64 `gorm:"column:DiscountAmount" json:"DiscountAmount"`
	DisuseCode             string          `gorm:"column:DisuseCode" json:"DisuseCode"`
	DisuseDate             time.Time       `gorm:"column:DisuseDate" json:"DisuseDate"`
	DisuseName             string          `gorm:"column:DisuseName" json:"DisuseName"`
	DisuseReason           string          `gorm:"column:DisuseReason" json:"DisuseReason"`
	EffectiveDay           sql.NullInt64   `gorm:"column:EffectiveDay" json:"EffectiveDay"`
	FeeLog                 sql.NullString  `gorm:"column:FeeLog" json:"FeeLog"`
	HandPhone              sql.NullString  `gorm:"column:HandPhone" json:"HandPhone"`
	ISShared               int             `gorm:"column:IS_Shared" json:"IS_Shared"`
	IntegralAmount         sql.NullFloat64 `gorm:"column:IntegralAmount" json:"IntegralAmount"`
	IntegralCardCode       sql.NullString  `gorm:"column:IntegralCardCode" json:"IntegralCardCode"`
	IntegralCardNo         sql.NullString  `gorm:"column:IntegralCardNo" json:"IntegralCardNo"`
	InvoiceDate            time.Time       `gorm:"column:InvoiceDate" json:"InvoiceDate"`
	InvoiceNo              sql.NullString  `gorm:"column:InvoiceNo" json:"InvoiceNo"`
	InvoiceType            sql.NullString  `gorm:"column:InvoiceType" json:"InvoiceType"`
	IsAccount              sql.NullInt64   `gorm:"column:IsAccount" json:"IsAccount"`
	IsActive               sql.NullInt64   `gorm:"column:IsActive" json:"IsActive"`
	IsBadDebts             sql.NullInt64   `gorm:"column:IsBadDebts" json:"IsBadDebts"`
	IsCanIntegral          sql.NullInt64   `gorm:"column:IsCanIntegral" json:"IsCanIntegral"`
	IsDisuse               int             `gorm:"column:IsDisuse" json:"IsDisuse"`
	IsUse                  sql.NullInt64   `gorm:"column:IsUse" json:"IsUse"`
	LastUseDate            time.Time       `gorm:"column:LastUseDate" json:"LastUseDate"`
	LimitDate              time.Time       `gorm:"column:LimitDate" json:"LimitDate"`
	Linkman                sql.NullString  `gorm:"column:Linkman" json:"Linkman"`
	MemberCardNo           sql.NullString  `gorm:"column:MemberCardNo" json:"MemberCardNo"`
	MemberCode             sql.NullString  `gorm:"column:MemberCode" json:"MemberCode"`
	MemberSetCardNo        sql.NullString  `gorm:"column:MemberSetCardNo" json:"MemberSetCardNo"`
	MemberSetCode          sql.NullString  `gorm:"column:MemberSetCode" json:"MemberSetCode"`
	MemberSetComno         string          `gorm:"column:MemberSetComno" json:"MemberSetComno"`
	MemberSetID            int             `gorm:"column:MemberSetID" json:"MemberSetID"`
	MemberSetName          sql.NullString  `gorm:"column:MemberSetName" json:"MemberSetName"`
	MemberSetSalesCode     sql.NullString  `gorm:"column:MemberSetSalesCode" json:"MemberSetSalesCode"`
	MemberSetType          sql.NullString  `gorm:"column:MemberSetType" json:"MemberSetType"`
	MemberType             sql.NullString  `gorm:"column:MemberType" json:"MemberType"`
	MemberYears            sql.NullInt64   `gorm:"column:MemberYears" json:"MemberYears"`
	ModifyDate             time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP               sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode       sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName       sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation      sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	ORPrice                sql.NullFloat64 `gorm:"column:ORPrice" json:"ORPrice"`
	OldCustomerCode        sql.NullString  `gorm:"column:OldCustomerCode" json:"OldCustomerCode"`
	OldVehicleCode         sql.NullString  `gorm:"column:OldVehicleCode" json:"OldVehicleCode"`
	OtherAccType           sql.NullString  `gorm:"column:OtherAccType" json:"OtherAccType"`
	Password               sql.NullString  `gorm:"column:Password" json:"Password"`
	Phone                  sql.NullString  `gorm:"column:Phone" json:"Phone"`
	PreReceiveAmount       sql.NullFloat64 `gorm:"column:PreReceiveAmount" json:"PreReceiveAmount"`
	PresentIntegral        sql.NullFloat64 `gorm:"column:PresentIntegral" json:"PresentIntegral"`
	PresentJob             sql.NullInt64   `gorm:"column:PresentJob" json:"PresentJob"`
	Price                  sql.NullFloat64 `gorm:"column:Price" json:"Price"`
	ReceivableAmount       sql.NullFloat64 `gorm:"column:ReceivableAmount" json:"ReceivableAmount"`
	RegisterNo             sql.NullString  `gorm:"column:RegisterNo" json:"RegisterNo"`
	Remarks                sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	SalesDate              time.Time       `gorm:"column:SalesDate" json:"SalesDate"`
	SalesmanCode           sql.NullString  `gorm:"column:SalesmanCode" json:"SalesmanCode"`
	SalesmanName           sql.NullString  `gorm:"column:SalesmanName" json:"SalesmanName"`
	SetMenu                sql.NullString  `gorm:"column:SetMenu" json:"SetMenu"`
	SetPassWorD            sql.NullString  `gorm:"column:SetPassWorD" json:"SetPassWorD"`
	SetSource              sql.NullInt64   `gorm:"column:SetSource" json:"SetSource"`
	SetValue               sql.NullString  `gorm:"column:SetValue" json:"SetValue"`
	Status                 sql.NullString  `gorm:"column:Status" json:"Status"`
	SumCostAmount          sql.NullFloat64 `gorm:"column:SumCostAmount" json:"SumCostAmount"`
	UseIntegral            sql.NullFloat64 `gorm:"column:UseIntegral" json:"UseIntegral"`
	UseValueAmount         sql.NullFloat64 `gorm:"column:UseValueAmount" json:"UseValueAmount"`
	ValueAmount            sql.NullFloat64 `gorm:"column:ValueAmount" json:"ValueAmount"`
	ValueCardNo            sql.NullString  `gorm:"column:ValueCardNo" json:"ValueCardNo"`
	VehicleCode            sql.NullString  `gorm:"column:VehicleCode" json:"VehicleCode"`
	Cid                    sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ID                     int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Membersetype           sql.NullString  `gorm:"column:membersetype" json:"membersetype"`
	OutDate                time.Time       `gorm:"column:OutDate" json:"OutDate"`
	OutReason              sql.NullString  `gorm:"column:OutReason" json:"OutReason"`
	OutPayType             sql.NullString  `gorm:"column:OutPayType" json:"OutPayType"`
	OutAmount              sql.NullFloat64 `gorm:"column:OutAmount" json:"OutAmount"`
}

// TableName sets the insert table name for this struct type
func (c *CMembersetsalesm) TableName() string {
	return "c_membersetsalesm"
}
