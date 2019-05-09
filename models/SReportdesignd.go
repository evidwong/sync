package models

import "database/sql"

type SReportdesignd struct {
	Functioncode sql.NullString `gorm:"column:functioncode" json:"functioncode"`
	Reportname   sql.NullString `gorm:"column:reportname" json:"reportname"`
	Urlname      sql.NullString `gorm:"column:urlname" json:"urlname"`
}

// TableName sets the insert table name for this struct type
func (s *SReportdesignd) TableName() string {
	return "s_reportdesignd"
}
