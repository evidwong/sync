package models

import "database/sql"

type InsuranceClerk struct {
	AddTime      int            `gorm:"column:add_time" json:"add_time"`
	Adminid      int            `gorm:"column:adminid" json:"adminid"`
	Aid          int            `gorm:"column:aid" json:"aid"`
	Auid         int            `gorm:"column:auid" json:"auid"`
	Cid          int            `gorm:"column:cid" json:"cid"`
	Coid         int            `gorm:"column:coid" json:"coid"`
	Email        sql.NullString `gorm:"column:email" json:"email"`
	ID           int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel        int            `gorm:"column:isdel" json:"isdel"`
	Jobid        sql.NullString `gorm:"column:jobid" json:"jobid"`
	Name         string         `gorm:"column:name" json:"name"`
	Phone        string         `gorm:"column:phone" json:"phone"`
	Rebate       float32        `gorm:"column:rebate" json:"rebate"`
	RebateAmount float32        `gorm:"column:rebate_amount" json:"rebate_amount"`
	Remark       sql.NullString `gorm:"column:remark" json:"remark"`
	RoleID       int            `gorm:"column:role_id" json:"role_id"`
	Status       int            `gorm:"column:status" json:"status"`
	UpdateTime   int            `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (i *InsuranceClerk) TableName() string {
	return "insurance_clerk"
}
