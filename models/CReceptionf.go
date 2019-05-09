package models

import "database/sql"

type CReceptionf struct {
	BadPic         sql.NullString `gorm:"column:BadPic" json:"BadPic"`
	CReceptionCode string         `gorm:"column:CReceptionCode" json:"CReceptionCode"`
	CheckM         string         `gorm:"column:CheckM" json:"CheckM"`
	CheckPeople    string         `gorm:"column:CheckPeople" json:"CheckPeople"`
	CheckPic       string         `gorm:"column:CheckPic" json:"CheckPic"`
	CheckType      string         `gorm:"column:CheckType" json:"CheckType"`
	CheckUp        sql.NullString `gorm:"column:CheckUp" json:"CheckUp"`
	CheckUpCode    sql.NullString `gorm:"column:CheckUpCode" json:"CheckUpCode"`
	CheckUpID      int            `gorm:"column:CheckUpId" json:"CheckUpId"`
	CheckUpName    sql.NullString `gorm:"column:CheckUpName" json:"CheckUpName"`
	CheckUpResult  sql.NullString `gorm:"column:CheckUpResult" json:"CheckUpResult"`
	CheckUpResult1 sql.NullString `gorm:"column:CheckUpResult1" json:"CheckUpResult1"`
	CheckUpResult2 sql.NullString `gorm:"column:CheckUpResult2" json:"CheckUpResult2"`
	CheckUpResult3 sql.NullString `gorm:"column:CheckUpResult3" json:"CheckUpResult3"`
	FunctionCode   string         `gorm:"column:FunctionCode" json:"FunctionCode"`
	IndexID        int            `gorm:"column:IndexID" json:"IndexID"`
	IsCheckItem    sql.NullInt64  `gorm:"column:IsCheckItem" json:"IsCheckItem"`
	IsStatus       sql.NullInt64  `gorm:"column:IsStatus" json:"IsStatus"`
	Remark         string         `gorm:"column:Remark" json:"Remark"`
	Result2Data    string         `gorm:"column:Result2_Data" json:"Result2_Data"`
	Result3Data    string         `gorm:"column:Result3_Data" json:"Result3_Data"`
	TimeAmount     float64        `gorm:"column:TimeAmount" json:"TimeAmount"`
	WorkTime       int            `gorm:"column:WorkTime" json:"WorkTime"`
	Cid            int            `gorm:"column:cid" json:"cid"`
	ID             int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	PositionID     int            `gorm:"column:position_id" json:"position_id"`
	ServiceID      sql.NullInt64  `gorm:"column:service_id" json:"service_id"`
	Type           int            `gorm:"column:type" json:"type"`
	TypeID         int            `gorm:"column:type_id" json:"type_id"`
}

// TableName sets the insert table name for this struct type
func (c *CReceptionf) TableName() string {
	return "c_receptionf"
}
