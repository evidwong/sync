package models

import "database/sql"

type EpcAcronym struct {
	ID   int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name sql.NullString `gorm:"column:name" json:"name"`
}

// TableName sets the insert table name for this struct type
func (e *EpcAcronym) TableName() string {
	return "epc_acronym"
}
