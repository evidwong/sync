package models

import "database/sql"

type OfferRecord struct {
	Addtime    int            `gorm:"column:addtime" json:"addtime"`
	CardID     int            `gorm:"column:card_id" json:"card_id"`
	Cid        int            `gorm:"column:cid" json:"cid"`
	Cname      string         `gorm:"column:cname" json:"cname"`
	Content    sql.NullString `gorm:"column:content" json:"content"`
	GdwzUserid int            `gorm:"column:gdwz_userid" json:"gdwz_userid"`
	ID         int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Itemid     int            `gorm:"column:itemid" json:"itemid"`
	Nickname   string         `gorm:"column:nickname" json:"nickname"`
	Offerid    int            `gorm:"column:offerid" json:"offerid"`
	Openid     string         `gorm:"column:openid" json:"openid"`
	Phone      string         `gorm:"column:phone" json:"phone"`
	UID        int            `gorm:"column:uid" json:"uid"`
}

// TableName sets the insert table name for this struct type
func (o *OfferRecord) TableName() string {
	return "offer_record"
}
