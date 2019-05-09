package models

import "database/sql"

type LotteryDeploy struct {
	ActivityID   sql.NullInt64  `gorm:"column:activity_id" json:"activity_id"`
	Cid          sql.NullInt64  `gorm:"column:cid" json:"cid"`
	ID           int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	InportLayout sql.NullString `gorm:"column:inport_layout" json:"inport_layout"`
	Isdel        sql.NullInt64  `gorm:"column:isdel" json:"isdel"`
	Isrequired   sql.NullInt64  `gorm:"column:isrequired" json:"isrequired"`
	NameShow     sql.NullString `gorm:"column:name_show" json:"name_show"`
	Storeid      sql.NullInt64  `gorm:"column:storeid" json:"storeid"`
}

// TableName sets the insert table name for this struct type
func (l *LotteryDeploy) TableName() string {
	return "lottery_deploy"
}
