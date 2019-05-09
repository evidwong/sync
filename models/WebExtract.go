package models

import "database/sql"

type WebExtract struct {
	AddTime      sql.NullInt64   `gorm:"column:add_time" json:"add_time"`
	Amount       sql.NullFloat64 `gorm:"column:amount" json:"amount"`
	Cid          sql.NullInt64   `gorm:"column:cid" json:"cid"`
	CustomerName sql.NullString  `gorm:"column:customer_name" json:"customer_name"`
	Flag         sql.NullInt64   `gorm:"column:flag" json:"flag"`
	FlagContent  sql.NullString  `gorm:"column:flag_content" json:"flag_content"`
	HandPhone    sql.NullString  `gorm:"column:hand_phone" json:"hand_phone"`
	ID           int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	OptUID       sql.NullInt64   `gorm:"column:opt_uid" json:"opt_uid"`
	OptUsername  sql.NullString  `gorm:"column:opt_username" json:"opt_username"`
	Quota        sql.NullInt64   `gorm:"column:quota" json:"quota"`
	Remark       sql.NullString  `gorm:"column:remark" json:"remark"`
	Status       sql.NullInt64   `gorm:"column:status" json:"status"`
	Type         sql.NullString  `gorm:"column:type" json:"type"`
	UID          sql.NullInt64   `gorm:"column:uid" json:"uid"`
	VerifyTime   sql.NullInt64   `gorm:"column:verify_time" json:"verify_time"`
}

// TableName sets the insert table name for this struct type
func (w *WebExtract) TableName() string {
	return "web_extract"
}
