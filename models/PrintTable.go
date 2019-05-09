package models

type PrintTable struct {
	ID int `gorm:"column:id;primary_key" json:"id;primary_key"`
}

// TableName sets the insert table name for this struct type
func (p *PrintTable) TableName() string {
	return "print_table"
}
