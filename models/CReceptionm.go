package models

import (
	"database/sql"
	"time"
)

type CReceptionm struct {
	AccAppDate              time.Time       `gorm:"column:AccAppDate" json:"AccAppDate"`
	AccAppPersonCode        sql.NullString  `gorm:"column:AccAppPersonCode" json:"AccAppPersonCode"`
	AccAppPersonName        sql.NullString  `gorm:"column:AccAppPersonName" json:"AccAppPersonName"`
	AccCardAmount           sql.NullFloat64 `gorm:"column:AccCardAmount" json:"AccCardAmount"`
	AccCashAmount           sql.NullFloat64 `gorm:"column:AccCashAmount" json:"AccCashAmount"`
	AccCheckAmount          sql.NullFloat64 `gorm:"column:AccCheckAmount" json:"AccCheckAmount"`
	AccOtherAmount          sql.NullFloat64 `gorm:"column:AccOtherAmount" json:"AccOtherAmount"`
	AccPaperAccount         sql.NullFloat64 `gorm:"column:AccPaperAccount" json:"AccPaperAccount"`
	AccPaperNo              sql.NullString  `gorm:"column:AccPaperNo" json:"AccPaperNo"`
	AccValueCardAmount      sql.NullFloat64 `gorm:"column:AccValueCardAmount" json:"AccValueCardAmount"`
	AccountPersonCode       sql.NullString  `gorm:"column:AccountPersonCode" json:"AccountPersonCode"`
	AccountPersonName       sql.NullString  `gorm:"column:AccountPersonName" json:"AccountPersonName"`
	AccountRemarks          sql.NullString  `gorm:"column:AccountRemarks" json:"AccountRemarks"`
	ActPresentIntegral      sql.NullFloat64 `gorm:"column:ActPresentIntegral" json:"ActPresentIntegral"`
	AddInCostAmount         sql.NullFloat64 `gorm:"column:AddInCostAmount" json:"AddInCostAmount"`
	Address                 sql.NullString  `gorm:"column:Address" json:"Address"`
	AllFee                  sql.NullFloat64 `gorm:"column:AllFee" json:"AllFee"`
	Appendant               sql.NullString  `gorm:"column:Appendant" json:"Appendant"`
	ArrearAmount            sql.NullFloat64 `gorm:"column:ArrearAmount" json:"ArrearAmount"`
	ArrearDiscountAmount    sql.NullFloat64 `gorm:"column:ArrearDiscountAmount" json:"ArrearDiscountAmount"`
	ArrearPreReceiveAmount  sql.NullFloat64 `gorm:"column:ArrearPreReceiveAmount" json:"ArrearPreReceiveAmount"`
	ArrearReceiveAmount     sql.NullFloat64 `gorm:"column:ArrearReceiveAmount" json:"ArrearReceiveAmount"`
	BUnitName               sql.NullString  `gorm:"column:BUnitName" json:"BUnitName"`
	BadDebtsAmount          sql.NullFloat64 `gorm:"column:BadDebtsAmount" json:"BadDebtsAmount"`
	BadDebtsDate            time.Time       `gorm:"column:BadDebtsDate" json:"BadDebtsDate"`
	BadDebtsPersonCode      sql.NullString  `gorm:"column:BadDebtsPersonCode" json:"BadDebtsPersonCode"`
	BadDebtsPersonName      sql.NullString  `gorm:"column:BadDebtsPersonName" json:"BadDebtsPersonName"`
	BadDebtsReason          sql.NullString  `gorm:"column:BadDebtsReason" json:"BadDebtsReason"`
	BalanceAmount           sql.NullFloat64 `gorm:"column:BalanceAmount" json:"BalanceAmount"`
	BalanceIntCardIntegral  sql.NullFloat64 `gorm:"column:BalanceIntCardIntegral" json:"BalanceIntCardIntegral"`
	BalanceMemberIntegral   sql.NullFloat64 `gorm:"column:BalanceMemberIntegral" json:"BalanceMemberIntegral"`
	BalanceValueAmount      sql.NullFloat64 `gorm:"column:BalanceValueAmount" json:"BalanceValueAmount"`
	Bank                    sql.NullString  `gorm:"column:Bank" json:"Bank"`
	BankAccount             sql.NullString  `gorm:"column:BankAccount" json:"BankAccount"`
	BookingDate             time.Time       `gorm:"column:BookingDate" json:"BookingDate"`
	Brand                   sql.NullString  `gorm:"column:Brand" json:"Brand"`
	BuyDate                 time.Time       `gorm:"column:BuyDate" json:"BuyDate"`
	CAddInAmount            sql.NullFloat64 `gorm:"column:CAddInAmount" json:"CAddInAmount"`
	CAddInCostAmount        sql.NullFloat64 `gorm:"column:CAddInCostAmount" json:"CAddInCostAmount"`
	CAddInDiscountAmount    sql.NullFloat64 `gorm:"column:CAddInDiscountAmount" json:"CAddInDiscountAmount"`
	CAmount                 sql.NullFloat64 `gorm:"column:CAmount" json:"CAmount"`
	CCostAmount             sql.NullFloat64 `gorm:"column:CCostAmount" json:"CCostAmount"`
	CDiscountAmount         sql.NullFloat64 `gorm:"column:CDiscountAmount" json:"CDiscountAmount"`
	CExpenseCode            sql.NullString  `gorm:"column:CExpenseCode" json:"CExpenseCode"`
	CInsBillCode            sql.NullString  `gorm:"column:CInsBillCode" json:"CInsBillCode"`
	CJManageAmount          sql.NullFloat64 `gorm:"column:CJManageAmount" json:"CJManageAmount"`
	CJManageRate            sql.NullFloat64 `gorm:"column:CJManageRate" json:"CJManageRate"`
	CModel                  sql.NullString  `gorm:"column:CModel" json:"CModel"`
	COMNo                   string          `gorm:"column:COMNo" json:"COMNo"`
	COtherAmount            sql.NullFloat64 `gorm:"column:COtherAmount" json:"COtherAmount"`
	CPManageAmount          sql.NullFloat64 `gorm:"column:CPManageAmount" json:"CPManageAmount"`
	CPManageRate            sql.NullFloat64 `gorm:"column:CPManageRate" json:"CPManageRate"`
	CPartAmount             sql.NullFloat64 `gorm:"column:CPartAmount" json:"CPartAmount"`
	CPartCostAmount         sql.NullFloat64 `gorm:"column:CPartCostAmount" json:"CPartCostAmount"`
	CPartDiscountAmount     sql.NullFloat64 `gorm:"column:CPartDiscountAmount" json:"CPartDiscountAmount"`
	CPreReceiveAmount       sql.NullFloat64 `gorm:"column:CPreReceiveAmount" json:"CPreReceiveAmount"`
	CQuotationCode          sql.NullString  `gorm:"column:CQuotationCode" json:"CQuotationCode"`
	CReceivableAmount       sql.NullFloat64 `gorm:"column:CReceivableAmount" json:"CReceivableAmount"`
	CReceptionCode          string          `gorm:"column:CReceptionCode" json:"CReceptionCode"`
	CTimeAmount             sql.NullFloat64 `gorm:"column:CTimeAmount" json:"CTimeAmount"`
	CTimeCostAmount         sql.NullFloat64 `gorm:"column:CTimeCostAmount" json:"CTimeCostAmount"`
	CTimeDiscountAmount     sql.NullFloat64 `gorm:"column:CTimeDiscountAmount" json:"CTimeDiscountAmount"`
	CTotalAmount            sql.NullFloat64 `gorm:"column:CTotalAmount" json:"CTotalAmount"`
	CTotalDiscountAmount    sql.NullFloat64 `gorm:"column:CTotalDiscountAmount" json:"CTotalDiscountAmount"`
	CTotalFreeAmount        sql.NullFloat64 `gorm:"column:CTotalFreeAmount" json:"CTotalFreeAmount"`
	CWarrantCode            sql.NullString  `gorm:"column:CWarrantCode" json:"CWarrantCode"`
	CanIntegralAmount       sql.NullFloat64 `gorm:"column:CanIntegralAmount" json:"CanIntegralAmount"`
	CarKeyNo                sql.NullString  `gorm:"column:CarKeyNo" json:"CarKeyNo"`
	CarProcess              sql.NullString  `gorm:"column:CarProcess" json:"CarProcess"`
	CarType                 sql.NullString  `gorm:"column:CarType" json:"CarType"`
	CardBank                sql.NullString  `gorm:"column:CardBank" json:"CardBank"`
	CardBankAccount         sql.NullString  `gorm:"column:CardBankAccount" json:"CardBankAccount"`
	CardMenu                sql.NullString  `gorm:"column:CardMenu" json:"CardMenu"`
	CardPref                sql.NullString  `gorm:"column:CardPref" json:"CardPref"`
	CardSelect              sql.NullString  `gorm:"column:CardSelect" json:"CardSelect"`
	CardValue               sql.NullString  `gorm:"column:CardValue" json:"CardValue"`
	CbRecMoneyCode          sql.NullString  `gorm:"column:CbRecMoneyCode" json:"CbRecMoneyCode"`
	Color                   sql.NullString  `gorm:"column:Color" json:"Color"`
	CompleteDate            time.Time       `gorm:"column:CompleteDate" json:"CompleteDate"`
	ContractNo              sql.NullString  `gorm:"column:ContractNo" json:"ContractNo"`
	CouponCardTimeID        string          `gorm:"column:CouponCardTimeID" json:"CouponCardTimeID"`
	CouponFee               sql.NullFloat64 `gorm:"column:CouponFee" json:"CouponFee"`
	CustomBillCode          sql.NullString  `gorm:"column:CustomBillCode" json:"CustomBillCode"`
	CustomerCode            sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName            sql.NullString  `gorm:"column:CustomerName" json:"CustomerName"`
	CustomerType            sql.NullString  `gorm:"column:CustomerType" json:"CustomerType"`
	DepartmentM             sql.NullString  `gorm:"column:DepartmentM" json:"DepartmentM"`
	DiscountFee             sql.NullFloat64 `gorm:"column:DiscountFee" json:"DiscountFee"`
	DiscountPerson          sql.NullString  `gorm:"column:DiscountPerson" json:"DiscountPerson"`
	DisuseCode              string          `gorm:"column:DisuseCode" json:"DisuseCode"`
	DisuseDate              time.Time       `gorm:"column:DisuseDate" json:"DisuseDate"`
	DisuseName              string          `gorm:"column:DisuseName" json:"DisuseName"`
	DisuseReason            sql.NullString  `gorm:"column:DisuseReason" json:"DisuseReason"`
	DriveLic                sql.NullString  `gorm:"column:DriveLic" json:"DriveLic"`
	DriveMileage            sql.NullFloat64 `gorm:"column:DriveMileage" json:"DriveMileage"`
	Driver                  sql.NullString  `gorm:"column:Driver" json:"Driver"`
	DriverHandPhone         sql.NullString  `gorm:"column:DriverHandPhone" json:"DriverHandPhone"`
	DriverPhone             sql.NullString  `gorm:"column:DriverPhone" json:"DriverPhone"`
	EngineNo                sql.NullString  `gorm:"column:EngineNo" json:"EngineNo"`
	EngineType              sql.NullString  `gorm:"column:EngineType" json:"EngineType"`
	ExpectDate              time.Time       `gorm:"column:ExpectDate" json:"ExpectDate"`
	Factory                 sql.NullString  `gorm:"column:Factory" json:"Factory"`
	FactoryCode             sql.NullString  `gorm:"column:FactoryCode" json:"FactoryCode"`
	FeeLog                  sql.NullString  `gorm:"column:FeeLog" json:"FeeLog"`
	FunctionCode            sql.NullString  `gorm:"column:FunctionCode" json:"FunctionCode"`
	GuaRepaire              sql.NullInt64   `gorm:"column:GuaRepaire" json:"GuaRepaire"`
	HSalesCode              sql.NullString  `gorm:"column:HSalesCode" json:"HSalesCode"`
	HandPhone               sql.NullString  `gorm:"column:HandPhone" json:"HandPhone"`
	IAddInAmount            sql.NullFloat64 `gorm:"column:IAddInAmount" json:"IAddInAmount"`
	IAddInCostAmount        sql.NullFloat64 `gorm:"column:IAddInCostAmount" json:"IAddInCostAmount"`
	IAddInDiscountAmount    sql.NullFloat64 `gorm:"column:IAddInDiscountAmount" json:"IAddInDiscountAmount"`
	IAmount                 sql.NullFloat64 `gorm:"column:IAmount" json:"IAmount"`
	ICostAmount             sql.NullFloat64 `gorm:"column:ICostAmount" json:"ICostAmount"`
	IPartAmount             sql.NullFloat64 `gorm:"column:IPartAmount" json:"IPartAmount"`
	IPartCostAmount         sql.NullFloat64 `gorm:"column:IPartCostAmount" json:"IPartCostAmount"`
	IPartDiscountAmount     sql.NullFloat64 `gorm:"column:IPartDiscountAmount" json:"IPartDiscountAmount"`
	ITimeAmount             sql.NullFloat64 `gorm:"column:ITimeAmount" json:"ITimeAmount"`
	ITimeCostAmount         sql.NullFloat64 `gorm:"column:ITimeCostAmount" json:"ITimeCostAmount"`
	ITimeDiscountAmount     sql.NullFloat64 `gorm:"column:ITimeDiscountAmount" json:"ITimeDiscountAmount"`
	ITotalAmount            sql.NullFloat64 `gorm:"column:ITotalAmount" json:"ITotalAmount"`
	ITotalDiscountAmount    sql.NullFloat64 `gorm:"column:ITotalDiscountAmount" json:"ITotalDiscountAmount"`
	ITotalFreeAmount        sql.NullFloat64 `gorm:"column:ITotalFreeAmount" json:"ITotalFreeAmount"`
	InDate                  time.Time       `gorm:"column:InDate" json:"InDate"`
	InMileage               sql.NullFloat64 `gorm:"column:InMileage" json:"InMileage"`
	InsuranceCode           sql.NullString  `gorm:"column:InsuranceCode" json:"InsuranceCode"`
	InsuranceCompany        sql.NullString  `gorm:"column:InsuranceCompany" json:"InsuranceCompany"`
	IntegralAmount          sql.NullFloat64 `gorm:"column:IntegralAmount" json:"IntegralAmount"`
	IntegralCardCode        sql.NullString  `gorm:"column:IntegralCardCode" json:"IntegralCardCode"`
	IntegralCardNo          sql.NullString  `gorm:"column:IntegralCardNo" json:"IntegralCardNo"`
	InvoiceDate             time.Time       `gorm:"column:InvoiceDate" json:"InvoiceDate"`
	InvoiceNo               sql.NullString  `gorm:"column:InvoiceNo" json:"InvoiceNo"`
	InvoiceStatus           sql.NullString  `gorm:"column:InvoiceStatus" json:"InvoiceStatus"`
	InvoiceType             sql.NullString  `gorm:"column:InvoiceType" json:"InvoiceType"`
	IsAccApp                sql.NullInt64   `gorm:"column:IsAccApp" json:"IsAccApp"`
	IsBadDebts              sql.NullInt64   `gorm:"column:IsBadDebts" json:"IsBadDebts"`
	IsCBRecMoney            sql.NullInt64   `gorm:"column:IsCBRecMoney" json:"IsCBRecMoney"`
	IsCStart                sql.NullInt64   `gorm:"column:IsCStart" json:"IsCStart"`
	IsCarPlus               sql.NullInt64   `gorm:"column:IsCarPlus" json:"IsCarPlus"`
	IsComplete              sql.NullInt64   `gorm:"column:IsComplete" json:"IsComplete"`
	IsDelete                sql.NullInt64   `gorm:"column:IsDelete" json:"IsDelete"`
	IsDisuse                sql.NullInt64   `gorm:"column:IsDisuse" json:"IsDisuse"`
	IsOutFactory            sql.NullInt64   `gorm:"column:IsOutFactory" json:"IsOutFactory"`
	IsRecMoney              sql.NullInt64   `gorm:"column:IsRecMoney" json:"IsRecMoney"`
	IsReturnWork            sql.NullInt64   `gorm:"column:IsReturnWork" json:"IsReturnWork"`
	IsSettle                sql.NullInt64   `gorm:"column:IsSettle" json:"IsSettle"`
	IsTowTruck              sql.NullInt64   `gorm:"column:IsTowTruck" json:"IsTowTruck"`
	Linkman                 sql.NullString  `gorm:"column:Linkman" json:"Linkman"`
	MaintainDate            time.Time       `gorm:"column:MaintainDate" json:"MaintainDate"`
	MaintainKind            sql.NullString  `gorm:"column:MaintainKind" json:"MaintainKind"`
	MaintainMileage         sql.NullFloat64 `gorm:"column:MaintainMileage" json:"MaintainMileage"`
	MemberCardNo            sql.NullString  `gorm:"column:MemberCardNo" json:"MemberCardNo"`
	MemberCode              sql.NullString  `gorm:"column:MemberCode" json:"MemberCode"`
	MemberSetCardNo         sql.NullString  `gorm:"column:MemberSetCardNo" json:"MemberSetCardNo"`
	MemberType              sql.NullString  `gorm:"column:MemberType" json:"MemberType"`
	MemberYears             sql.NullInt64   `gorm:"column:MemberYears" json:"MemberYears"`
	ModifyDate              time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP                sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode        sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName        sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation       sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	NeedVisitDate           time.Time       `gorm:"column:NeedVisitDate" json:"NeedVisitDate"`
	OilScale                sql.NullString  `gorm:"column:OilScale" json:"OilScale"`
	OperatorCode            sql.NullString  `gorm:"column:OperatorCode" json:"OperatorCode"`
	OperatorName            sql.NullString  `gorm:"column:OperatorName" json:"OperatorName"`
	OtherAccType            sql.NullString  `gorm:"column:OtherAccType" json:"OtherAccType"`
	OutDate                 time.Time       `gorm:"column:OutDate" json:"OutDate"`
	OutFactoryPersonCode    sql.NullString  `gorm:"column:OutFactoryPersonCode" json:"OutFactoryPersonCode"`
	OutFactoryPersonName    sql.NullString  `gorm:"column:OutFactoryPersonName" json:"OutFactoryPersonName"`
	OutMileage              sql.NullFloat64 `gorm:"column:OutMileage" json:"OutMileage"`
	OutSatisfy              sql.NullString  `gorm:"column:OutSatisfy" json:"OutSatisfy"`
	PAddInAmount            sql.NullFloat64 `gorm:"column:PAddInAmount" json:"PAddInAmount"`
	PAddInCostAmount        sql.NullFloat64 `gorm:"column:PAddInCostAmount" json:"PAddInCostAmount"`
	PAddInDiscountAmount    sql.NullFloat64 `gorm:"column:PAddInDiscountAmount" json:"PAddInDiscountAmount"`
	PAmount                 sql.NullFloat64 `gorm:"column:PAmount" json:"PAmount"`
	PArrearAmount           sql.NullFloat64 `gorm:"column:PArrearAmount" json:"PArrearAmount"`
	PArrearDiscountAmount   sql.NullFloat64 `gorm:"column:PArrearDiscountAmount" json:"PArrearDiscountAmount"`
	PArrearPreReceiveAmount sql.NullFloat64 `gorm:"column:PArrearPreReceiveAmount" json:"PArrearPreReceiveAmount"`
	PArrearReceiveAmount    sql.NullFloat64 `gorm:"column:PArrearReceiveAmount" json:"PArrearReceiveAmount"`
	PBalanceAmount          sql.NullFloat64 `gorm:"column:PBalanceAmount" json:"PBalanceAmount"`
	PCostAmount             sql.NullFloat64 `gorm:"column:PCostAmount" json:"PCostAmount"`
	POutCode                sql.NullString  `gorm:"column:POutCode" json:"POutCode"`
	PPartAmount             sql.NullFloat64 `gorm:"column:PPartAmount" json:"PPartAmount"`
	PPartCostAmount         sql.NullFloat64 `gorm:"column:PPartCostAmount" json:"PPartCostAmount"`
	PPartDiscountAmount     sql.NullFloat64 `gorm:"column:PPartDiscountAmount" json:"PPartDiscountAmount"`
	PPayFor                 sql.NullString  `gorm:"column:PPayFor" json:"PPayFor"`
	PRejectAmount           sql.NullFloat64 `gorm:"column:PRejectAmount" json:"PRejectAmount"`
	PSettleAddInAmount      sql.NullFloat64 `gorm:"column:PSettleAddInAmount" json:"PSettleAddInAmount"`
	PSettleBillCode         sql.NullString  `gorm:"column:PSettleBillCode" json:"PSettleBillCode"`
	PSettleCPayAmount       sql.NullFloat64 `gorm:"column:PSettleCPayAmount" json:"PSettleCPayAmount"`
	PSettleDate             time.Time       `gorm:"column:PSettleDate" json:"PSettleDate"`
	PSettlePartsAmount      sql.NullFloat64 `gorm:"column:PSettlePartsAmount" json:"PSettlePartsAmount"`
	PSettleTimeAmount       sql.NullFloat64 `gorm:"column:PSettleTimeAmount" json:"PSettleTimeAmount"`
	PSettleType             sql.NullString  `gorm:"column:PSettleType" json:"PSettleType"`
	PTimeAmount             sql.NullFloat64 `gorm:"column:PTimeAmount" json:"PTimeAmount"`
	PTimeCostAmount         sql.NullFloat64 `gorm:"column:PTimeCostAmount" json:"PTimeCostAmount"`
	PTimeDiscountAmount     sql.NullFloat64 `gorm:"column:PTimeDiscountAmount" json:"PTimeDiscountAmount"`
	PTotalAmount            sql.NullFloat64 `gorm:"column:PTotalAmount" json:"PTotalAmount"`
	PTotalDiscountAmount    sql.NullFloat64 `gorm:"column:PTotalDiscountAmount" json:"PTotalDiscountAmount"`
	PTotalFreeAmount        sql.NullFloat64 `gorm:"column:PTotalFreeAmount" json:"PTotalFreeAmount"`
	PartsCostAmount         sql.NullFloat64 `gorm:"column:PartsCostAmount" json:"PartsCostAmount"`
	Phone                   sql.NullString  `gorm:"column:Phone" json:"Phone"`
	PreReceiveAmount        sql.NullFloat64 `gorm:"column:PreReceiveAmount" json:"PreReceiveAmount"`
	PresentIntegral         sql.NullFloat64 `gorm:"column:PresentIntegral" json:"PresentIntegral"`
	ProFactory              sql.NullString  `gorm:"column:ProFactory" json:"ProFactory"`
	ProductYear             sql.NullString  `gorm:"column:ProductYear" json:"ProductYear"`
	Proposer                sql.NullString  `gorm:"column:Proposer" json:"Proposer"`
	ProposerPhone           sql.NullString  `gorm:"column:ProposerPhone" json:"ProposerPhone"`
	QuotationDate           time.Time       `gorm:"column:QuotationDate" json:"QuotationDate"`
	RWReceptionCode         sql.NullString  `gorm:"column:RWReceptionCode" json:"RWReceptionCode"`
	RecDate                 time.Time       `gorm:"column:RecDate" json:"RecDate"`
	RecMoneyDate            time.Time       `gorm:"column:RecMoneyDate" json:"RecMoneyDate"`
	RecMoneyPersonCode      sql.NullString  `gorm:"column:RecMoneyPersonCode" json:"RecMoneyPersonCode"`
	RecMoneyPersonName      sql.NullString  `gorm:"column:RecMoneyPersonName" json:"RecMoneyPersonName"`
	RecPersonCode           sql.NullString  `gorm:"column:RecPersonCode" json:"RecPersonCode"`
	RecPersonName           sql.NullString  `gorm:"column:RecPersonName" json:"RecPersonName"`
	RecType                 sql.NullString  `gorm:"column:RecType" json:"RecType"`
	ReceivableIntegral      sql.NullInt64   `gorm:"column:ReceivableIntegral" json:"ReceivableIntegral"`
	ReceptionSource         sql.NullString  `gorm:"column:ReceptionSource" json:"ReceptionSource"`
	RegisterNo              sql.NullString  `gorm:"column:RegisterNo" json:"RegisterNo"`
	RelaSalesCode           sql.NullString  `gorm:"column:RelaSalesCode" json:"RelaSalesCode"`
	RelaSalesman            sql.NullString  `gorm:"column:RelaSalesman" json:"RelaSalesman"`
	Remarks                 sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	ReturnVistPlanID        int             `gorm:"column:ReturnVistPlanID" json:"ReturnVistPlanID"`
	SalesCode               sql.NullString  `gorm:"column:SalesCode" json:"SalesCode"`
	SalesLeader             sql.NullString  `gorm:"column:SalesLeader" json:"SalesLeader"`
	SalesName               sql.NullString  `gorm:"column:SalesName" json:"SalesName"`
	SelfExcept              sql.NullString  `gorm:"column:SelfExcept" json:"SelfExcept"`
	SendCarCostAmount       sql.NullFloat64 `gorm:"column:SendCarCostAmount" json:"SendCarCostAmount"`
	SendCarMileage          sql.NullFloat64 `gorm:"column:SendCarMileage" json:"SendCarMileage"`
	SendCarPerson           sql.NullString  `gorm:"column:SendCarPerson" json:"SendCarPerson"`
	SetMenu                 sql.NullString  `gorm:"column:SetMenu" json:"SetMenu"`
	SetValue                sql.NullString  `gorm:"column:SetValue" json:"SetValue"`
	SettleDate              time.Time       `gorm:"column:SettleDate" json:"SettleDate"`
	SettleRemarks           sql.NullString  `gorm:"column:SettleRemarks" json:"SettleRemarks"`
	SourceBillCode          sql.NullString  `gorm:"column:SourceBillCode" json:"SourceBillCode"`
	SpeRemarks              sql.NullString  `gorm:"column:SpeRemarks" json:"SpeRemarks"`
	SumAddInAmount          sql.NullFloat64 `gorm:"column:SumAddInAmount" json:"SumAddInAmount"`
	SumAmount               sql.NullFloat64 `gorm:"column:SumAmount" json:"SumAmount"`
	SumCostAmount           sql.NullFloat64 `gorm:"column:SumCostAmount" json:"SumCostAmount"`
	SumDiscountAmount       sql.NullFloat64 `gorm:"column:SumDiscountAmount" json:"SumDiscountAmount"`
	SumFreeAmount           sql.NullFloat64 `gorm:"column:SumFreeAmount" json:"SumFreeAmount"`
	SumPartAmount           sql.NullFloat64 `gorm:"column:SumPartAmount" json:"SumPartAmount"`
	SumPartIntegral         float64         `gorm:"column:SumPartIntegral" json:"SumPartIntegral"`
	SumQtyTotal             int             `gorm:"column:SumQtyTotal" json:"SumQtyTotal"`
	SumTimeAmount           sql.NullFloat64 `gorm:"column:SumTimeAmount" json:"SumTimeAmount"`
	SumTimeIntegral         float64         `gorm:"column:SumTimeIntegral" json:"SumTimeIntegral"`
	SumTotalAmount          sql.NullFloat64 `gorm:"column:SumTotalAmount" json:"SumTotalAmount"`
	SumTotalIntegral        float64         `gorm:"column:SumTotalIntegral" json:"SumTotalIntegral"`
	TaxAmount               sql.NullFloat64 `gorm:"column:TaxAmount" json:"TaxAmount"`
	TaxRate                 sql.NullFloat64 `gorm:"column:TaxRate" json:"TaxRate"`
	TestCarCode             sql.NullString  `gorm:"column:TestCarCode" json:"TestCarCode"`
	TestCarName             sql.NullString  `gorm:"column:TestCarName" json:"TestCarName"`
	TestDriveReport         sql.NullString  `gorm:"column:TestDriveReport" json:"TestDriveReport"`
	TimeCostAmount          sql.NullFloat64 `gorm:"column:TimeCostAmount" json:"TimeCostAmount"`
	TransferNo              sql.NullString  `gorm:"column:TransferNo" json:"TransferNo"`
	UseIntegral             sql.NullFloat64 `gorm:"column:UseIntegral" json:"UseIntegral"`
	ValueCardCOMNo          sql.NullString  `gorm:"column:ValueCardCOMNo" json:"ValueCardCOMNo"`
	ValueCardNo             sql.NullString  `gorm:"column:ValueCardNo" json:"ValueCardNo"`
	VehicleCode             sql.NullString  `gorm:"column:VehicleCode" json:"VehicleCode"`
	VehicleLevel            sql.NullString  `gorm:"column:VehicleLevel" json:"VehicleLevel"`
	VehicleName             sql.NullString  `gorm:"column:VehicleName" json:"VehicleName"`
	Vin                     sql.NullString  `gorm:"column:Vin" json:"Vin"`
	WAddInAmount            sql.NullFloat64 `gorm:"column:WAddInAmount" json:"WAddInAmount"`
	WAddInCostAmount        sql.NullFloat64 `gorm:"column:WAddInCostAmount" json:"WAddInCostAmount"`
	WAddInDiscountAmount    sql.NullFloat64 `gorm:"column:WAddInDiscountAmount" json:"WAddInDiscountAmount"`
	WAmount                 sql.NullFloat64 `gorm:"column:WAmount" json:"WAmount"`
	WArrearAmount           sql.NullFloat64 `gorm:"column:WArrearAmount" json:"WArrearAmount"`
	WArrearDiscountAmount   sql.NullFloat64 `gorm:"column:WArrearDiscountAmount" json:"WArrearDiscountAmount"`
	WArrearPreReceiveAmount sql.NullFloat64 `gorm:"column:WArrearPreReceiveAmount" json:"WArrearPreReceiveAmount"`
	WArrearReceiveAmount    sql.NullFloat64 `gorm:"column:WArrearReceiveAmount" json:"WArrearReceiveAmount"`
	WBalanceAmount          sql.NullFloat64 `gorm:"column:WBalanceAmount" json:"WBalanceAmount"`
	WCostAmount             sql.NullFloat64 `gorm:"column:WCostAmount" json:"WCostAmount"`
	WPartAmount             sql.NullFloat64 `gorm:"column:WPartAmount" json:"WPartAmount"`
	WPartCostAmount         sql.NullFloat64 `gorm:"column:WPartCostAmount" json:"WPartCostAmount"`
	WPartDiscountAmount     sql.NullFloat64 `gorm:"column:WPartDiscountAmount" json:"WPartDiscountAmount"`
	WSettleAddInAmount      sql.NullFloat64 `gorm:"column:WSettleAddInAmount" json:"WSettleAddInAmount"`
	WSettleBillCode         sql.NullString  `gorm:"column:WSettleBillCode" json:"WSettleBillCode"`
	WSettleCPayAmount       sql.NullFloat64 `gorm:"column:WSettleCPayAmount" json:"WSettleCPayAmount"`
	WSettleDate             time.Time       `gorm:"column:WSettleDate" json:"WSettleDate"`
	WSettlePartsAmount      sql.NullFloat64 `gorm:"column:WSettlePartsAmount" json:"WSettlePartsAmount"`
	WSettleTimeAmount       sql.NullFloat64 `gorm:"column:WSettleTimeAmount" json:"WSettleTimeAmount"`
	WSettleType             sql.NullString  `gorm:"column:WSettleType" json:"WSettleType"`
	WTimeAmount             sql.NullFloat64 `gorm:"column:WTimeAmount" json:"WTimeAmount"`
	WTimeCostAmount         sql.NullFloat64 `gorm:"column:WTimeCostAmount" json:"WTimeCostAmount"`
	WTimeDiscountAmount     sql.NullFloat64 `gorm:"column:WTimeDiscountAmount" json:"WTimeDiscountAmount"`
	WTotalAmount            sql.NullFloat64 `gorm:"column:WTotalAmount" json:"WTotalAmount"`
	WTotalDiscountAmount    sql.NullFloat64 `gorm:"column:WTotalDiscountAmount" json:"WTotalDiscountAmount"`
	WTotalFreeAmount        sql.NullFloat64 `gorm:"column:WTotalFreeAmount" json:"WTotalFreeAmount"`
	WorkCheckerCode         sql.NullString  `gorm:"column:WorkCheckerCode" json:"WorkCheckerCode"`
	WorkCheckerName         sql.NullString  `gorm:"column:WorkCheckerName" json:"WorkCheckerName"`
	WorkLeaderCode          sql.NullString  `gorm:"column:WorkLeaderCode" json:"WorkLeaderCode"`
	WorkLeaderName          sql.NullString  `gorm:"column:WorkLeaderName" json:"WorkLeaderName"`
	ZAddInAmount            sql.NullFloat64 `gorm:"column:ZAddInAmount" json:"ZAddInAmount"`
	ZPartAmount             sql.NullFloat64 `gorm:"column:ZPartAmount" json:"ZPartAmount"`
	ZReceivableAmount       sql.NullFloat64 `gorm:"column:ZReceivableAmount" json:"ZReceivableAmount"`
	ZTimeAmount             sql.NullFloat64 `gorm:"column:ZTimeAmount" json:"ZTimeAmount"`
	ZTotalAmount            sql.NullFloat64 `gorm:"column:ZTotalAmount" json:"ZTotalAmount"`
	ZTotalDiscountAmount    sql.NullFloat64 `gorm:"column:ZTotalDiscountAmount" json:"ZTotalDiscountAmount"`
	ZTotalFreeAmount        sql.NullFloat64 `gorm:"column:ZTotalFreeAmount" json:"ZTotalFreeAmount"`
	Cid                     int             `gorm:"column:cid" json:"cid"`
	ID                      int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	IsQuotation             int             `gorm:"column:isQuotation" json:"isQuotation"`
	IsQuotedprice           sql.NullInt64   `gorm:"column:isQuotedprice" json:"isQuotedprice"`
	IsTesting               sql.NullInt64   `gorm:"column:isTesting" json:"isTesting"`
	Pickingtype             int             `gorm:"column:pickingtype" json:"pickingtype"`
	ReserveID               int             `gorm:"column:reserve_id" json:"reserve_id"`
}

// TableName sets the insert table name for this struct type
func (c *CReceptionm) TableName() string {
	return "c_receptionm"
}
