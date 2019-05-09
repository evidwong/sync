package models

import "database/sql"

type OpenUser struct {
	Agentid         int            `gorm:"column:agentid" json:"agentid"`
	Aid             sql.NullInt64  `gorm:"column:aid" json:"aid"`
	Cid             int            `gorm:"column:cid" json:"cid"`
	City            string         `gorm:"column:city" json:"city"`
	Comno           string         `gorm:"column:comno" json:"comno"`
	Country         string         `gorm:"column:country" json:"country"`
	GdwzOpenid      string         `gorm:"column:gdwz_openid" json:"gdwz_openid"`
	GdwzUserid      int            `gorm:"column:gdwz_userid" json:"gdwz_userid"`
	Groupid         int            `gorm:"column:groupid" json:"groupid"`
	Headimgurl      string         `gorm:"column:headimgurl" json:"headimgurl"`
	ID              int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isnew           int            `gorm:"column:isnew" json:"isnew"`
	Lat             float32        `gorm:"column:lat" json:"lat"`
	Lng             float32        `gorm:"column:lng" json:"lng"`
	LocationAddress string         `gorm:"column:location_address" json:"location_address"`
	LocationCity    string         `gorm:"column:location_city" json:"location_city"`
	Nickname        string         `gorm:"column:nickname" json:"nickname"`
	Openid          string         `gorm:"column:openid" json:"openid"`
	Province        string         `gorm:"column:province" json:"province"`
	Sex             int            `gorm:"column:sex" json:"sex"`
	Status          int            `gorm:"column:status" json:"status"`
	Subscribe       int            `gorm:"column:subscribe" json:"subscribe"`
	SubscribeTime   int            `gorm:"column:subscribe_time" json:"subscribe_time"`
	Token           sql.NullString `gorm:"column:token" json:"token"`
	Unionid         sql.NullString `gorm:"column:unionid" json:"unionid"`
	UnsubscribeTime int            `gorm:"column:unsubscribe_time" json:"unsubscribe_time"`
	UpdateTime      sql.NullInt64  `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (o *OpenUser) TableName() string {
	return "open_user"
}
