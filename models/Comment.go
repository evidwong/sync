package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	AddTime           time.Time      `gorm:"column:add_time" json:"add_time"`
	Cid               sql.NullInt64  `gorm:"column:cid" json:"cid"`
	Comno             sql.NullString `gorm:"column:comno" json:"comno"`
	Content           sql.NullString `gorm:"column:content" json:"content"`
	ID                int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Image             sql.NullString `gorm:"column:image" json:"image"`
	Nickname          sql.NullString `gorm:"column:nickname" json:"nickname"`
	Openid            sql.NullString `gorm:"column:openid" json:"openid"`
	Phone             sql.NullString `gorm:"column:phone" json:"phone"`
	ReplyComment      sql.NullString `gorm:"column:reply_comment" json:"reply_comment"`
	ReplyTime         time.Time      `gorm:"column:reply_time" json:"reply_time"`
	Score             sql.NullInt64  `gorm:"column:score" json:"score"`
	ShowImg           sql.NullString `gorm:"column:show_img" json:"show_img"`
	ShowThumb         sql.NullString `gorm:"column:show_thumb" json:"show_thumb"`
	StoreID           sql.NullInt64  `gorm:"column:store_id" json:"store_id"`
	Superaddition     sql.NullString `gorm:"column:superaddition" json:"superaddition"`
	SuperadditionTime time.Time      `gorm:"column:superaddition_time" json:"superaddition_time"`
	Thumb             sql.NullString `gorm:"column:thumb" json:"thumb"`
	Type              sql.NullString `gorm:"column:type" json:"type"`
	UID               sql.NullInt64  `gorm:"column:uid" json:"uid"`
}

// TableName sets the insert table name for this struct type
func (c *Comment) TableName() string {
	return "comment"
}
