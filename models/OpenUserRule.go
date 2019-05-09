package models

type OpenUserRule struct {
	Cid           int    `gorm:"column:cid" json:"cid"`
	DayReport     int    `gorm:"column:day_report" json:"day_report"`
	GroupID       int    `gorm:"column:group_id" json:"group_id"`
	ID            int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	Isdel         int    `gorm:"column:isdel" json:"isdel"`
	Openid        string `gorm:"column:openid" json:"openid"`
	Permission    string `gorm:"column:permission" json:"permission"`
	ReportTime    int    `gorm:"column:report_time" json:"report_time"`
	ReserveNotify int    `gorm:"column:reserve_notify" json:"reserve_notify"`
	Status        int    `gorm:"column:status" json:"status"`
	Tagid         int    `gorm:"column:tagid" json:"tagid"`
	UID           int    `gorm:"column:uid" json:"uid"`
	Updatetime    int    `gorm:"column:updatetime" json:"updatetime"`
}

// TableName sets the insert table name for this struct type
func (o *OpenUserRule) TableName() string {
	return "open_user_rule"
}
