package models

import "database/sql"

type WebDistributor struct {
	AddTime       int            `gorm:"column:add_time" json:"add_time"`
	ApplyTime     int            `gorm:"column:apply_time" json:"apply_time"`
	CatID         int            `gorm:"column:cat_id" json:"cat_id"`
	CatName       sql.NullString `gorm:"column:cat_name" json:"cat_name"`
	Childrens     sql.NullInt64  `gorm:"column:childrens" json:"childrens"`
	Cid           int            `gorm:"column:cid" json:"cid"`
	ClientName    sql.NullString `gorm:"column:client_name" json:"client_name"`
	Comno         sql.NullString `gorm:"column:comno" json:"comno"`
	CustomerName  sql.NullString `gorm:"column:customer_name" json:"customer_name"`
	Fname         sql.NullString `gorm:"column:fname" json:"fname"`
	Ftype         sql.NullString `gorm:"column:ftype" json:"ftype"`
	HandPhone     string         `gorm:"column:hand_phone" json:"hand_phone"`
	ID            int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Integral      int            `gorm:"column:integral" json:"integral"`
	IntegralLave  int            `gorm:"column:integral_lave" json:"integral_lave"`
	LashShareTime int            `gorm:"column:lash_share_time" json:"lash_share_time"`
	Level         int            `gorm:"column:level" json:"level"`
	Openid        sql.NullString `gorm:"column:openid" json:"openid"`
	OptUID        int            `gorm:"column:opt_uid" json:"opt_uid"`
	OptUsername   sql.NullString `gorm:"column:opt_username" json:"opt_username"`
	Pid           int            `gorm:"column:pid" json:"pid"`
	Recommended   sql.NullString `gorm:"column:recommended" json:"recommended"`
	Status        int            `gorm:"column:status" json:"status"`
	StoreID       int            `gorm:"column:store_id" json:"store_id"`
	UID           int            `gorm:"column:uid" json:"uid"`
	UpdateTime    int            `gorm:"column:update_time" json:"update_time"`
	VerifyTime    int            `gorm:"column:verify_time" json:"verify_time"`
}

// TableName sets the insert table name for this struct type
func (w *WebDistributor) TableName() string {
	return "web_distributor"
}
