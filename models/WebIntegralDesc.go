package models

import "database/sql"

type WebIntegralDesc struct {
	AddTime        sql.NullInt64  `gorm:"column:add_time" json:"add_time"`
	ID             int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	IntegralLimit  sql.NullString `gorm:"column:integral_limit" json:"integral_limit"`
	IntegralOffset sql.NullString `gorm:"column:integral_offset" json:"integral_offset"`
	Level          sql.NullString `gorm:"column:level" json:"level"`
	OptUID         sql.NullInt64  `gorm:"column:opt_uid" json:"opt_uid"`
	OptUsername    sql.NullString `gorm:"column:opt_username" json:"opt_username"`
	Pid            int            `gorm:"column:pid" json:"pid"`
	Remark         sql.NullString `gorm:"column:remark" json:"remark"`
	Status         sql.NullInt64  `gorm:"column:status" json:"status"`
	UpdateTime     sql.NullInt64  `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (w *WebIntegralDesc) TableName() string {
	return "web_integral_desc"
}
