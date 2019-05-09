package models

type MessageTmpl struct {
	Cid     int    `gorm:"column:cid" json:"cid"`
	Content string `gorm:"column:content" json:"content"`
	ID      int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name    string `gorm:"column:name" json:"name"`
	Status  int    `gorm:"column:status" json:"status"`
	Type    string `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (m *MessageTmpl) TableName() string {
	return "message_tmpl"
}
