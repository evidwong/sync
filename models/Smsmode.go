package models

import "database/sql"

type Smsmode struct {
	Cid      sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Comno    sql.NullString `gorm:"column:comno" json:"comno"`
	Content  sql.NullString `gorm:"column:content" json:"content"`
	ID       int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	ModeCode sql.NullString `gorm:"column:mode_code" json:"mode_code"`
	ModeType sql.NullInt64  `gorm:"column:mode_type" json:"mode_type"`
	Modeid   sql.NullString `gorm:"column:modeid" json:"modeid"`
	Name     sql.NullString `gorm:"column:name" json:"name"`
	Number   sql.NullInt64  `gorm:"column:number" json:"number"`
	Storeid  sql.NullInt64  `gorm:"column:storeid" json:"storeid"`
	Type     sql.NullInt64  `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (s *Smsmode) TableName() string {
	return "smsmode"
}
