package models

import "database/sql"

type CTemplate struct {
	Cid          int           `gorm:"column:cid" json:"cid"`
	Comno        string        `gorm:"column:comno" json:"comno"`
	Default      string        `gorm:"column:default" json:"default"`
	Display      sql.NullInt64 `gorm:"column:display" json:"display"`
	Field        string        `gorm:"column:field" json:"field"`
	FieldType    string        `gorm:"column:field_type" json:"field_type"`
	ID           int           `gorm:"column:id;primary_key" json:"id;primary_key"`
	Identifier   int           `gorm:"column:identifier" json:"identifier"`
	Name         string        `gorm:"column:name" json:"name"`
	Readonly     int           `gorm:"column:readonly" json:"readonly"`
	Required     int           `gorm:"column:required" json:"required"`
	Sort         int           `gorm:"column:sort" json:"sort"`
	TemplateType int           `gorm:"column:template_type" json:"template_type"`
	UID          int           `gorm:"column:u_id" json:"u_id"`
	Width        int           `gorm:"column:width" json:"width"`
}

// TableName sets the insert table name for this struct type
func (c *CTemplate) TableName() string {
	return "c_template"
}
