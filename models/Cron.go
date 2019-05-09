package models

import (
	"database/sql"
	"time"
)

type Cron struct {
	Cid         int            `gorm:"column:cid" json:"cid"`
	EndAt       time.Time      `gorm:"column:end_at" json:"end_at"`
	EndTime     time.Time      `gorm:"column:end_time" json:"end_time"`
	ExpireStep  sql.NullString `gorm:"column:expire_step" json:"expire_step"`
	ExpireType  sql.NullString `gorm:"column:expire_type" json:"expire_type"`
	ID          int64          `gorm:"column:id;primary_key" json:"id;primary_key"`
	LastRunTime time.Time      `gorm:"column:last_run_time" json:"last_run_time"`
	PushType    sql.NullString `gorm:"column:push_type" json:"push_type"`
	StartAt     time.Time      `gorm:"column:start_at" json:"start_at"`
	StartTime   time.Time      `gorm:"column:start_time" json:"start_time"`
	Status      sql.NullInt64  `gorm:"column:status" json:"status"`
	Step        sql.NullInt64  `gorm:"column:step" json:"step"`
	StepType    sql.NullString `gorm:"column:step_type" json:"step_type"`
	Times       sql.NullInt64  `gorm:"column:times" json:"times"`
	Title       sql.NullString `gorm:"column:title" json:"title"`
	Type        sql.NullString `gorm:"column:type" json:"type"`
	Workerid    int            `gorm:"column:workerid" json:"workerid"`
}

// TableName sets the insert table name for this struct type
func (c *Cron) TableName() string {
	return "cron"
}
