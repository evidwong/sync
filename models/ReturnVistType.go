package models

type ReturnVistType struct {
	AddTime    int    `gorm:"column:add_time" json:"add_time"`
	Aid        int    `gorm:"column:aid" json:"aid"`
	ID         int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Title      string `gorm:"column:title" json:"title"`
	Type       int    `gorm:"column:type" json:"type"`
	UpdateTime int    `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (r *ReturnVistType) TableName() string {
	return "return_vist_type"
}
