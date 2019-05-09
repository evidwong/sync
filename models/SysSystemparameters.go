package models

import "database/sql"

type SysSystemparameters struct {
	COMNo             string         `gorm:"column:COMNo" json:"COMNo"`
	COMNoName         string         `gorm:"column:COMNoName" json:"COMNoName"`
	DefaultValue      sql.NullString `gorm:"column:DefaultValue" json:"DefaultValue"`
	GroupID           sql.NullInt64  `gorm:"column:GroupId" json:"GroupId"`
	GroupName         sql.NullString `gorm:"column:GroupName" json:"GroupName"`
	IsCanEmpty        sql.NullInt64  `gorm:"column:IsCanEmpty" json:"IsCanEmpty"`
	IsComno           int            `gorm:"column:IsComno" json:"IsComno"`
	IsMoreChoise      sql.NullInt64  `gorm:"column:IsMoreChoise" json:"IsMoreChoise"`
	IsRange           sql.NullInt64  `gorm:"column:IsRange" json:"IsRange"`
	IsRangeClass      sql.NullInt64  `gorm:"column:IsRangeClass" json:"IsRangeClass"`
	KeyItem           sql.NullString `gorm:"column:KeyItem" json:"KeyItem"`
	ModifyDate        sql.NullInt64  `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP          sql.NullString `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode  sql.NullString `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName  sql.NullString `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation sql.NullString `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	Options           sql.NullString `gorm:"column:Options" json:"Options"`
	ParamName         string         `gorm:"column:ParamName" json:"ParamName"`
	ParamType         sql.NullString `gorm:"column:ParamType" json:"ParamType"`
	ParamValue        sql.NullString `gorm:"column:ParamValue" json:"ParamValue"`
	Range             sql.NullString `gorm:"column:Range" json:"Range"`
	Remarks           sql.NullString `gorm:"column:Remarks" json:"Remarks"`
	ShortDesc         sql.NullString `gorm:"column:ShortDesc" json:"ShortDesc"`
	StrLen            sql.NullInt64  `gorm:"column:StrLen" json:"StrLen"`
	Version           sql.NullString `gorm:"column:Version" json:"Version"`
	Cid               int            `gorm:"column:cid" json:"cid"`
	ID                int            `gorm:"column:id;primary_key" json:"id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (s *SysSystemparameters) TableName() string {
	return "sys_systemparameters"
}
