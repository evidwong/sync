package models

import "database/sql"

type Member struct {
	AddTime         int            `gorm:"column:add_time" json:"add_time"`
	Agentid         sql.NullInt64  `gorm:"column:agentid" json:"agentid"`
	Aid             sql.NullInt64  `gorm:"column:aid" json:"aid"`
	Cid             int            `gorm:"column:cid" json:"cid"`
	Comno           sql.NullString `gorm:"column:comno" json:"comno"`
	Customercode    string         `gorm:"column:customercode" json:"customercode"`
	Customername    string         `gorm:"column:customername" json:"customername"`
	Default         string         `gorm:"column:default" json:"default"`
	Email           string         `gorm:"column:email" json:"email"`
	Fuid            int            `gorm:"column:fuid" json:"fuid"`
	Gid             int            `gorm:"column:gid" json:"gid"`
	Isnew           sql.NullInt64  `gorm:"column:isnew" json:"isnew"`
	Lastloginip     string         `gorm:"column:lastloginip" json:"lastloginip"`
	Lastlogintime   int            `gorm:"column:lastlogintime" json:"lastlogintime"`
	Lat             float32        `gorm:"column:lat" json:"lat"`
	Lng             float32        `gorm:"column:lng" json:"lng"`
	LocationAddress string         `gorm:"column:location_address" json:"location_address"`
	LocationCity    string         `gorm:"column:location_city" json:"location_city"`
	Myid            string         `gorm:"column:myid" json:"myid"`
	Myidkey         string         `gorm:"column:myidkey" json:"myidkey"`
	Nickname        sql.NullString `gorm:"column:nickname" json:"nickname"`
	Password        string         `gorm:"column:password" json:"password"`
	Phone           string         `gorm:"column:phone" json:"phone"`
	Regdate         int            `gorm:"column:regdate" json:"regdate"`
	Regip           string         `gorm:"column:regip" json:"regip"`
	Salt            string         `gorm:"column:salt" json:"salt"`
	Secques         string         `gorm:"column:secques" json:"secques"`
	UID             int            `gorm:"column:uid;primary_key" json:"uid;primary_key"`
	Updatetime      int            `gorm:"column:updatetime" json:"updatetime"`
	Username        string         `gorm:"column:username" json:"username"`
	Vin             string         `gorm:"column:vin" json:"vin"`
}

// TableName sets the insert table name for this struct type
func (m *Member) TableName() string {
	return "member"
}
