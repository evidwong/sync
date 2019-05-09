package models

import (
	"database/sql"
	"time"
)

type CMembersetm struct {
	AddTime           time.Time       `gorm:"column:add_time" json:"add_time"`
	Approvercode      sql.NullString  `gorm:"column:approvercode" json:"approvercode"`
	Approvername      sql.NullString  `gorm:"column:approvername" json:"approvername"`
	Bonusamount       sql.NullFloat64 `gorm:"column:bonusamount" json:"bonusamount"`
	Canconsumemoney   sql.NullFloat64 `gorm:"column:canconsumemoney" json:"canconsumemoney"`
	Cid               sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Cmodel            sql.NullString  `gorm:"column:cmodel" json:"cmodel"`
	Comno             sql.NullString  `gorm:"column:comno" json:"comno"`
	Comnocode         sql.NullString  `gorm:"column:comnocode" json:"comnocode"`
	ID                int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Integral          sql.NullInt64   `gorm:"column:integral" json:"integral"`
	Iscanintegral     sql.NullInt64   `gorm:"column:iscanintegral" json:"iscanintegral"`
	Isshow            sql.NullInt64   `gorm:"column:isshow" json:"isshow"`
	Isupload          sql.NullInt64   `gorm:"column:isupload" json:"isupload"`
	Isvalid           sql.NullInt64   `gorm:"column:isvalid" json:"isvalid"`
	Memberinfo        sql.NullString  `gorm:"column:memberinfo" json:"memberinfo"`
	Memberpic         sql.NullString  `gorm:"column:memberpic" json:"memberpic"`
	Memberprice       sql.NullFloat64 `gorm:"column:memberprice" json:"memberprice"`
	Membersetcode     sql.NullString  `gorm:"column:membersetcode" json:"membersetcode"`
	Membersetcomno    sql.NullString  `gorm:"column:membersetcomno" json:"membersetcomno"`
	Membersetmtype    sql.NullInt64   `gorm:"column:membersetmtype" json:"membersetmtype"`
	Membersetname     sql.NullString  `gorm:"column:membersetname" json:"membersetname"`
	Membersetype      sql.NullString  `gorm:"column:membersetype" json:"membersetype"`
	Membertype        sql.NullString  `gorm:"column:membertype" json:"membertype"`
	Modifydate        time.Time       `gorm:"column:modifydate" json:"modifydate"`
	Modifyip          sql.NullString  `gorm:"column:modifyip" json:"modifyip"`
	Modifypersoncode  sql.NullString  `gorm:"column:modifypersoncode" json:"modifypersoncode"`
	Modifypersonname  sql.NullString  `gorm:"column:modifypersonname" json:"modifypersonname"`
	Modifyworkstation sql.NullString  `gorm:"column:modifyworkstation" json:"modifyworkstation"`
	Presentjob        sql.NullInt64   `gorm:"column:presentjob" json:"presentjob"`
	Price             sql.NullFloat64 `gorm:"column:price" json:"price"`
	Profile           sql.NullString  `gorm:"column:profile" json:"profile"`
	Setlimit          sql.NullInt64   `gorm:"column:setlimit" json:"setlimit"`
	Setlimitunit      sql.NullString  `gorm:"column:setlimitunit" json:"setlimitunit"`
	StoreID           sql.NullString  `gorm:"column:store_id" json:"store_id"`
	StoreName         sql.NullString  `gorm:"column:store_name" json:"store_name"`
	Storedpackage     sql.NullString  `gorm:"column:storedpackage" json:"storedpackage"`
	Thumb             sql.NullString  `gorm:"column:thumb" json:"thumb"`
	UpdateTime        time.Time       `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (c *CMembersetm) TableName() string {
	return "c_membersetm"
}
