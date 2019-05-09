package models

import (
	"database/sql"
	"time"
)

type JBespeak struct {
	AddTime           time.Time      `gorm:"column:add_time" json:"add_time"`
	Advancetime       sql.NullInt64  `gorm:"column:advancetime" json:"advancetime"`
	Cid               sql.NullInt64  `gorm:"column:cid" json:"cid"`
	ID                int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Intervaltime      sql.NullInt64  `gorm:"column:intervaltime" json:"intervaltime"`
	Pid               sql.NullInt64  `gorm:"column:pid" json:"pid"`
	Restamcycle       sql.NullInt64  `gorm:"column:restamcycle" json:"restamcycle"`
	Restamendtime     time.Time      `gorm:"column:restamendtime" json:"restamendtime"`
	Restamstarttime   time.Time      `gorm:"column:restamstarttime" json:"restamstarttime"`
	Restamtotal       sql.NullInt64  `gorm:"column:restamtotal" json:"restamtotal"`
	Restpmcycle       sql.NullInt64  `gorm:"column:restpmcycle" json:"restpmcycle"`
	Restpmendtime     time.Time      `gorm:"column:restpmendtime" json:"restpmendtime"`
	Restpmstarttime   time.Time      `gorm:"column:restpmstarttime" json:"restpmstarttime"`
	Restpmtotal       sql.NullInt64  `gorm:"column:restpmtotal" json:"restpmtotal"`
	Storeid           sql.NullString `gorm:"column:storeid" json:"storeid"`
	UpdateTime        time.Time      `gorm:"column:update_time" json:"update_time"`
	Validityendtime   time.Time      `gorm:"column:validityendtime" json:"validityendtime"`
	Validitystarttime time.Time      `gorm:"column:validitystarttime" json:"validitystarttime"`
	Workeramcycle     sql.NullInt64  `gorm:"column:workeramcycle" json:"workeramcycle"`
	Workeramendtime   time.Time      `gorm:"column:workeramendtime" json:"workeramendtime"`
	Workeramstarttime time.Time      `gorm:"column:workeramstarttime" json:"workeramstarttime"`
	Workeramtotal     sql.NullInt64  `gorm:"column:workeramtotal" json:"workeramtotal"`
	Workerpmcycle     sql.NullInt64  `gorm:"column:workerpmcycle" json:"workerpmcycle"`
	Workerpmendtime   time.Time      `gorm:"column:workerpmendtime;primary_key" json:"workerpmendtime;primary_key"`
	Workerpmstarttime time.Time      `gorm:"column:workerpmstarttime" json:"workerpmstarttime"`
	Workerpmtotal     sql.NullInt64  `gorm:"column:workerpmtotal" json:"workerpmtotal"`
}

// TableName sets the insert table name for this struct type
func (j *JBespeak) TableName() string {
	return "j_bespeak"
}
