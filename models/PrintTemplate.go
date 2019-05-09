package models

import "database/sql"

type PrintTemplate struct {
	Cid          int            `gorm:"column:cid" json:"cid"`
	Comno        sql.NullString `gorm:"column:comno" json:"comno"`
	Descs        sql.NullString `gorm:"column:descs" json:"descs"`
	Functioncode string         `gorm:"column:functioncode" json:"functioncode"`
	ID           int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	OptTime      sql.NullInt64  `gorm:"column:opt_time" json:"opt_time"`
	OptUID       sql.NullInt64  `gorm:"column:opt_uid" json:"opt_uid"`
	OptUsername  sql.NullString `gorm:"column:opt_username" json:"opt_username"`
	Reportname   sql.NullString `gorm:"column:reportname" json:"reportname"`
	Status       sql.NullInt64  `gorm:"column:status" json:"status"`
	StoreID      sql.NullString `gorm:"column:store_id" json:"store_id"`
	StoreName    sql.NullString `gorm:"column:store_name" json:"store_name"`
	Title        string         `gorm:"column:title" json:"title"`
	Urlname      sql.NullString `gorm:"column:urlname" json:"urlname"`
}

// TableName sets the insert table name for this struct type
func (p *PrintTemplate) TableName() string {
	return "print_template"
}
