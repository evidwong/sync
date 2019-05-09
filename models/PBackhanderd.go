package models

import (
	"database/sql"
	"time"
)

type PBackhanderd struct {
	ActualAmount    sql.NullFloat64 `gorm:"column:ActualAmount" json:"ActualAmount"`
	Attribute1      sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2      sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3      sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4      sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5      sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6      sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7      sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8      sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	BHAmount        sql.NullFloat64 `gorm:"column:BHAmount" json:"BHAmount"`
	BHDate          time.Time       `gorm:"column:BHDate" json:"BHDate"`
	BHPayee         sql.NullString  `gorm:"column:BHPayee" json:"BHPayee"`
	BHPercent       sql.NullFloat64 `gorm:"column:BHPercent" json:"BHPercent"`
	COMNo           sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	CustomerCode    sql.NullString  `gorm:"column:CustomerCode" json:"CustomerCode"`
	CustomerName    sql.NullString  `gorm:"column:CustomerName" json:"CustomerName"`
	EnglishName     sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode           sql.NullString  `gorm:"column:FCode" json:"FCode"`
	FunctionCode    sql.NullString  `gorm:"column:FunctionCode" json:"FunctionCode"`
	GroupNo         sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	IndexID         int             `gorm:"column:IndexID;primary_key" json:"IndexID;primary_key"`
	ItemIndexID     sql.NullInt64   `gorm:"column:ItemIndexID" json:"ItemIndexID"`
	ItemName        sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	PBackHanderCode string          `gorm:"column:PBackHanderCode;primary_key" json:"PBackHanderCode;primary_key"`
	POutCode        sql.NullString  `gorm:"column:POutCode" json:"POutCode"`
	ReturnAmount    sql.NullFloat64 `gorm:"column:ReturnAmount" json:"ReturnAmount"`
	SalesAmount     sql.NullFloat64 `gorm:"column:SalesAmount" json:"SalesAmount"`
	SalesLeader     sql.NullString  `gorm:"column:SalesLeader" json:"SalesLeader"`
	SalesmanCode    sql.NullString  `gorm:"column:SalesmanCode" json:"SalesmanCode"`
	SalesmanName    sql.NullString  `gorm:"column:SalesmanName" json:"SalesmanName"`
	Unit            sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode     sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	Cid             sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode        sql.NullString  `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (p *PBackhanderd) TableName() string {
	return "p_backhanderd"
}
