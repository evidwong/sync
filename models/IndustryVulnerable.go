package models

import "database/sql"

type IndustryVulnerable struct {
	Cid             sql.NullInt64  `gorm:"column:cid" json:"cid"`
	ID              int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Index           sql.NullInt64  `gorm:"column:index" json:"index"`
	VulnerableName  sql.NullString `gorm:"column:vulnerable_name" json:"vulnerable_name"`
	VulnerablePic   sql.NullString `gorm:"column:vulnerable_pic" json:"vulnerable_pic"`
	VulnerableThumb sql.NullString `gorm:"column:vulnerable_thumb" json:"vulnerable_thumb"`
}

// TableName sets the insert table name for this struct type
func (i *IndustryVulnerable) TableName() string {
	return "industry_vulnerable"
}
