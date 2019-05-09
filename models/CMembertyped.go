package models

import "database/sql"

type CMembertyped struct {
	Cid         sql.NullInt64   `gorm:"column:cid" json:"cid"`
	DepositRate sql.NullFloat64 `gorm:"column:deposit_rate" json:"deposit_rate"`
	ID          int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Integrate   sql.NullFloat64 `gorm:"column:integrate" json:"integrate"`
	Membercode  sql.NullString  `gorm:"column:membercode" json:"membercode"`
	Pid         sql.NullInt64   `gorm:"column:pid" json:"pid"`
	StoreID     sql.NullInt64   `gorm:"column:store_id" json:"store_id"`
	Worktype    sql.NullString  `gorm:"column:worktype" json:"worktype"`
}

// TableName sets the insert table name for this struct type
func (c *CMembertyped) TableName() string {
	return "c_membertyped"
}
