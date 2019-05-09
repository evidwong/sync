package models

import (
	"database/sql"
	"time"
)

type WebArticle struct {
	ArticleContent  sql.NullString `gorm:"column:article_content" json:"article_content"`
	ArticleDescribe string         `gorm:"column:article_describe" json:"article_describe"`
	ArticleImage    string         `gorm:"column:article_image" json:"article_image"`
	ArticleThumb    string         `gorm:"column:article_thumb" json:"article_thumb"`
	ArticleTitle    string         `gorm:"column:article_title" json:"article_title"`
	CategoryID      sql.NullInt64  `gorm:"column:category_id" json:"category_id"`
	Cid             sql.NullInt64  `gorm:"column:cid" json:"cid"`
	CreateTime      time.Time      `gorm:"column:create_time" json:"create_time"`
	CreateUser      string         `gorm:"column:create_user" json:"create_user"`
	ID              int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	ReadNum         int            `gorm:"column:read_num" json:"read_num"`
	SharedNum       int            `gorm:"column:shared_num" json:"shared_num"`
	Sort            sql.NullInt64  `gorm:"column:sort" json:"sort"`
	UpdateTime      time.Time      `gorm:"column:update_time" json:"update_time"`
	UpdateUser      string         `gorm:"column:update_user" json:"update_user"`
	URL             sql.NullString `gorm:"column:url" json:"url"`
}

// TableName sets the insert table name for this struct type
func (w *WebArticle) TableName() string {
	return "web_article"
}
