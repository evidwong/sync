package models

import "database/sql"

type WechatTempl struct {
	Cron       sql.NullString `gorm:"column:cron" json:"cron"`
	ID         int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Obj        int            `gorm:"column:obj" json:"obj"`
	Status     int            `gorm:"column:status" json:"status"`
	TemplDesc  sql.NullString `gorm:"column:templ_desc" json:"templ_desc"`
	TemplIndex string         `gorm:"column:templ_index" json:"templ_index"`
	Templid    string         `gorm:"column:templid" json:"templid"`
	Title      string         `gorm:"column:title" json:"title"`
	UpdateTime int            `gorm:"column:update_time" json:"update_time"`
	Workerid   int            `gorm:"column:workerid" json:"workerid"`
}

// TableName sets the insert table name for this struct type
func (w *WechatTempl) TableName() string {
	return "wechat_templ"
}
