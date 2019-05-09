package models

import "database/sql"

type ArticleUserHistory struct {
	RegisterNo       sql.NullString `gorm:"column:RegisterNo" json:"RegisterNo"`
	ActionType       sql.NullString `gorm:"column:action_type" json:"action_type"`
	ArticleID        sql.NullInt64  `gorm:"column:article_id" json:"article_id"`
	ArticleTitle     sql.NullString `gorm:"column:article_title" json:"article_title"`
	ArticleURL       sql.NullString `gorm:"column:article_url" json:"article_url"`
	Cid              int            `gorm:"column:cid" json:"cid"`
	CreatedAt        sql.NullInt64  `gorm:"column:created_at" json:"created_at"`
	FromUser         sql.NullString `gorm:"column:from_user" json:"from_user"`
	FromUserNickname sql.NullString `gorm:"column:from_user_nickname" json:"from_user_nickname"`
	FromUserPhone    sql.NullString `gorm:"column:from_user_phone" json:"from_user_phone"`
	ID               int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	IPAddress        sql.NullString `gorm:"column:ip_address" json:"ip_address"`
	ItemType         sql.NullString `gorm:"column:item_type" json:"item_type"`
	Nickname         sql.NullString `gorm:"column:nickname" json:"nickname"`
	Openid           sql.NullString `gorm:"column:openid" json:"openid"`
	Phone            sql.NullString `gorm:"column:phone" json:"phone"`
	Type             sql.NullInt64  `gorm:"column:type" json:"type"`
	UID              sql.NullInt64  `gorm:"column:uid" json:"uid"`
}

// TableName sets the insert table name for this struct type
func (a *ArticleUserHistory) TableName() string {
	return "article_user_history"
}
