package models

import "database/sql"

type CouponCard struct {
	AddTime          sql.NullInt64   `gorm:"column:add_time" json:"add_time"`
	Brand            string          `gorm:"column:brand" json:"brand"`
	CanGive          sql.NullInt64   `gorm:"column:can_give" json:"can_give"`
	Cartype          string          `gorm:"column:cartype" json:"cartype"`
	Cid              sql.NullInt64   `gorm:"column:cid" json:"cid"`
	Cmodel           string          `gorm:"column:cmodel" json:"cmodel"`
	ComeFrom         sql.NullString  `gorm:"column:comeFrom" json:"comeFrom"`
	Comno            sql.NullString  `gorm:"column:comno" json:"comno"`
	ComnoName        sql.NullString  `gorm:"column:comno_name" json:"comno_name"`
	ComplateCode     sql.NullString  `gorm:"column:complate_code" json:"complate_code"`
	ExpireTime       sql.NullInt64   `gorm:"column:expire_time" json:"expire_time"`
	GiveTime         sql.NullInt64   `gorm:"column:give_time" json:"give_time"`
	GiveUID          sql.NullInt64   `gorm:"column:give_uid" json:"give_uid"`
	GoodsID          sql.NullInt64   `gorm:"column:goods_id" json:"goods_id"`
	GoodsName        sql.NullString  `gorm:"column:goods_name" json:"goods_name"`
	ID               int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	Idcode           sql.NullString  `gorm:"column:idcode" json:"idcode"`
	IsConvert        int             `gorm:"column:is_convert" json:"is_convert"`
	Isdel            sql.NullInt64   `gorm:"column:isdel" json:"isdel"`
	Item             sql.NullString  `gorm:"column:item" json:"item"`
	Nickname         sql.NullString  `gorm:"column:nickname" json:"nickname"`
	Openid           sql.NullString  `gorm:"column:openid" json:"openid"`
	Orderid          sql.NullString  `gorm:"column:orderid" json:"orderid"`
	Phone            sql.NullString  `gorm:"column:phone" json:"phone"`
	Price            sql.NullFloat64 `gorm:"column:price" json:"price"`
	Pwcode           sql.NullString  `gorm:"column:pwcode" json:"pwcode"`
	Randcode         sql.NullString  `gorm:"column:randcode" json:"randcode"`
	Registerno       string          `gorm:"column:registerno" json:"registerno"`
	Sendname         string          `gorm:"column:sendname" json:"sendname"`
	Status           int             `gorm:"column:status" json:"status"`
	StoreID          sql.NullInt64   `gorm:"column:store_id" json:"store_id"`
	Type             int             `gorm:"column:type" json:"type"`
	UID              sql.NullInt64   `gorm:"column:uid" json:"uid"`
	UseTime          sql.NullInt64   `gorm:"column:use_time" json:"use_time"`
	UserCode         sql.NullString  `gorm:"column:user_code" json:"user_code"`
	VCid             sql.NullInt64   `gorm:"column:v_cid" json:"v_cid"`
	VComno           sql.NullString  `gorm:"column:v_comno" json:"v_comno"`
	VComnoName       sql.NullString  `gorm:"column:v_comno_name" json:"v_comno_name"`
	VCompany         sql.NullString  `gorm:"column:v_company" json:"v_company"`
	VNickname        sql.NullString  `gorm:"column:v_nickname" json:"v_nickname"`
	VOpenid          sql.NullString  `gorm:"column:v_openid" json:"v_openid"`
	Vin              string          `gorm:"column:vin" json:"vin"`
	WechatExpirePush int             `gorm:"column:wechat_expire_push" json:"wechat_expire_push"`
}

// TableName sets the insert table name for this struct type
func (c *CouponCard) TableName() string {
	return "coupon_card"
}
