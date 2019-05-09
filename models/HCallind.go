package models

import (
	"database/sql"
	"time"
)

type HCallind struct {
	CustFeedback     sql.NullString `gorm:"column:CustFeedback" json:"CustFeedback"`
	HCallInCode      sql.NullString `gorm:"column:HCallInCode" json:"HCallInCode"`
	IndexID          sql.NullInt64  `gorm:"column:IndexID" json:"IndexID"`
	IntentType       sql.NullString `gorm:"column:IntentType" json:"IntentType"`
	OldIntenType     sql.NullString `gorm:"column:OldIntenType" json:"OldIntenType"`
	PlanVisitContent sql.NullString `gorm:"column:PlanVisitContent" json:"PlanVisitContent"`
	PlanVisitDate    time.Time      `gorm:"column:PlanVisitDate" json:"PlanVisitDate"`
	PlanVisitType    sql.NullString `gorm:"column:PlanVisitType" json:"PlanVisitType"`
	SalesAdviser     sql.NullString `gorm:"column:SalesAdviser" json:"SalesAdviser"`
	SalesAdviserCode sql.NullString `gorm:"column:SalesAdviserCode" json:"SalesAdviserCode"`
	VisitContent     sql.NullString `gorm:"column:VisitContent" json:"VisitContent"`
	VisitDate        time.Time      `gorm:"column:VisitDate" json:"VisitDate"`
	VisitDesc        sql.NullString `gorm:"column:VisitDesc" json:"VisitDesc"`
	VisitType        sql.NullString `gorm:"column:VisitType" json:"VisitType"`
	Cid              sql.NullInt64  `gorm:"column:cid" json:"cid"`
	ID               int            `gorm:"column:id;primary_key" json:"id;primary_key"`
	Pid              int            `gorm:"column:pid" json:"pid"`
}

// TableName sets the insert table name for this struct type
func (h *HCallind) TableName() string {
	return "h_callind"
}
