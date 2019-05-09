package models

import "database/sql"

type PMonthschedule struct {
	Attribute1    sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2    sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3    sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4    sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5    sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6    sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7    sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8    sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	COMNo         sql.NullString  `gorm:"column:COMNo" json:"COMNo"`
	EnglishName   sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode         sql.NullString  `gorm:"column:FCode" json:"FCode"`
	GroupNo       sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	ItemName      sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	PScheduleCode string          `gorm:"column:PScheduleCode;primary_key" json:"PScheduleCode;primary_key"`
	ScheQty       sql.NullFloat64 `gorm:"column:ScheQty" json:"ScheQty"`
	Unit          sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode   sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	YearMonth     sql.NullString  `gorm:"column:YearMonth" json:"YearMonth"`
	Cid           sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode      sql.NullString  `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (p *PMonthschedule) TableName() string {
	return "p_monthschedule"
}
