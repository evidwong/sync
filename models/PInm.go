package models

import (
	"database/sql"
	"time"
)

type PInm struct {
	AccAppDate                  time.Time       `gorm:"column:AccAppDate" json:"AccAppDate"`
	AccAppPersonCode            sql.NullString  `gorm:"column:AccAppPersonCode" json:"AccAppPersonCode"`
	AccAppPersonName            sql.NullString  `gorm:"column:AccAppPersonName" json:"AccAppPersonName"`
	AccType                     sql.NullString  `gorm:"column:AccType" json:"AccType"`
	AccountDate                 time.Time       `gorm:"column:AccountDate" json:"AccountDate"`
	AccountPersonCode           sql.NullString  `gorm:"column:AccountPersonCode" json:"AccountPersonCode"`
	AccountPersonName           sql.NullString  `gorm:"column:AccountPersonName" json:"AccountPersonName"`
	AccountRemarks              sql.NullString  `gorm:"column:AccountRemarks" json:"AccountRemarks"`
	Approval                    int             `gorm:"column:Approval" json:"Approval"`
	ApprovalDate                time.Time       `gorm:"column:ApprovalDate" json:"ApprovalDate"`
	ApproverCode                sql.NullString  `gorm:"column:ApproverCode" json:"ApproverCode"`
	ApproverName                sql.NullString  `gorm:"column:ApproverName" json:"ApproverName"`
	ArrearDiscountAmount        sql.NullFloat64 `gorm:"column:ArrearDiscountAmount" json:"ArrearDiscountAmount"`
	ArrearPayAmount             sql.NullFloat64 `gorm:"column:ArrearPayAmount" json:"ArrearPayAmount"`
	ArrearPrePayAmount          sql.NullFloat64 `gorm:"column:ArrearPrePayAmount" json:"ArrearPrePayAmount"`
	BalanceAmount               sql.NullFloat64 `gorm:"column:BalanceAmount" json:"BalanceAmount"`
	Bank                        sql.NullString  `gorm:"column:Bank" json:"Bank"`
	BankAccount                 sql.NullString  `gorm:"column:BankAccount" json:"BankAccount"`
	BuyerCode                   sql.NullString  `gorm:"column:BuyerCode" json:"BuyerCode"`
	BuyerName                   sql.NullString  `gorm:"column:BuyerName" json:"BuyerName"`
	CarryCode                   sql.NullString  `gorm:"column:CarryCode" json:"CarryCode"`
	CarryName                   sql.NullString  `gorm:"column:CarryName" json:"CarryName"`
	CheckerCode                 sql.NullString  `gorm:"column:CheckerCode" json:"CheckerCode"`
	CheckerName                 sql.NullString  `gorm:"column:CheckerName" json:"CheckerName"`
	ComNo                       sql.NullString  `gorm:"column:ComNo" json:"ComNo"`
	Currency                    sql.NullString  `gorm:"column:Currency" json:"Currency"`
	CustomBillCode              sql.NullString  `gorm:"column:CustomBillCode" json:"CustomBillCode"`
	DeliveryCompany             sql.NullString  `gorm:"column:DeliveryCompany" json:"DeliveryCompany"`
	DeliveryDate                time.Time       `gorm:"column:DeliveryDate" json:"DeliveryDate"`
	DeliveryManCode             sql.NullString  `gorm:"column:DeliveryManCode" json:"DeliveryManCode"`
	DeliveryManName             sql.NullString  `gorm:"column:DeliveryManName" json:"DeliveryManName"`
	DeliveryRemarks             sql.NullString  `gorm:"column:DeliveryRemarks" json:"DeliveryRemarks"`
	DepartmentM                 sql.NullString  `gorm:"column:DepartmentM" json:"DepartmentM"`
	DiscountAmount              sql.NullFloat64 `gorm:"column:DiscountAmount" json:"DiscountAmount"`
	DisuseDate                  time.Time       `gorm:"column:DisuseDate" json:"DisuseDate"`
	DisuseReason                sql.NullString  `gorm:"column:DisuseReason" json:"DisuseReason"`
	DocDate                     time.Time       `gorm:"column:DocDate" json:"DocDate"`
	ExchangeRate                sql.NullFloat64 `gorm:"column:ExchangeRate" json:"ExchangeRate"`
	FastOutCode                 sql.NullString  `gorm:"column:FastOutCode" json:"FastOutCode"`
	FunctionCode                string          `gorm:"column:FunctionCode" json:"FunctionCode"`
	InType                      sql.NullString  `gorm:"column:InType" json:"InType"`
	InvoiceDate                 time.Time       `gorm:"column:InvoiceDate" json:"InvoiceDate"`
	InvoiceNo                   sql.NullString  `gorm:"column:InvoiceNo" json:"InvoiceNo"`
	InvoiceType                 sql.NullString  `gorm:"column:InvoiceType" json:"InvoiceType"`
	IsAccApp                    sql.NullInt64   `gorm:"column:IsAccApp" json:"IsAccApp"`
	IsArrear                    sql.NullInt64   `gorm:"column:IsArrear" json:"IsArrear"`
	IsFast                      sql.NullInt64   `gorm:"column:IsFast" json:"IsFast"`
	IsNeedDelivery              sql.NullInt64   `gorm:"column:IsNeedDelivery" json:"IsNeedDelivery"`
	IsPref                      sql.NullInt64   `gorm:"column:IsPref" json:"IsPref"`
	IsTax                       sql.NullInt64   `gorm:"column:IsTax" json:"IsTax"`
	ModifyDate                  time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP                    sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode            sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName            sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation           sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	MustDeliveryDate            time.Time       `gorm:"column:MustDeliveryDate" json:"MustDeliveryDate"`
	OperatorCode                sql.NullString  `gorm:"column:OperatorCode" json:"OperatorCode"`
	OperatorName                sql.NullString  `gorm:"column:OperatorName" json:"OperatorName"`
	OtherAmount                 sql.NullFloat64 `gorm:"column:OtherAmount" json:"OtherAmount"`
	PInCode                     string          `gorm:"column:PInCode" json:"PInCode"`
	PInPreReceiveAmount         sql.NullFloat64 `gorm:"column:PInPreReceiveAmount" json:"PInPreReceiveAmount"`
	PQualityJudgeCode           sql.NullString  `gorm:"column:PQualityJudgeCode" json:"PQualityJudgeCode"`
	PayAmount                   sql.NullFloat64 `gorm:"column:PayAmount" json:"PayAmount"`
	PayableAmount               sql.NullFloat64 `gorm:"column:PayableAmount" json:"PayableAmount"`
	PrePayAmount                sql.NullFloat64 `gorm:"column:PrePayAmount" json:"PrePayAmount"`
	PrefRate                    sql.NullFloat64 `gorm:"column:PrefRate" json:"PrefRate"`
	Remarks                     sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	ReturnReason                sql.NullString  `gorm:"column:ReturnReason" json:"ReturnReason"`
	ReturnType                  sql.NullString  `gorm:"column:ReturnType" json:"ReturnType"`
	SPrePayAmount               sql.NullFloat64 `gorm:"column:SPrePayAmount" json:"SPrePayAmount"`
	SupplierCode                sql.NullString  `gorm:"column:SupplierCode" json:"SupplierCode"`
	SupplierName                sql.NullString  `gorm:"column:SupplierName" json:"SupplierName"`
	TaxRate                     sql.NullFloat64 `gorm:"column:TaxRate" json:"TaxRate"`
	TotalCostAmount             sql.NullFloat64 `gorm:"column:TotalCostAmount" json:"TotalCostAmount"`
	TotalCount                  sql.NullInt64   `gorm:"column:TotalCount" json:"TotalCount"`
	TotalInCostAmount           sql.NullFloat64 `gorm:"column:TotalInCostAmount" json:"TotalInCostAmount"`
	TotalNoTaxCostAmount        sql.NullFloat64 `gorm:"column:TotalNoTaxCostAmount" json:"TotalNoTaxCostAmount"`
	TotalPresentAmount          sql.NullFloat64 `gorm:"column:TotalPresentAmount" json:"TotalPresentAmount"`
	TotalQty                    sql.NullFloat64 `gorm:"column:TotalQty" json:"TotalQty"`
	TotalRefPrice4Amount        sql.NullFloat64 `gorm:"column:TotalRefPrice4Amount" json:"TotalRefPrice4Amount"`
	TotalReturnCostProfitAmount sql.NullFloat64 `gorm:"column:TotalReturnCostProfitAmount" json:"TotalReturnCostProfitAmount"`
	TotalReturnPriceAmount      sql.NullFloat64 `gorm:"column:TotalReturnPriceAmount" json:"TotalReturnPriceAmount"`
	TotalType                   sql.NullInt64   `gorm:"column:TotalType" json:"TotalType"`
	TransferAmount              sql.NullFloat64 `gorm:"column:TransferAmount" json:"TransferAmount"`
	Transport                   sql.NullString  `gorm:"column:Transport" json:"Transport"`
	Cid                         sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ID                          int64           `gorm:"column:id;primary_key" json:"id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (p *PInm) TableName() string {
	return "p_inm"
}
