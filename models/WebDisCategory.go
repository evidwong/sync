package models

import "database/sql"

type WebDisCategory struct {
	AddTime     sql.NullInt64  `gorm:"column:add_time" json:"add_time"`
	DisType     sql.NullString `gorm:"column:dis_type" json:"dis_type"`
	ID          int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Level1      sql.NullInt64  `gorm:"column:level_1" json:"level_1"`
	Level2      sql.NullInt64  `gorm:"column:level_2" json:"level_2"`
	Level3      sql.NullInt64  `gorm:"column:level_3" json:"level_3"`
	OptUID      sql.NullInt64  `gorm:"column:opt_uid" json:"opt_uid"`
	OptUsername sql.NullString `gorm:"column:opt_username" json:"opt_username"`
	OrderType   sql.NullString `gorm:"column:order_type" json:"order_type"`
	Pid         int            `gorm:"column:pid" json:"pid"`
	Remark      sql.NullString `gorm:"column:remark" json:"remark"`
	Status      int            `gorm:"column:status" json:"status"`
}

// TableName sets the insert table name for this struct type
func (w *WebDisCategory) TableName() string {
	return "web_dis_category"
}
