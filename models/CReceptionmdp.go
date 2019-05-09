package models

import (
	"database/sql"
	"time"
)

type CReceptionmdp struct {
	Amount        sql.NullFloat64 `gorm:"column:amount" json:"amount"`
	BillingTime   time.Time       `gorm:"column:billingTime" json:"billingTime"`
	BusinessType  sql.NullString  `gorm:"column:businessType" json:"businessType"`
	CarModel      sql.NullString  `gorm:"column:carModel" json:"carModel"`
	CashierTime   time.Time       `gorm:"column:cashierTime" json:"cashierTime"`
	Cid           sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Contact       sql.NullString  `gorm:"column:contact" json:"contact"`
	Cost          sql.NullFloat64 `gorm:"column:cost" json:"cost"`
	CustomerCode  sql.NullString  `gorm:"column:customerCode" json:"customerCode"`
	CustomerName  sql.NullString  `gorm:"column:customerName" json:"customerName"`
	HandPhone     sql.NullString  `gorm:"column:handPhone" json:"handPhone"`
	ID            int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	OptUID        sql.NullInt64   `gorm:"column:opt_uid" json:"opt_uid"`
	OptUsername   sql.NullString  `gorm:"column:opt_username" json:"opt_username"`
	OrderCode     sql.NullString  `gorm:"column:orderCode" json:"orderCode"`
	OrgID         sql.NullInt64   `gorm:"column:org_id" json:"org_id"`
	OutTime       time.Time       `gorm:"column:outTime" json:"outTime"`
	ReceiptType   sql.NullString  `gorm:"column:receiptType" json:"receiptType"`
	RegisterNo    sql.NullString  `gorm:"column:registerNo" json:"registerNo"`
	Remark        sql.NullString  `gorm:"column:remark" json:"remark"`
	ServicePerson sql.NullString  `gorm:"column:servicePerson" json:"servicePerson"`
	Status        sql.NullString  `gorm:"column:status" json:"status"`
	Type          sql.NullInt64   `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (c *CReceptionmdp) TableName() string {
	return "c_receptionmdp"
}
