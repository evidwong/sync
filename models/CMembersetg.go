package models

import "database/sql"

type CMembersetg struct {
	Cid           sql.NullInt64   `gorm:"column:cid" json:"cid"`
	ExpireDate    sql.NullInt64   `gorm:"column:expire_date" json:"expire_date"`
	GoodsID       sql.NullInt64   `gorm:"column:goods_id" json:"goods_id"`
	ID            int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Identify      sql.NullString  `gorm:"column:identify" json:"identify"`
	Membersetcode sql.NullString  `gorm:"column:membersetcode" json:"membersetcode"`
	Number        int             `gorm:"column:number" json:"number"`
	Pid           sql.NullInt64   `gorm:"column:pid" json:"pid"`
	Price         sql.NullFloat64 `gorm:"column:price" json:"price"`
	Title         sql.NullString  `gorm:"column:title" json:"title"`
}

// TableName sets the insert table name for this struct type
func (c *CMembersetg) TableName() string {
	return "c_membersetg"
}
