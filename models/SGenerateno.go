package models

import (
	"database/sql"
	"time"
)

type SGenerateno struct {
	Function          sql.NullString `gorm:"column:Function" json:"Function"`
	FunctionCode      sql.NullString `gorm:"column:FunctionCode" json:"FunctionCode"`
	GenerateType      sql.NullString `gorm:"column:GenerateType" json:"GenerateType"`
	IsPrefixUseCOMNo  sql.NullInt64  `gorm:"column:IsPrefixUseCOMNo" json:"IsPrefixUseCOMNo"`
	IsUse             sql.NullInt64  `gorm:"column:IsUse" json:"IsUse"`
	IsValid           sql.NullInt64  `gorm:"column:IsValid" json:"IsValid"`
	KeyFieldName      sql.NullString `gorm:"column:KeyFieldName" json:"KeyFieldName"`
	LastDate          time.Time      `gorm:"column:LastDate" json:"LastDate"`
	LastNo            int            `gorm:"column:LastNo" json:"LastNo"`
	Length            sql.NullInt64  `gorm:"column:Length" json:"Length"`
	ModifyDate        time.Time      `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP          sql.NullString `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode  sql.NullString `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName  sql.NullString `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation sql.NullString `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	Prefix            sql.NullString `gorm:"column:Prefix" json:"Prefix"`
	TableName         sql.NullString `gorm:"column:TableName" json:"TableName"`
	Cid               sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Comno             sql.NullString `gorm:"column:comno" json:"comno"`
	ID                int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	StoreID           sql.NullInt64  `gorm:"column:store_id" json:"store_id"`
	StoreName         sql.NullString `gorm:"column:store_name" json:"store_name"`
}