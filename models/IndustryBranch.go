package models

import "database/sql"

type IndustryBranch struct {
	Branch sql.NullString `gorm:"column:branch" json:"branch"`
	ID     int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Remark sql.NullString `gorm:"column:remark" json:"remark"`
	Tel    sql.NullString `gorm:"column:tel" json:"tel"`
}

// TableName sets the insert table name for this struct type
func (i *IndustryBranch) TableName() string {
	return "industry_branch"
}
