package models

import "database/sql"

type Offers struct {
	Auditname     string         `gorm:"column:auditname" json:"auditname"`
	Audittime     string         `gorm:"column:audittime" json:"audittime"`
	AutoTrigger   int            `gorm:"column:auto_trigger" json:"auto_trigger"`
	Cid           int            `gorm:"column:cid" json:"cid"`
	Enddate       int            `gorm:"column:enddate" json:"enddate"`
	ID            int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel         int            `gorm:"column:isdel" json:"isdel"`
	Item          sql.NullString `gorm:"column:item" json:"item"`
	Itemid        int            `gorm:"column:itemid" json:"itemid"`
	LimitInterval int            `gorm:"column:limit_interval" json:"limit_interval"`
	MatchMoney    float32        `gorm:"column:match_money" json:"match_money"`
	Name          string         `gorm:"column:name" json:"name"`
	Num           int            `gorm:"column:num" json:"num"`
	OfferType     string         `gorm:"column:offer_type" json:"offer_type"`
	OfferTypeText string         `gorm:"column:offer_type_text" json:"offer_type_text"`
	Sended        int            `gorm:"column:sended" json:"sended"`
	SingleMoney   float32        `gorm:"column:single_money" json:"single_money"`
	SingleNum     int            `gorm:"column:single_num" json:"single_num"`
	Startdate     int            `gorm:"column:startdate" json:"startdate"`
	Status        int            `gorm:"column:status" json:"status"`
	StoreVal      sql.NullString `gorm:"column:store_val" json:"store_val"`
	Title         string         `gorm:"column:title" json:"title"`
	TotalMoney    float32        `gorm:"column:total_money" json:"total_money"`
	Updatetime    int            `gorm:"column:updatetime" json:"updatetime"`
}

// TableName sets the insert table name for this struct type
func (o *Offers) TableName() string {
	return "offers"
}
