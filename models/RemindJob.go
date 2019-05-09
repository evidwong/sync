package models

import "time"

type RemindJob struct {
	Action       string    `gorm:"column:action" json:"action"`
	ActionAt     time.Time `gorm:"column:action_at" json:"action_at"`
	Cid          int       `gorm:"column:cid" json:"cid"`
	Comno        string    `gorm:"column:comno" json:"comno"`
	CreateAt     time.Time `gorm:"column:create_at" json:"create_at"`
	FailContent  string    `gorm:"column:fail_content" json:"fail_content"`
	FlagNum      int       `gorm:"column:flag_num" json:"flag_num"`
	FlagTime     string    `gorm:"column:flag_time" json:"flag_time"`
	FromID       int       `gorm:"column:from_id" json:"from_id"`
	FunctionCode string    `gorm:"column:function_code" json:"function_code"`
	ID           int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	Job          string    `gorm:"column:job" json:"job"`
	LimitAt      time.Time `gorm:"column:limit_at" json:"limit_at"`
	OptUID       int       `gorm:"column:opt_uid" json:"opt_uid"`
	OptUsercode  string    `gorm:"column:opt_usercode" json:"opt_usercode"`
	OptUsername  string    `gorm:"column:opt_username" json:"opt_username"`
	Phone        string    `gorm:"column:phone" json:"phone"`
	Property     string    `gorm:"column:property" json:"property"`
	RelationCode string    `gorm:"column:relation_code" json:"relation_code"`
	Status       int       `gorm:"column:status" json:"status"`
	Type         string    `gorm:"column:type" json:"type"`
	UpdateAt     time.Time `gorm:"column:update_at" json:"update_at"`
}

// TableName sets the insert table name for this struct type
func (r *RemindJob) TableName() string {
	return "remind_job"
}
