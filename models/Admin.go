package models

import "database/sql"
//Admin ...
type Admin struct {
	Department       sql.NullString `gorm:"column:Department" json:"Department"`
	ActionList       sql.NullString `gorm:"column:action_list" json:"action_list"`
	AddTime          int            `gorm:"column:add_time" json:"add_time"`
	Cid              int            `gorm:"column:cid" json:"cid"`
	Email            string         `gorm:"column:email" json:"email"`
	ExtendPermission sql.NullString `gorm:"column:extend_permission" json:"extend_permission"`
	GroupID          sql.NullString `gorm:"column:groupId" json:"groupId"`
	GroupIds          sql.NullString `gorm:"column:group_id" json:"group_id"`
	ID               int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isclerk          sql.NullInt64  `gorm:"column:isclerk" json:"isclerk"`
	Jobid            sql.NullString `gorm:"column:jobid" json:"jobid"`
	LastIP           string         `gorm:"column:last_ip" json:"last_ip"`
	LastLogin        int            `gorm:"column:last_login" json:"last_login"`
	Mid              int            `gorm:"column:mid" json:"mid"`
	Name             sql.NullString `gorm:"column:name" json:"name"`
	Nickname         sql.NullString `gorm:"column:nickname" json:"nickname"`
	Openid           sql.NullString `gorm:"column:openid" json:"openid"`
	Password         string         `gorm:"column:password" json:"password"`
	Permission       sql.NullString `gorm:"column:permission" json:"permission"`
	Phone            sql.NullString `gorm:"column:phone" json:"phone"`
	Pid              int            `gorm:"column:pid" json:"pid"`
	Remark           sql.NullString `gorm:"column:remark" json:"remark"`
	RoleID           sql.NullInt64  `gorm:"column:role_id" json:"role_id"`
	Roleid           sql.NullInt64  `gorm:"column:roleid" json:"roleid"`
	Salt             string         `gorm:"column:salt" json:"salt"`
	Status           sql.NullInt64  `gorm:"column:status" json:"status"`
	StoreID          sql.NullString `gorm:"column:store_id" json:"store_id"`
	Type             int            `gorm:"column:type" json:"type"`
	UID              sql.NullInt64  `gorm:"column:uid" json:"uid"`
	Username         string         `gorm:"column:username" json:"username"`
	WechatName       sql.NullString `gorm:"column:wechat_name" json:"wechat_name"`
}

// TableName sets the insert table name for this struct type
func (a *Admin) TableName() string {
	return "admin"
}
