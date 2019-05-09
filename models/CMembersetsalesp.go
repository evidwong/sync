package models

import (
	"database/sql"
	"time"
)

type CMembersetsalesp struct {
	Attribute1         sql.NullString  `gorm:"column:Attribute1" json:"Attribute1"`
	Attribute2         sql.NullString  `gorm:"column:Attribute2" json:"Attribute2"`
	Attribute3         sql.NullString  `gorm:"column:Attribute3" json:"Attribute3"`
	Attribute4         sql.NullString  `gorm:"column:Attribute4" json:"Attribute4"`
	Attribute5         sql.NullString  `gorm:"column:Attribute5" json:"Attribute5"`
	Attribute6         sql.NullString  `gorm:"column:Attribute6" json:"Attribute6"`
	Attribute7         sql.NullString  `gorm:"column:Attribute7" json:"Attribute7"`
	Attribute8         sql.NullString  `gorm:"column:Attribute8" json:"Attribute8"`
	BalanceOutQty      float64         `gorm:"column:BalanceOutQty" json:"BalanceOutQty"`
	CanOutQty          float64         `gorm:"column:CanOutQty" json:"CanOutQty"`
	ChargeAmount       sql.NullFloat64 `gorm:"column:ChargeAmount" json:"ChargeAmount"`
	ComNo              sql.NullString  `gorm:"column:ComNo" json:"ComNo"`
	EnglishName        sql.NullString  `gorm:"column:EnglishName" json:"EnglishName"`
	FCode              sql.NullString  `gorm:"column:FCode" json:"FCode"`
	GroupNo            sql.NullString  `gorm:"column:GroupNo" json:"GroupNo"`
	IndexID            sql.NullInt64   `gorm:"column:IndexID" json:"IndexID"`
	IsAdditionalItem   sql.NullInt64   `gorm:"column:IsAdditionalItem" json:"IsAdditionalItem"`
	IsPresentItem      sql.NullInt64   `gorm:"column:IsPresentItem" json:"IsPresentItem"`
	ItemCode           sql.NullString  `gorm:"column:ItemCode" json:"ItemCode"`
	ItemName           sql.NullString  `gorm:"column:ItemName" json:"ItemName"`
	LastOutDate        time.Time       `gorm:"column:LastOutDate" json:"LastOutDate"`
	LimitDate          time.Time       `gorm:"column:LimitDate" json:"LimitDate"`
	MemberSetSalesCode sql.NullString  `gorm:"column:MemberSetSalesCode" json:"MemberSetSalesCode"`
	OrderID            sql.NullInt64   `gorm:"column:OrderID" json:"OrderID"`
	Remarks            sql.NullString  `gorm:"column:Remarks" json:"Remarks"`
	Unit               sql.NullString  `gorm:"column:Unit" json:"Unit"`
	UseOutQty          float64         `gorm:"column:UseOutQty" json:"UseOutQty"`
	UseRegisterNo      sql.NullString  `gorm:"column:UseRegisterNo" json:"UseRegisterNo"`
	UseVehicleCode     sql.NullString  `gorm:"column:UseVehicleCode" json:"UseVehicleCode"`
	UserDefCode        sql.NullString  `gorm:"column:UserDefCode" json:"UserDefCode"`
	Attributecode      string          `gorm:"column:attributecode" json:"attributecode"`
	Cid                int             `gorm:"column:cid" json:"cid"`
	ConsumePrice       sql.NullFloat64 `gorm:"column:consume_price" json:"consume_price"`
	Costprice          sql.NullFloat64 `gorm:"column:costprice" json:"costprice"`
	ID                 int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isintegral         sql.NullInt64   `gorm:"column:isintegral" json:"isintegral"`
	StandardPrice      sql.NullFloat64 `gorm:"column:standard_price" json:"standard_price"`
}

// TableName sets the insert table name for this struct type
func (c *CMembersetsalesp) TableName() string {
	return "c_membersetsalesp"
}
