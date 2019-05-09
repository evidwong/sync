package models

import (
	"database/sql"
	"time"
)

type VGItemlib struct {
	ArriveDays        sql.NullInt64   `gorm:"column:ArriveDays" json:"ArriveDays"`
	Attribute1        sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2        sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3        sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4        sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5        sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6        sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7        sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8        sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	BOMCode           sql.NullString  `gorm:"column:BOMCode" json:"BOMCode"`
	Box               sql.NullString  `gorm:"column:Box" json:"Box"`
	DeliverPeriod     sql.NullInt64   `gorm:"column:DeliverPeriod" json:"DeliverPeriod"`
	EnglishName       sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode             sql.NullString  `gorm:"column:FCode" json:"FCode"`
	Grade             sql.NullString  `gorm:"column:Grade" json:"Grade"`
	GroupNo           sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	Height            sql.NullFloat64 `gorm:"column:Height" json:"Height"`
	IsCanIntegral     sql.NullInt64   `gorm:"column:IsCanIntegral" json:"IsCanIntegral"`
	IsCanPresent      sql.NullInt64   `gorm:"column:IsCanPresent" json:"IsCanPresent"`
	IsCancel          sql.NullInt64   `gorm:"column:IsCancel" json:"IsCancel"`
	ItemLib           string          `gorm:"column:ItemLib" json:"ItemLib"`
	ItemMemo1         sql.NullString  `gorm:"column:ItemMemo1" json:"ItemMemo1"`
	ItemMemo2         sql.NullString  `gorm:"column:ItemMemo2" json:"ItemMemo2"`
	ItemMemo3         sql.NullString  `gorm:"column:ItemMemo3" json:"ItemMemo3"`
	ItemName          sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	Length            sql.NullFloat64 `gorm:"column:Length" json:"Length"`
	MinOrderQty       sql.NullFloat64 `gorm:"column:MinOrderQty" json:"MinOrderQty"`
	ModifyDate        time.Time       `gorm:"column:ModifyDate" json:"ModifyDate"`
	ModifyIP          sql.NullString  `gorm:"column:ModifyIP" json:"ModifyIP"`
	ModifyPersonCode  sql.NullString  `gorm:"column:ModifyPersonCode" json:"ModifyPersonCode"`
	ModifyPersonName  sql.NullString  `gorm:"column:ModifyPersonName" json:"ModifyPersonName"`
	ModifyWorkStation sql.NullString  `gorm:"column:ModifyWorkStation" json:"ModifyWorkStation"`
	OrderPeriod       sql.NullInt64   `gorm:"column:OrderPeriod" json:"OrderPeriod"`
	OutOrderQty       sql.NullFloat64 `gorm:"column:OutOrderQty" json:"OutOrderQty"`
	PYCode            sql.NullString  `gorm:"column:PYCode" json:"PYCode"`
	PackQty           sql.NullFloat64 `gorm:"column:PackQty" json:"PackQty"`
	PackUnit          sql.NullString  `gorm:"column:PackUnit" json:"PackUnit"`
	PartsCOMNo        sql.NullString  `gorm:"column:PartsCOMNo" json:"PartsCOMNo"`
	PartsIntegral     sql.NullFloat64 `gorm:"column:PartsIntegral" json:"PartsIntegral"`
	RelationCode      sql.NullString  `gorm:"column:RelationCode" json:"RelationCode"`
	RelationNo        sql.NullInt64   `gorm:"column:RelationNo" json:"RelationNo"`
	SingleVol         sql.NullFloat64 `gorm:"column:SingleVol" json:"SingleVol"`
	SuitQty           sql.NullFloat64 `gorm:"column:SuitQty" json:"SuitQty"`
	TmpField1         sql.NullString  `gorm:"column:TmpField1" json:"TmpField1"`
	TmpField2         time.Time       `gorm:"column:TmpField2" json:"TmpField2"`
	Type1             sql.NullString  `gorm:"column:Type1" json:"Type1"`
	Type2             sql.NullString  `gorm:"column:Type2" json:"Type2"`
	Type3             sql.NullString  `gorm:"column:Type3" json:"Type3"`
	Type4             sql.NullString  `gorm:"column:Type4" json:"Type4"`
	Type5             sql.NullString  `gorm:"column:Type5" json:"Type5"`
	Unit              sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UserDefCode       sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	Vol               sql.NullFloat64 `gorm:"column:Vol" json:"Vol"`
	WBCode            sql.NullString  `gorm:"column:WBCode" json:"WBCode"`
	Weight            sql.NullFloat64 `gorm:"column:Weight" json:"Weight"`
	Width             sql.NullFloat64 `gorm:"column:Width" json:"Width"`
	Cid               sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Itemcode          string          `gorm:"column:itemcode" json:"itemcode"`
}

// TableName sets the insert table name for this struct type
func (v *VGItemlib) TableName() string {
	return "v_g_itemlib"
}
