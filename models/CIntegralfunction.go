package models

import "database/sql"

type CIntegralfunction struct {
	Integral          sql.NullFloat64 `gorm:"column:Integral" json:"Integral"`
	IntegralFunction  string          `gorm:"column:IntegralFunction" json:"IntegralFunction"`
	ModifyDate        sql.NullInt64   `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP          sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode  sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName  sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	Cid               sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Comno             sql.NullString  `gorm:"column:comno" json:"comno"`
	Comnocode         sql.NullString  `gorm:"column:comnocode" json:"comnocode"`
	ID                int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	StoreID           sql.NullString  `gorm:"column:store_id" json:"store_id"`
	StoreName         sql.NullString  `gorm:"column:store_name" json:"store_name"`
}

// TableName sets the insert table name for this struct type
func (c *CIntegralfunction) TableName() string {
	return "c_integralfunction"
}
