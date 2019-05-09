package models

import (
	"database/sql"
	"time"
)

type Historyreport struct {
	CdAmount          sql.NullFloat64 `gorm:"column:cd_Amount" json:"cd_Amount"`
	CdClass           sql.NullString  `gorm:"column:cd_Class" json:"cd_Class"`
	CdCostPrice       sql.NullFloat64 `gorm:"column:cd_CostPrice" json:"cd_CostPrice"`
	CdJobName         sql.NullString  `gorm:"column:cd_JobName" json:"cd_JobName"`
	CdJobType         sql.NullString  `gorm:"column:cd_JobType" json:"cd_JobType"`
	CdMainWorker      sql.NullString  `gorm:"column:cd_MainWorker" json:"cd_MainWorker"`
	CdRecTypeD        sql.NullString  `gorm:"column:cd_RecTypeD" json:"cd_RecTypeD"`
	CdSalesLeader     sql.NullString  `gorm:"column:cd_SalesLeader" json:"cd_SalesLeader"`
	CdTimeAmount      sql.NullFloat64 `gorm:"column:cd_TimeAmount" json:"cd_TimeAmount"`
	CdWorkTime        sql.NullFloat64 `gorm:"column:cd_WorkTime" json:"cd_WorkTime"`
	CdID              sql.NullInt64   `gorm:"column:cd_id" json:"cd_id"`
	CmBrand           sql.NullString  `gorm:"column:cm_Brand" json:"cm_Brand"`
	CmCReceptionCode  sql.NullString  `gorm:"column:cm_CReceptionCode" json:"cm_CReceptionCode"`
	CmInvoiceType     sql.NullString  `gorm:"column:cm_InvoiceType" json:"cm_InvoiceType"`
	CmIsRecMoney      int             `gorm:"column:cm_IsRecMoney" json:"cm_IsRecMoney"`
	CmRecMoneyDate    time.Time       `gorm:"column:cm_RecMoneyDate" json:"cm_RecMoneyDate"`
	CmRecPersonName   sql.NullString  `gorm:"column:cm_RecPersonName" json:"cm_RecPersonName"`
	CmReceptionSource sql.NullString  `gorm:"column:cm_ReceptionSource" json:"cm_ReceptionSource"`
	CmRegisterNo      sql.NullString  `gorm:"column:cm_RegisterNo" json:"cm_RegisterNo"`
	CmVehicleCode     sql.NullString  `gorm:"column:cm_VehicleCode" json:"cm_VehicleCode"`
	CmVehicleLevel    sql.NullString  `gorm:"column:cm_VehicleLevel" json:"cm_VehicleLevel"`
	CmCid             sql.NullInt64   `gorm:"column:cm_cid" json:"cm_cid"`
	CmComno           sql.NullString  `gorm:"column:cm_comno" json:"cm_comno"`
	CmID              sql.NullInt64   `gorm:"column:cm_id" json:"cm_id"`
	ID                int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	PdAmount          sql.NullFloat64 `gorm:"column:pd_Amount" json:"pd_Amount"`
	PdAttribute5      sql.NullString  `gorm:"column:pd_Attribute5" json:"pd_Attribute5"`
	PdAttritube3      sql.NullString  `gorm:"column:pd_Attritube3" json:"pd_Attritube3"`
	PdCostPrice       sql.NullFloat64 `gorm:"column:pd_CostPrice" json:"pd_CostPrice"`
	PdItemCode        sql.NullString  `gorm:"column:pd_ItemCode" json:"pd_ItemCode"`
	PdItemName        sql.NullString  `gorm:"column:pd_ItemName" json:"pd_ItemName"`
	PdPrice           sql.NullFloat64 `gorm:"column:pd_Price" json:"pd_Price"`
	PdQty             sql.NullFloat64 `gorm:"column:pd_Qty" json:"pd_Qty"`
	PdRecTypeD        sql.NullString  `gorm:"column:pd_RecTypeD" json:"pd_RecTypeD"`
	PdSalesLeader     sql.NullString  `gorm:"column:pd_SalesLeader" json:"pd_SalesLeader"`
	PdClass           sql.NullString  `gorm:"column:pd_class" json:"pd_class"`
	PdID              sql.NullInt64   `gorm:"column:pd_id" json:"pd_id"`
	PmPOutCode        sql.NullString  `gorm:"column:pm_POutCode" json:"pm_POutCode"`
	PmID              sql.NullInt64   `gorm:"column:pm_id" json:"pm_id"`
}

// TableName sets the insert table name for this struct type
func (h *Historyreport) TableName() string {
	return "historyreport"
}
