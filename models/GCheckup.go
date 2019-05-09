package models

import "database/sql"

type GCheckup struct {
	CheckPic        sql.NullString  `gorm:"column:CheckPic" json:"CheckPic"`
	CheckPicThumb   sql.NullString  `gorm:"column:CheckPic_thumb" json:"CheckPic_thumb"`
	CheckType       sql.NullString  `gorm:"column:CheckType" json:"CheckType"`
	CheckUp         string          `gorm:"column:CheckUp" json:"CheckUp"`
	CheckUpPosition sql.NullString  `gorm:"column:CheckUpPosition" json:"CheckUpPosition"`
	CheckUpResult1  sql.NullString  `gorm:"column:CheckUpResult1" json:"CheckUpResult1"`
	CheckUpResult2  sql.NullString  `gorm:"column:CheckUpResult2" json:"CheckUpResult2"`
	CheckUpResult3  sql.NullString  `gorm:"column:CheckUpResult3" json:"CheckUpResult3"`
	Result2Data     string          `gorm:"column:Result2_Data" json:"Result2_Data"`
	Result3Data     string          `gorm:"column:Result3_Data" json:"Result3_Data"`
	TimeAmount      sql.NullFloat64 `gorm:"column:TimeAmount" json:"TimeAmount"`
	WorkTime        sql.NullFloat64 `gorm:"column:WorkTime" json:"WorkTime"`
	Cid             int             `gorm:"column:cid" json:"cid"`
	Commission      float32         `gorm:"column:commission" json:"commission"`
	Comno           sql.NullString  `gorm:"column:comno" json:"comno"`
	Comnocode       sql.NullString  `gorm:"column:comnocode" json:"comnocode"`
	CreatedAt       sql.NullInt64   `gorm:"column:created_at" json:"created_at"`
	ID              int             `gorm:"column:id;primary_key" json:"id;primary_key"`
	OptName         sql.NullString  `gorm:"column:opt_name" json:"opt_name"`
	OptUID          sql.NullInt64   `gorm:"column:opt_uid" json:"opt_uid"`
	PositionID      sql.NullInt64   `gorm:"column:position_id" json:"position_id"`
	Remark          sql.NullString  `gorm:"column:remark" json:"remark"`
	ServiceID       sql.NullInt64   `gorm:"column:service_id" json:"service_id"`
	SortID          sql.NullString  `gorm:"column:sort_id" json:"sort_id"`
	Status          int             `gorm:"column:status" json:"status"`
	TypeID          sql.NullInt64   `gorm:"column:type_id" json:"type_id"`
	UpdatedAt       sql.NullInt64   `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (g *GCheckup) TableName() string {
	return "g_checkup"
}
