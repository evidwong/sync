package models

import (
	"database/sql"
	"time"
)

type CMembertype struct {
	AddTime           time.Time       `gorm:"column:add_time" json:"add_time"`
	Bonusamount       sql.NullFloat64 `gorm:"column:bonusamount" json:"bonusamount"`
	Cid               sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Classorder        sql.NullInt64   `gorm:"column:classorder" json:"classorder"`
	Comno             sql.NullString  `gorm:"column:comno" json:"comno"`
	Comnocode         sql.NullString  `gorm:"column:comnocode" json:"comnocode"`
	ID                int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Integralmin       sql.NullFloat64 `gorm:"column:integralmin" json:"integralmin"`
	Integralrate      sql.NullFloat64 `gorm:"column:integralrate" json:"integralrate"`
	Integralrule      sql.NullFloat64 `gorm:"column:integralrule" json:"integralrule"`
	Isshow            sql.NullInt64   `gorm:"column:isshow" json:"isshow"`
	Isupload          sql.NullInt64   `gorm:"column:isupload" json:"isupload"`
	Jobdiscount       sql.NullFloat64 `gorm:"column:jobdiscount" json:"jobdiscount"`
	Jobdiscounttype   sql.NullString  `gorm:"column:jobdiscounttype" json:"jobdiscounttype"`
	Limitmonth        sql.NullInt64   `gorm:"column:limitmonth" json:"limitmonth"`
	Memberamount      sql.NullFloat64 `gorm:"column:memberamount" json:"memberamount"`
	Memberinfo        sql.NullString  `gorm:"column:memberinfo" json:"memberinfo"`
	Memberpic         sql.NullString  `gorm:"column:memberpic" json:"memberpic"`
	Membertype        string          `gorm:"column:membertype" json:"membertype"`
	Modifydate        time.Time       `gorm:"column:modifydate" json:"modifydate"`
	Modifyip          sql.NullString  `gorm:"column:modifyip" json:"modifyip"`
	Modifypersoncode  sql.NullString  `gorm:"column:modifypersoncode" json:"modifypersoncode"`
	Modifypersonname  sql.NullString  `gorm:"column:modifypersonname" json:"modifypersonname"`
	Modifyworkstation sql.NullString  `gorm:"column:modifyworkstation" json:"modifyworkstation"`
	Partsdiscounttype sql.NullString  `gorm:"column:partsdiscounttype" json:"partsdiscounttype"`
	Profile           sql.NullString  `gorm:"column:profile" json:"profile"`
	Remarks           sql.NullString  `gorm:"column:remarks" json:"remarks"`
	Setdiscounttype   sql.NullString  `gorm:"column:setdiscounttype" json:"setdiscounttype"`
	StoreID           sql.NullString  `gorm:"column:store_id" json:"store_id"`
	StoreName         sql.NullString  `gorm:"column:store_name" json:"store_name"`
	Storedpackage     sql.NullString  `gorm:"column:storedpackage" json:"storedpackage"`
	Thumb             sql.NullString  `gorm:"column:thumb" json:"thumb"`
	UpdateTime        time.Time       `gorm:"column:update_time" json:"update_time"`
	Upgradeintegral   sql.NullFloat64 `gorm:"column:upgradeintegral" json:"upgradeintegral"`
}

// TableName sets the insert table name for this struct type
func (c *CMembertype) TableName() string {
	return "c_membertype"
}
