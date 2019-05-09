package models

import "database/sql"

type Violation struct {
	Cid               sql.NullInt64 `gorm:"column:cid" json:"cid"`
	Dealfee           int           `gorm:"column:dealfee" json:"dealfee"`
	Dealflag          int           `gorm:"column:dealflag" json:"dealflag"`
	Dealmsg           string        `gorm:"column:dealmsg" json:"dealmsg"`
	Dealtime          string        `gorm:"column:dealtime" json:"dealtime"`
	Exist             int           `gorm:"column:exist" json:"exist"`
	Finefee           int           `gorm:"column:finefee" json:"finefee"`
	ID                int           `gorm:"column:id;primary_key" json:"id;primary_key"`
	Latefee           int           `gorm:"column:latefee" json:"latefee"`
	Livebill          int           `gorm:"column:livebill" json:"livebill"`
	Otherway          string        `gorm:"column:otherway" json:"otherway"`
	Policeaddress     string        `gorm:"column:policeaddress" json:"policeaddress"`
	Policecontact     string        `gorm:"column:policecontact" json:"policecontact"`
	Province          string        `gorm:"column:province" json:"province"`
	Referfee          int           `gorm:"column:referfee" json:"referfee"`
	Score             int           `gorm:"column:score" json:"score"`
	Status            int           `gorm:"column:status" json:"status"`
	Totalfee          int           `gorm:"column:totalfee" json:"totalfee"`
	Vehicle           string        `gorm:"column:vehicle" json:"vehicle"`
	Viloationdetail   string        `gorm:"column:viloationdetail" json:"viloationdetail"`
	Viloationlocation string        `gorm:"column:viloationlocation" json:"viloationlocation"`
	Violationid       int           `gorm:"column:violationid" json:"violationid"`
	Violationtime     string        `gorm:"column:violationtime" json:"violationtime"`
	Writetime         int           `gorm:"column:writetime" json:"writetime"`
	Wsh               string        `gorm:"column:wsh" json:"wsh"`
	Wzdm              string        `gorm:"column:wzdm" json:"wzdm"`
}

// TableName sets the insert table name for this struct type
func (v *Violation) TableName() string {
	return "violation"
}
