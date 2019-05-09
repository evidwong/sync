package models

import "database/sql"

type Package struct {
	Comno              string         `gorm:"column:comno" json:"comno"`
	Customercode       sql.NullString `gorm:"column:customercode" json:"customercode"`
	Customername       sql.NullString `gorm:"column:customername" json:"customername"`
	ID                 int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isover             int            `gorm:"column:isover" json:"isover"`
	JobName            sql.NullString `gorm:"column:job_name" json:"job_name"`
	Membersetsalescode sql.NullString `gorm:"column:membersetsalescode" json:"membersetsalescode"`
	PackageCode        string         `gorm:"column:package_code" json:"package_code"`
	PackageName        string         `gorm:"column:package_name" json:"package_name"`
	Salesdate          int            `gorm:"column:salesdate" json:"salesdate"`
	ServiceNumber      int            `gorm:"column:service_number" json:"service_number"`
	Status             sql.NullString `gorm:"column:status" json:"status"`
	Surplus            int            `gorm:"column:surplus" json:"surplus"`
	UID                int            `gorm:"column:uid" json:"uid"`
	Used               int            `gorm:"column:used" json:"used"`
	Validity           int            `gorm:"column:validity" json:"validity"`
	Vehicle            sql.NullString `gorm:"column:vehicle" json:"vehicle"`
	Vehiclecode        sql.NullString `gorm:"column:vehiclecode" json:"vehiclecode"`
}

// TableName sets the insert table name for this struct type
func (p *Package) TableName() string {
	return "package"
}
