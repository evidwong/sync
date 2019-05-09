package models

import "database/sql"

type InsuranceCompany struct {
	AddTime    int            `gorm:"column:add_time" json:"add_time"`
	Aid        int            `gorm:"column:aid" json:"aid"`
	Auid       int            `gorm:"column:auid" json:"auid"`
	Cid        int            `gorm:"column:cid" json:"cid"`
	ID         int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel      int            `gorm:"column:isdel" json:"isdel"`
	Name       string         `gorm:"column:name" json:"name"`
	Pic        sql.NullString `gorm:"column:pic" json:"pic"`
	Rebate     float32        `gorm:"column:rebate" json:"rebate"`
	RebateType int            `gorm:"column:rebate_type" json:"rebate_type"`
	Recommend  int            `gorm:"column:recommend" json:"recommend"`
	Remark     sql.NullString `gorm:"column:remark" json:"remark"`
	ShortName  sql.NullString `gorm:"column:short_name" json:"short_name"`
	Status     int            `gorm:"column:status" json:"status"`
	Thumb      sql.NullString `gorm:"column:thumb" json:"thumb"`
	UpdateTime int            `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (i *InsuranceCompany) TableName() string {
	return "insurance_company"
}
